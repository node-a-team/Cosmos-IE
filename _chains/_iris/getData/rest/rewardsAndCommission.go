package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

type rewardsAndCommisson struct {
	Total       []Coin
	Delegations []struct {
		Validator string
		Reward    []Coin
	}
	Commission []Coin
}

func getRewardsAndCommisson(accAddr string, log *zap.Logger) ([]Coin, []Coin) {

	var rc rewardsAndCommisson
	var reward []Coin

	res, _ := runRESTCommand("/distribution/" + accAddr + "/rewards")
	json.Unmarshal(res, &rc)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Rewards&Commission"))
	}

	for _, value := range rc.Delegations {
		if value.Validator == OperAddr {
			reward = value.Reward
		}
	}

	return reward, rc.Commission
}
