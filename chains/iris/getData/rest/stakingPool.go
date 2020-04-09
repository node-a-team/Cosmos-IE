package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

type stakingPool struct {
	Loose_tokens  string `json:"loose_tokens"`
	Bonded_tokens string `json:"bonded_tokens"`
	Total_supply  string `json:"total_supply"`
	Bonded_ratio  string `json:"bonded_ratio"`
}

func getStakingPool(log *zap.Logger) stakingPool {

	var sp stakingPool

	res, _ := runRESTCommand("/stake/pool")
	json.Unmarshal(res, &sp)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Staking Pool"))
	}

	return sp
}
