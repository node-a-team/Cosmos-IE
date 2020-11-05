package terra

import (
	"strings"
	"go.uber.org/zap"
	"encoding/json"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

var (
	FeeDenom string
)

type Oracle struct {
	Miss	float64
	Feeder Feeder
}

type Feeder struct {
	Address  string
        Balance  Balance
}
type Balance  struct {
        Denom string
        Amount string
}

type Result struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func GetOracle(restServer string, operAddr string, o Oracle, log *zap.Logger) Oracle {

        var r Result

	// Oracle Missing
        res, _ := utils.RunRESTCommand(restServer, "/oracle/voters/" +operAddr +"/miss")
        json.Unmarshal(res, &r)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
		log.Info("\t", zap.Bool("Success", true), zap.String("Terra Oracle Miss", r.Result ),)
        }

	o.Miss = utils.StringToFloat64(r.Result)


	// Feeder Info
	res, _ = utils.RunRESTCommand(restServer, "/oracle/voters/" +operAddr +"/feeder")
        json.Unmarshal(res, &r)
        // log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
		log.Info("\t", zap.Bool("Success", true), zap.String("Terra Oracle Feeder's Address", r.Result),)
        }

	o.Feeder.Address = r.Result

        return o
}

