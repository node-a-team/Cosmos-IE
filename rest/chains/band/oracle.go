package band

import (
	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type Oracle struct {
	Status struct {
		Is_active	bool
	}
}

func CheckOracleActive(restServer string, operAddr string, log *zap.Logger) float64 {

        var o Oracle

	// Oracle activation check
        res, _ := utils.RunRESTCommand(restServer, "/oracle/v1/validators/" +operAddr)
        json.Unmarshal(res, &o)

	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
		log.Info("\t", zap.Bool("Success", true), zap.String("Oracle Active", fmt.Sprintf("%v", o.Status.Is_active) ),)
        }

	result := utils.BoolToFloat64(o.Status.Is_active)


        return result
}

