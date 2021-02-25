package rest

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type delegations struct {
	Height	string	`json:"height"`
	Result	[]delegation
}

type delegation struct {
	Delegator_address	string	`json:"delegator_address"`
	Validator_address	string	`json:"validator_address"`
	Shares			string	`json:"shares"`
	Balance			string	`json:"balance"`
}

type delegationInfo struct {
	DelegationCount	float64
	SelfDelegation	float64
}

func getDelegations(accAddr string, log *zap.Logger) delegationInfo {

	var d delegations
	var dInfo delegationInfo

	res, _ := runRESTCommand("/staking/validators/" +OperAddr +"/delegations")
	json.Unmarshal(res, &d)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Delegations"),)
        }

	dInfo.DelegationCount = float64(len(d.Result))

	for _, value := range d.Result {
		if accAddr == value.Delegator_address {
			dInfo.SelfDelegation = utils.StringToFloat64(value.Shares)
		}
	}

	return dInfo
}
