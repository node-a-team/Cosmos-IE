package rest

import (
//	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type inflation struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

type inflation_iris struct {
	Params struct {
		Mint_Denom	string
		Inflation	string
	}
}

type inflation_Emoney struct {
        Height string   `json:"height"`
        Result struct {
		Assets []struct {
			Denom		string	`json:"denom"`
			Inflation	string	`json:"inflation"`
			Accum		string	`json:"accum"`
		}
	}
}


func getInflation(chain string, denom string, log *zap.Logger) float64 {

	var result string

	switch chain {
	case "iris":
		var i inflation_iris

		res, _ := runRESTCommand("/irishub/mint/params")
                json.Unmarshal(res, &i)

                // log
                if strings.Contains(string(res), "not found") {
                        // handle error
                        log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
                } else {
                        result = i.Params.Inflation
                        log.Info("\t", zap.Bool("Success", true), zap.String("Inflation", result),)
                }

	case "terra":
		break
	case "emoney":
		var i inflation_Emoney

		res, _ := runRESTCommand("/inflation/current")
		json.Unmarshal(res, &i)

		// log
	        if strings.Contains(string(res), "not found") {
                        // handle error
                        log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
                } else {

			for _, value := range i.Result.Assets {
				if value.Denom == denom {
                                        result = value.Inflation
                                }
                        }

                        log.Info("\t", zap.Bool("Success", true), zap.String("Inflation", result),)
                }
	default:
		var i inflation

		// Does not work
		// res, _ := runRESTCommand("/cosmos/mint/v1beta1/inflation")

		res, _ := runRESTCommand("/minting/inflation")
		json.Unmarshal(res, &i)

		// log 
		if strings.Contains(string(res), "not found") {
	                // handle error
	                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
	        } else {
			result = i.Result
	                log.Info("\t", zap.Bool("Success", true), zap.String("Inflation", result),)
	        }
	}

	return utils.StringToFloat64(result)
}

