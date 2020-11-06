package exporter

import (
	"go.uber.org/zap"

	rest "github.com/node-a-team/Cosmos-IE/rest/common"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)


func SetMetric(currentBlock int64, restData *rest.RESTData, log *zap.Logger) {

	operAddr := rest.OperAddr
	consPubKey := restData.Validators.ConsPubKey
	consAddr := restData.Validatorsets[consPubKey][0]

	//// network
	metricData.Network.ChainID = restData.Commit.ChainId
        metricData.Network.BlockHeight = currentBlock

	metricData.Network.Staking.NotBondedTokens = utils.StringToFloat64(restData.StakingPool.Result.Not_bonded_tokens)
	metricData.Network.Staking.BondedTokens = utils.StringToFloat64(restData.StakingPool.Result.Bonded_tokens)
	metricData.Network.Staking.TotalSupply = restData.StakingPool.Result.Total_supply
	metricData.Network.Staking.BondedRatio = metricData.Network.Staking.BondedTokens / metricData.Network.Staking.TotalSupply

	// minting
	metricData.Network.Minting.Inflation = restData.Inflation
	metricData.Network.Minting.ActualInflation = metricData.Network.Minting.Inflation / metricData.Network.Staking.BondedRatio

	//gov
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
	metricData.Validator.Proposer.Status = restData.Commit.ValidatorProposingStatus

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
//	metricData.Validator.Commit.VoteType = restData.Commit.VoteType
        metricData.Validator.Commit.PrecommitStatus = restData.Commit.ValidatorPrecommitStatus

	// oracle
	metricData.Validator.Oracle.Miss = restData.Oracle.Miss
	//for Terra
	metricData.Validator.Oracle.FeederBalance = utils.StringToFloat64(restData.Oracle.Feeder.Balance.Amount)

}

func GetMetric() *metric {

	return &metricData
}

