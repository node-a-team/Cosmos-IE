package kava

import (
	"fmt"
	"net/http"
	"go.uber.org/zap"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/node-a-team/Cosmos-IE/chains/kava/exporter"
)

const (
	bech32MainPrefix = "kava"
	bip44CoinType = 459
)

func Main(port string) {

	log,_ := zap.NewDevelopment()
        defer log.Sync()

	config := sdk.GetConfig()
	config.SetCoinType(bip44CoinType)
	config.SetBech32PrefixForAccount(bech32MainPrefix, bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)
	config.Seal()

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":" +port, nil)
	// log
        if err != nil {
                // handle error
                log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: " +port +")"),)
        }
}
