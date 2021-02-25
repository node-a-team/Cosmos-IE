package rest

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type oracle struct {
	Miss	float64
}

type oracleMiss struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func getOracleMiss(log *zap.Logger) oracle {

	var o oracle
        var om oracleMiss

        res, _ := runRESTCommand("/oracle/voters/" +OperAddr +"/miss")
        json.Unmarshal(res, &om)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Oracle Miss"),)
        }

	o.Miss = utils.StringToFloat64(om.Result)

        return o
}
