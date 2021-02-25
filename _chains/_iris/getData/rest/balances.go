package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

type balances struct {
	//	Height	string	`json:"height"`
	Type  string
	Value struct {
		address string
		Coins   []Coin
	}
}

type Coin struct {
	Denom  string
	Amount string
}

func getBalances(accAddr string, log *zap.Logger) []Coin {

	var b balances

	res, _ := runRESTCommand("/bank/accounts/" + accAddr)
	json.Unmarshal(res, &b)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Staking Pool"))
	}

	return b.Value.Coins
}
