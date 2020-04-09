package rest

import (
	"go.uber.org/zap"
	"os/exec"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

var (
        Addr string
	OperAddr string
)


type RESTData struct {

	BlockHeight	int64
	StakingPool	stakingPool

	Validatorsets	map[string][]string
	Validators	validator
	Delegations	delegationInfo
	Balances	[]Coin
	Rewards		[]Coin
	Commission	[]Coin

	Oracle		oracle
	Gov		govInfo
}

func newRESTData(blockHeight int64) *RESTData {

	rd := &RESTData {
		BlockHeight:	blockHeight,
		Validatorsets:	make(map[string][]string),
        }

	return rd
}

func GetData(blockHeight int64, log *zap.Logger) (*RESTData, string) {


	accAddr := utils.GetAccAddrFromOperAddr(OperAddr, log)

	rd := newRESTData(blockHeight)
	rd.StakingPool = getStakingPool(log)

	rd.Validatorsets = getValidatorsets(blockHeight, log)
	rd.Validators = getValidators(log)
	rd.Delegations = getDelegations(accAddr, log)
	rd.Balances = getBalances(accAddr, log)
	rd.Rewards, rd.Commission = getRewardsAndCommisson(log)

	rd.Oracle = getOracleMiss(log)
	rd.Gov = getGovInfo(log)

	consHexAddr := utils.Bech32AddrToHexAddr(rd.Validatorsets[rd.Validators.ConsPubKey][0], log)
	return rd, consHexAddr
}

func runRESTCommand(str string) ([]uint8, error) {
        cmd := "curl -s -XGET " +Addr +str +" -H \"accept:application/json\""
        out, err := exec.Command("/bin/bash", "-c", cmd).Output()

        return out, err
}
