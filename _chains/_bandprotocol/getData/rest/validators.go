package rest

import (
	"strings"
	"encoding/json"
	"go.uber.org/zap"
)

type validators struct {
	Height string `json:"height"`
	Result validator
}

type validator struct {
	OperAddr        string `json:"operator_address"`
	ConsPubKey      string `json:"consensus_pubkey"`
	Jailed          bool   `json:"jailed"`
	Status          int    `json:"status"`
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}
	UnbondingHeight string `json:"unbonding_height"`
	UnbondingTime   string `json:"unbonding_time"`
	Commission      struct {
		Commission_rates struct {
			Rate          string `json:"rate"`
			Max_rate       string `json:"max_rate"`
			Max_change_rate string `json:"max_change_rate"`
		}
		UpdateTime string `json:"update_time"`
	}
	MinSelfDelegation string `json:"min_self_delegation"`
}

func getValidators(log *zap.Logger) validator {

	var v validators

	res, _ := runRESTCommand("/staking/validators/" +OperAddr)
	json.Unmarshal(res, &v)

        // log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Validators"),)
        }

	return v.Result
}
