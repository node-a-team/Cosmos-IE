package rest

import (
	"go.uber.org/zap"
	"os/exec"

	utils "github.com/node-a-team/Cosmos-IE/utils"
	terra "github.com/node-a-team/Cosmos-IE/rest/chains/terra"
)

var (
        Addr string
	OperAddr string
)


type RESTData struct {

	BlockHeight	int64
	Commit		commitInfo
	StakingPool	stakingPool

	Validatorsets	map[string][]string
	Validators	validator
	Delegations	delegationInfo
	Balances	[]Coin
	Rewards		[]Coin
	Commission	[]Coin

	Oracle		terra.Oracle
	Gov		govInfo
}

func newRESTData(blockHeight int64) *RESTData {

	rd := &RESTData {
		BlockHeight:	blockHeight,
		Validatorsets:	make(map[string][]string),
        }

	return rd
}

func GetData(blockHeight int64, blockData Blocks, log *zap.Logger) (*RESTData) {


	accAddr := utils.GetAccAddrFromOperAddr(OperAddr, log)

	rd := newRESTData(blockHeight)
	rd.StakingPool = getStakingPool(log)

	rd.Validatorsets = getValidatorsets(blockHeight, log)
	rd.Validators = getValidators(log)
	rd.Delegations = getDelegations(accAddr, log)
	rd.Balances = getBalances(accAddr, log)
	rd.Rewards, rd.Commission = getRewardsAndCommisson(log)
	rd.Gov = getGovInfo(log)

	// for Terra
	rd.Oracle = terra.GetOracle(Addr, OperAddr,
		terra.Oracle{Feeder: terra.Feeder{Balance: terra.Balance{Denom: terra.FeeDenom }} },
		log,
	)

	for _, v := range getBalances(rd.Oracle.Feeder.Address, log) {
		if v.Denom == terra.FeeDenom {
			rd.Oracle.Feeder.Balance.Denom = v.Denom
			rd.Oracle.Feeder.Balance.Amount = v.Amount
		}
	}

	consHexAddr := utils.Bech32AddrToHexAddr(rd.Validatorsets[rd.Validators.ConsPubKey][0], log)
	rd.Commit = getCommit(blockData, consHexAddr)

	return rd
}

func runRESTCommand(str string) ([]uint8, error) {
        cmd := "curl -s -XGET " +Addr +str +" -H \"accept:application/json\""
        out, err := exec.Command("/bin/bash", "-c", cmd).Output()

        return out, err
}
