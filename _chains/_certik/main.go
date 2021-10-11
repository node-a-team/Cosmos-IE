package certik

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/node-a-team/Cosmos-IE/chains/certik/exporter"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	rpc "github.com/node-a-team/Cosmos-IE/chains/certik/getData/rpc"
)

func Main(port string) {

	log, _ := zap.NewDevelopment()
	defer log.Sync()

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
