package metric

import (
	"go.uber.org/zap"

	rpc "github.com/node-a-team/Cosmos-IE/chains/iov/getData/rpc"
)

var (
	metricData metric

	DenomList = []string{"iov"}
	GaugesNamespaceList = [...]string{"blockHeight",
				"commitVoteType",
				"precommitStatus",
				}
)

type metric struct {

	Network struct {
		ChainID         string
		BlockHeight	int64
	}

	Validator struct {
		Commit struct {
			VoteType                float64
	                PrecommitStatus         float64
		}

	}
}



func SetMetric(currentBlock int64, rpcData *rpc.RPCData, log *zap.Logger) {

	//// network
	metricData.Network.ChainID = rpcData.Commit.ChainId
        metricData.Network.BlockHeight = currentBlock

	//// validator
	// commit
	metricData.Validator.Commit.VoteType = rpcData.Commit.VoteType
        metricData.Validator.Commit.PrecommitStatus = rpcData.Commit.ValidatorPrecommitStatus
}

func GetMetric() *metric {

	return &metricData
}

