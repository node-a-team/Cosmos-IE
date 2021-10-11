package exporter

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	sdk "github.com/cosmos/cosmos-sdk/types"
	iris "github.com/irisnet/irishub/address"
	"github.com/node-a-team/Cosmos-IE/common"

	terra "github.com/terra-project/core/types"
	//	kava "github.com/kava-labs/kava/app"
	emoney "github.com/e-money/em-ledger/types"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ()

func Go(chain string, port string) {

	log, _ := zap.NewDevelopment()
	defer log.Sync()

	setConfig(chain)

	http.Handle("/metrics", promhttp.Handler())
	go Start(chain, log)

	err := http.ListenAndServe(":"+port, nil)
	// log
	if err != nil {
		// handle error
		log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
	} else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: "+port+")"))
	}

}

func setConfig(chain string) {

	config := sdk.GetConfig()

	switch chain {
	case "iris":
		iris.ConfigureBech32Prefix()
	case "band":
		bech32MainPrefix := "band"
		var bip44CoinType uint32 = 494

		accountPrefix := bech32MainPrefix
		validatorPrefix := bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
		consensusPrefix := bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus

		config.SetBech32PrefixForAccount(accountPrefix, accountPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForValidator(validatorPrefix, validatorPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForConsensusNode(consensusPrefix, consensusPrefix+sdk.PrefixPublic)
		config.SetCoinType(bip44CoinType)

		//		fmt.Println(accountPrefix, validatorPrefix, consensusPrefix)

	case "terra":
		config.SetCoinType(terra.CoinType)
		config.SetFullFundraiserPath(terra.FullFundraiserPath)
		config.SetBech32PrefixForAccount(terra.Bech32PrefixAccAddr, terra.Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(terra.Bech32PrefixValAddr, terra.Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(terra.Bech32PrefixConsAddr, terra.Bech32PrefixConsPub)
		//		case "kava":
		//			kava.SetBech32AddressPrefixes(config)
		//			kava.SetBip44CoinType(config)
	case "emoney":
		emoney.ConfigureSDK()
	case "starname":
		Bech32Prefix := "star"
		Bech32PrefixAccAddr := Bech32Prefix
		Bech32PrefixAccPub := Bech32Prefix + sdk.PrefixPublic
		Bech32PrefixValAddr := Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
		Bech32PrefixValPub := Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
		Bech32PrefixConsAddr := Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
		Bech32PrefixConsPub := Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	case "certik":
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
		config.Seal()
	}

	config.Seal()
}
