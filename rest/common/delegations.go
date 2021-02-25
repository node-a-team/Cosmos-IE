package rest

import (
	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type delegations struct {
//	Height	string	`json:"height"`

	Delegation_responses []struct {
		Delegation delegation
	}

	Pagination struct {
		Total	string
	}
}

type selfDelegation struct {
//      Height  string  `json:"height"`
        Delegation_response struct {
		Delegation delegation
	}
}

type delegation struct {
	Delegator_address	string	`json:"delegator_address"`
	Validator_address	string	`json:"validator_address"`
	Shares			string	`json:"shares"`
}

type delegationInfo struct {
	DelegationCount	float64
	SelfDelegation	float64
}

var (
	dInfo delegationInfo
)

func getDelegations(log *zap.Logger) delegationInfo {
	var d delegations

	res, _ := runRESTCommand("/cosmos/staking/v1beta1/validators/" +OperAddr +"/delegations?pagination.limit=10000")
	json.Unmarshal(res, &d)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("", zap.Bool("Success", true), zap.String("Delegation Count", fmt.Sprint(len(d.Delegation_responses))),)
        }

	dInfo.DelegationCount = float64(len(d.Delegation_responses))

	for _, value := range d.Delegation_responses {
		if AccAddr == value.Delegation.Delegator_address {
			dInfo.SelfDelegation = utils.StringToFloat64(value.Delegation.Shares)
		}
	}


	return dInfo
}
