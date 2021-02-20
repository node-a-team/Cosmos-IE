package rest

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type inflation struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func getInflation(log *zap.Logger) float64 {

	var i inflation

	res, _ := runRESTCommand("/minting/inflation")
	json.Unmarshal(res, &i)

	// log 
	if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Inflation"),)
        }

        return utils.StringToFloat64(i.Result)
}
