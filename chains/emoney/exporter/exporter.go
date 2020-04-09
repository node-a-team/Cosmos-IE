package exporter

import (
	"fmt"
	"time"
	"go.uber.org/zap"

	rpc "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rpc"
	rest "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rest"
	metric "github.com/node-a-team/Cosmos-IE/chains/emoney/exporter/metric"
	utils "github.com/node-a-team/Cosmos-IE/utils"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	previousBlockHeight	int64

)

func Start(log *zap.Logger) {

	gaugesNamespaceList := metric.GaugesNamespaceList

	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))
        var gaugesDenom []prometheus.Gauge = make([]prometheus.Gauge, len(metric.DenomList)*3) // wallet, rewards, commission


	// nomal guages
	for i := 0; i < len(gaugesNamespaceList); i++ {
                gauges[i] = utils.NewGauge("exporter", gaugesNamespaceList[i], "")
                prometheus.MustRegister(gauges[i])
        }


	// denom gagues
	count := 0
	for i := 0; i < len(metric.DenomList)*3; i += 3 {
		gaugesDenom[i] = utils.NewGauge("exporter_balances", metric.DenomList[count], "")
		gaugesDenom[i+1] = utils.NewGauge("exporter_commission", metric.DenomList[count], "")
		gaugesDenom[i+2] = utils.NewGauge("exporter_rewards", metric.DenomList[count], "")
		prometheus.MustRegister(gaugesDenom[i])
		prometheus.MustRegister(gaugesDenom[i+1])
		prometheus.MustRegister(gaugesDenom[i+2])

		count++
	}


	// labels
	labels := []string{"chainId", "moniker", "operatorAddress", "accountAddress", "consHexAddress"}
	gaugesForLabel := utils.NewCounterVec("exporter", "labels", "", labels)

	prometheus.MustRegister(gaugesForLabel)


	for {
		func() {
			defer func() {

				if r := recover(); r != nil {
					//Error Log
				}

				time.Sleep(500 * time.Millisecond)

			}()


			currentBlockHeight := rpc.BlockHeight()

			if previousBlockHeight != currentBlockHeight {

				fmt.Println("")
				log.Info("RPC-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Block Height: " +fmt.Sprint(currentBlockHeight)))


				restData, consHexAddr := rest.GetData(currentBlockHeight, log)
				rpcData := rpc.GetData(currentBlockHeight, consHexAddr, log)

				metric.SetMetric(currentBlockHeight, restData, rpcData, log)

				metricData := metric.GetMetric()
				denomList := metric.GetDenomList()


				count := 0
				for i := 0; i < len(denomList); i++ {

					for _, value := range metricData.Validator.Account.Balances {
						if value.Denom == denomList[i] {
							gaugesDenom[count].Set(utils.StringToFloat64(value.Amount))
							count++
						}
					}
					for _, value := range metricData.Validator.Account.Commission {
                                                if value.Denom == denomList[i] {
							gaugesDenom[count].Set(utils.StringToFloat64(value.Amount))
							count++
                                                }
                                        }
					for _, value := range metricData.Validator.Account.Rewards {
                                                if value.Denom == denomList[i] {
							gaugesDenom[count].Set(utils.StringToFloat64(value.Amount))
							count++
                                                }
                                        }
				}


				gaugesValue := [...]float64{
					float64(metricData.Network.BlockHeight),

					metricData.Network.Staking.NotBondedTokens,
					metricData.Network.Staking.BondedTokens,
					metricData.Network.Staking.TotalSupply,
					metricData.Network.Staking.BondedRatio,

					metricData.Network.Gov.TotalProposalCount,
					metricData.Network.Gov.VotingProposalCount,

					metricData.Validator.VotingPower,
					metricData.Validator.MinSelfDelegation,
					metricData.Validator.JailStatus,

					metricData.Validator.Proposer.Ranking,
					metricData.Validator.Proposer.Status,

					metricData.Validator.Delegation.Shares,
					metricData.Validator.Delegation.Ratio,
					metricData.Validator.Delegation.DelegatorCount,
					metricData.Validator.Delegation.Self,

					metricData.Validator.Commission.Rate,
					metricData.Validator.Commission.MaxRate,
					metricData.Validator.Commission.MaxChangeRate,
					metricData.Validator.Commit.VoteType,
					metricData.Validator.Commit.PrecommitStatus,
//					metricData.Validator.Oracle.Miss,
				}

				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}


				gaugesForLabel.WithLabelValues(metricData.Network.ChainID,
								metricData.Validator.Moniker,
								metricData.Validator.Address.Operator,
								metricData.Validator.Address.Account,
								metricData.Validator.Address.ConsensusHex,
				).Add(0)

			}

			previousBlockHeight = currentBlockHeight
		}()
	}
}


