package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"

	utils "github.com/node-a-team/cosmos_metric/utils"
)

/*
type delegations struct {
	Height string `json:"height"`
	Result []delegation
}
*/

type delegation struct {
	Delegator_address string `json:"delegator_addr"`
	Validator_address string `json:"validator_addr"`
	Shares            string `json:"shares"`
	Balance           string `json:"balance"`
}

type delegationInfo struct {
	DelegationCount float64
	SelfDelegation  float64
}

func getDelegations(accAddr string, log *zap.Logger) delegationInfo {

	var d []delegation
	var dInfo delegationInfo

	res, _ := runRESTCommand("/stake/validators/" + OperAddr + "/delegations")
	json.Unmarshal(res, &d)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Delegations"))
	}

	dInfo.DelegationCount = float64(len(d))

	for _, value := range d {
		if accAddr == value.Delegator_address {
			dInfo.SelfDelegation = utils.StringToFloat64(value.Shares)
		}
	}

	return dInfo
}
