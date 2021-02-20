package rest

import (
	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"
)

type commission struct {

//	Height	string		`json:"height"`
	Commission struct {
		Commission []Coin
	}

}

func getCommission(log *zap.Logger) []Coin {

	var c commission

	res, _ := runRESTCommand("/cosmos/distribution/v1beta1/validators/"+OperAddr +"/commission")
	json.Unmarshal(res, &c)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("\t", zap.Bool("Success", true), zap.String("Commission", fmt.Sprint(c.Commission.Commission)),)
        }

	return c.Commission.Commission
}
