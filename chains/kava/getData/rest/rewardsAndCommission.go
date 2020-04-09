package rest

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"
)

type rewardsAndCommisson struct {

	Height	string		`json:"height"`
	Result	struct {
                Operator_Address        string  `"json:"operator_address"`
                Self_bond_rewards         []Coin  `"json:"self_bond_rewards"`
                Val_commission      []Coin  `"json:"val_commission"`
	}

}

func getRewardsAndCommisson(log *zap.Logger) ([]Coin, []Coin) {

	var rc rewardsAndCommisson

	res, _ := runRESTCommand("/distribution/validators/" +OperAddr)
	json.Unmarshal(res, &rc)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Rewards&Commission"),)
        }

	return rc.Result.Self_bond_rewards, rc.Result.Val_commission
}
