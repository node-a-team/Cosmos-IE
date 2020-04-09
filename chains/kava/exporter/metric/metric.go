package metric

import (
	"go.uber.org/zap"

	rest "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rest"
	rpc "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rpc"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)

var (
	metricData metric

	DenomList = []string{"ukava"}
	GaugesNamespaceList = [...]string{"blockHeight",
				"notBondedTokens",
				"bondedTokens",
				"totalSupply",
				"bondedRatio",
				"totalProposalCount",
				"votingProposalCount",
				"votingPower",
				"minSelfDelegation",
				"jailStatus",
				"proposerRanking",
				"proposerStatus",
				"delegationShares",
				"delegationRatio",
				"delegatorCount",
				"delegationSelf",
				"commissionRate",
				"commissionMaxRate",
				"commissionMaxChangeRate",
				"commitVoteType",
				"precommitStatus",
				"inflation",
				"actualInflation",
				}
)

type metric struct {

	Network struct {
		ChainID         string
		BlockHeight	int64
		PrecommitRate	float64

		Staking struct {
			NotBondedTokens	float64
			BondedTokens	float64
			TotalSupply	float64
			BondedRatio	float64
		}

		Minting struct {
			Inflation	float64
			ActualInflation	float64
		}

		Gov struct{
                        TotalProposalCount      float64
                        VotingProposalCount     float64
                }
	}

	Validator struct {
		Moniker			string
		VotingPower		float64
		MinSelfDelegation       float64
                JailStatus              float64



		Address struct {
			Account		string
			Operator	string
			ConsensusHex	string
		}
		Proposer struct {
			Ranking		float64
			Status		float64
		}

		Delegation struct {
			Shares		float64
			Ratio		float64
			DelegatorCount	float64
			Self		float64
		}

		Commission struct {
			Rate		float64
			MaxRate		float64
			MaxChangeRate	float64
		}

		Account struct {
			Balances	[]rest.Coin
			Commission	[]rest.Coin
			Rewards		[]rest.Coin
		}

		Commit struct {
			VoteType                float64
	                PrecommitStatus         float64
		}

	}
}



func SetMetric(currentBlock int64, restData *rest.RESTData, rpcData *rpc.RPCData, log *zap.Logger) {

	operAddr := rest.OperAddr
	consPubKey := restData.Validators.ConsPubKey
	consAddr := restData.Validatorsets[consPubKey][0]

	//// network

	metricData.Network.ChainID = rpcData.Commit.ChainId
        metricData.Network.BlockHeight = currentBlock

	metricData.Network.Staking.NotBondedTokens = utils.StringToFloat64(restData.StakingPool.Result.Not_bonded_tokens)
	metricData.Network.Staking.BondedTokens = utils.StringToFloat64(restData.StakingPool.Result.Bonded_tokens)
	metricData.Network.Staking.TotalSupply = restData.StakingPool.Result.Total_supply
	metricData.Network.Staking.BondedRatio = metricData.Network.Staking.BondedTokens / metricData.Network.Staking.TotalSupply

	// minting
	metricData.Network.Minting.Inflation = restData.Inflation
	metricData.Network.Minting.ActualInflation = metricData.Network.Minting.Inflation / metricData.Network.Staking.BondedRatio

	// gov
	metricData.Network.Gov.TotalProposalCount = restData.Gov.TotalProposalCount
        metricData.Network.Gov.VotingProposalCount = restData.Gov.VotingProposalCount


	//// validator
	metricData.Validator.Moniker = restData.Validators.Description.Moniker
        metricData.Validator.VotingPower = utils.StringToFloat64(restData.Validatorsets[consPubKey][1])
	metricData.Validator.MinSelfDelegation = utils.StringToFloat64(restData.Validators.MinSelfDelegation)
	metricData.Validator.JailStatus = utils.BoolToFloat64(restData.Validators.Jailed)

	// address
	metricData.Validator.Address.Operator = operAddr
	metricData.Validator.Address.Account = utils.GetAccAddrFromOperAddr(operAddr, log)
	metricData.Validator.Address.ConsensusHex = utils.Bech32AddrToHexAddr(consAddr, log)

	// proposer
	metricData.Validator.Proposer.Ranking = utils.StringToFloat64(restData.Validatorsets[consPubKey][3])
	metricData.Validator.Proposer.Status = rpcData.Commit.ValidatorProposingStatus

	// delegation
	metricData.Validator.Delegation.Shares = utils.StringToFloat64(restData.Validators.DelegatorShares)
	metricData.Validator.Delegation.Ratio = metricData.Validator.Delegation.Shares / metricData.Network.Staking.BondedTokens
	metricData.Validator.Delegation.DelegatorCount = restData.Delegations.DelegationCount
	metricData.Validator.Delegation.Self = restData.Delegations.SelfDelegation

	// commission
	metricData.Validator.Commission.Rate = utils.StringToFloat64(restData.Validators.Commission.Commission_rates.Rate)
	metricData.Validator.Commission.MaxRate = utils.StringToFloat64(restData.Validators.Commission.Commission_rates.Max_rate)
	metricData.Validator.Commission.MaxChangeRate = utils.StringToFloat64(restData.Validators.Commission.Commission_rates.Max_change_rate)

	// account
	metricData.Validator.Account.Balances = restData.Balances
	metricData.Validator.Account.Commission = restData.Commission
	metricData.Validator.Account.Rewards = restData.Rewards

	// commit
	metricData.Validator.Commit.VoteType = rpcData.Commit.VoteType
        metricData.Validator.Commit.PrecommitStatus = rpcData.Commit.ValidatorPrecommitStatus



}

func GetMetric() *metric {

	return &metricData
}

func GetDenomList() []string {
	return  DenomList
}
