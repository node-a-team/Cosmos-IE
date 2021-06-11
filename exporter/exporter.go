package exporter

import (
	"fmt"
//	"time"
	"go.uber.org/zap"
	"strconv"

	rest "github.com/node-a-team/Cosmos-IE/rest/common"
	utils "github.com/node-a-team/Cosmos-IE/utils"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	defaultGauges []prometheus.Gauge
	additionalGauges []prometheus.Gauge

	gaugesDenom []prometheus.Gauge
)

func Start(chain string, log *zap.Logger) {

	denomList := getDenomList(chain)

	defaultGauges = make([]prometheus.Gauge, len(gaugesNamespaceList))
	gaugesDenom = make([]prometheus.Gauge, len(denomList)*3)

	// nomal guages
	for i := 0; i < len(gaugesNamespaceList); i++ {
                defaultGauges[i] = utils.NewGauge("exporter", gaugesNamespaceList[i], "")
                prometheus.MustRegister(defaultGauges[i])
        }

	// denom gagues
	count := 0
	for i := 0; i < len(denomList)*3; i += 3 {

		gaugesDenom[i] = utils.NewGauge("exporter_balances", denomList[count], "")
		gaugesDenom[i+1] = utils.NewGauge("exporter_commission", denomList[count], "")
		gaugesDenom[i+2] = utils.NewGauge("exporter_rewards", denomList[count], "")
		prometheus.MustRegister(gaugesDenom[i])
		prometheus.MustRegister(gaugesDenom[i+1])
		prometheus.MustRegister(gaugesDenom[i+2])

		count++
	}


	// labels
	labels := []string{"chainId", "moniker", "operatorAddress", "accountAddress", "consHexAddress"}
//	labels := []string{"chainId", "moniker", "operatorAddress", "accountAddress"}
	gaugesForLabel := utils.NewCounterVec("exporter", "labels", "", labels)

	prometheus.MustRegister(gaugesForLabel)


	for {
		func() {
/*			
			defer func() {

				if r := recover(); r != nil {
					//Error Log
				}

				time.Sleep(500 * time.Millisecond)

			}()
*/

			blockData := rest.GetBlocks(log)
                        currentBlockHeight, _:= strconv.ParseInt(blockData.Block.Header.Height, 10, 64)

			if previousBlockHeight != currentBlockHeight {

				fmt.Println("")
				log.Info("\t", zap.Bool("Success", true), zap.String("Block Height", fmt.Sprint(currentBlockHeight)))
				restData := rest.GetData(chain, currentBlockHeight, blockData, denomList[0], log)

				SetMetric(currentBlockHeight, restData, log)
				metricData := GetMetric()

				// balances, commission, rewards,
				count := 0
				for i := 0; i < len(denomList)*3; i +=3 {
					for _, value := range metricData.Validator.Account.Balances {
						if value.Denom == denomList[count] {
							gaugesDenom[i].Set(utils.StringToFloat64(value.Amount))
						}
					}
					for _, value := range metricData.Validator.Account.Commission {
                                                if value.Denom == denomList[count] {
							gaugesDenom[i+1].Set(utils.StringToFloat64(value.Amount))
						}
                                        }
					for _, value := range metricData.Validator.Account.Rewards {
                                                if value.Denom == denomList[count] {
							gaugesDenom[i+2].Set(utils.StringToFloat64(value.Amount))
						}
                                        }
					count++
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

//					metricData.Validator.Proposer.Ranking,
//					metricData.Validator.Proposer.Status,

					metricData.Validator.Delegation.Shares,
					metricData.Validator.Delegation.Ratio,
					metricData.Validator.Delegation.DelegatorCount,
					metricData.Validator.Delegation.Self,

					metricData.Validator.Commission.Rate,
					metricData.Validator.Commission.MaxRate,
					metricData.Validator.Commission.MaxChangeRate,
//					metricData.Validator.Commit.VoteType,
					metricData.Validator.Commit.PrecommitStatus,

					metricData.Network.Minting.Inflation,
					metricData.Network.Minting.ActualInflation,
				}

				for i:=0; i < len(gaugesNamespaceList); i++ {
					defaultGauges[i].Set(gaugesValue[i])
				}

				gaugesForLabel.WithLabelValues(metricData.Network.ChainID,
								metricData.Validator.Moniker,
								metricData.Validator.Address.Operator,
								metricData.Validator.Address.Account,
								metricData.Validator.Address.ConsensusHex,
				).Add(0)

				addGauges(chain, metricData, log)

			}

			previousBlockHeight = currentBlockHeight
		}()
	}
}

func addGauges(chain string, metricData *metric, log *zap.Logger) {

	if chain == "band" {
                if len(additionalGauges) == 0 {
                        additionalGauges = make([]prometheus.Gauge, len(gaugesNamespaceList_Band))

                        for i := 0; i < len(gaugesNamespaceList_Band); i++ {
                                additionalGauges[i] = utils.NewGauge("exporter", gaugesNamespaceList_Band[i], "")
                                prometheus.MustRegister(additionalGauges[i])
                        }
                } else {
                        gaugesValue := [...]float64{
				metricData.Validator.Oracle.Active,
                        }
                        for i:=0; i < len(gaugesNamespaceList_Band); i++ {
                                additionalGauges[i].Set(gaugesValue[i])
                        }
                }
        } else if chain == "terra" {
		if len(additionalGauges) == 0 {
			additionalGauges = make([]prometheus.Gauge, len(gaugesNamespaceList_Terra))

			for i := 0; i < len(gaugesNamespaceList_Terra); i++ {
                                additionalGauges[i] = utils.NewGauge("exporter", gaugesNamespaceList_Terra[i], "")
                                prometheus.MustRegister(additionalGauges[i])
			}
		} else {
			gaugesValue := [...]float64{
				metricData.Validator.Oracle.Miss,
				metricData.Validator.Oracle.FeederBalance,
		        }
			for i:=0; i < len(gaugesNamespaceList_Terra); i++ {
		                additionalGauges[i].Set(gaugesValue[i])
		        }
		}
	}
}
