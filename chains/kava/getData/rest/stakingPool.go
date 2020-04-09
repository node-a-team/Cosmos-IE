package rest

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type stakingPool struct {
	Height	string	`json:"height"`
	Result	struct {
		Not_bonded_tokens	string	`json:"not_bonded_tokens"`
		Bonded_tokens		string	`json:"bonded_tokens"`
		Total_supply		float64
	}
}

type totalSupply struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func getStakingPool(log *zap.Logger) stakingPool {

	var sp stakingPool

	res, _ := runRESTCommand("/staking/pool")
	json.Unmarshal(res, &sp)

	// log 
	if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Staking Pool"),)
        }

	sp.Result.Total_supply = getTotalSupply("kava", log)

	return sp
}

func getTotalSupply(denom string, log *zap.Logger) float64 {

        var ts totalSupply

        res, _ := runRESTCommand("/supply/total/u" +denom)
        json.Unmarshal(res, &ts)

	// log
	if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Total Supply"),)
        }

        return utils.StringToFloat64(ts.Result)
}
