package rest

import (
	"fmt"
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
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("\t", zap.Bool("Success", true), zap.String("Rewards", fmt.Sprint(rc.Result.Self_bond_rewards)),)
                log.Info("\t", zap.Bool("Success", true), zap.String("Commission", fmt.Sprint(rc.Result.Val_commission)),)
        }

	return rc.Result.Self_bond_rewards, rc.Result.Val_commission
}
