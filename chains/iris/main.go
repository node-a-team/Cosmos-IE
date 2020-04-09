package iris

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/node-a-team/cosmos_metric/chains/iris/exporter"
	rpc "github.com/node-a-team/cosmos_metric/chains/iris/getData/rpc"
)

func Main(port string) {

	log, _ := zap.NewDevelopment()
	defer log.Sync()

//	cfg.ConfigPath = os.Args[1]

//	port := cfg.Init()
	rpc.OpenSocket(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":"+port, nil)

	// log
	if err != nil {
		// handle error
		log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
	} else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: "+port+")"))
	}

}
