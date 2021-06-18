package exporter

import (
	"go.uber.org/zap"

	rest "github.com/node-a-team/Cosmos-IE/rest/common"
	utils "github.com/node-a-team/Cosmos-IE/utils"
)


func SetMetric(currentBlock int64, restData *rest.RESTData, log *zap.Logger) {

	operAddr := rest.OperAddr
	consPubKey := restData.Validator.Consensus_pubkey.Key
	consAddr := restData.Validatorsets[consPubKey][0]

	/* network */
	metricData.Network.ChainID = restData.Commit.ChainId
        metricData.Network.BlockHeight = currentBlock

	metricData.Network.Staking.NotBondedTokens = utils.StringToFloat64(restData.StakingPool.Pool.Not_bonded_tokens)
	metricData.Network.Staking.BondedTokens = utils.StringToFloat64(restData.StakingPool.Pool.Bonded_tokens)
	metricData.Network.Staking.TotalSupply = restData.StakingPool.Pool.Total_supply
	metricData.Network.Staking.BondedRatio = metricData.Network.Staking.BondedTokens / metricData.Network.Staking.TotalSupply

	//// minting
	metricData.Network.Minting.Inflation = restData.Inflation
	metricData.Network.Minting.ActualInflation = metricData.Network.Minting.Inflation / metricData.Network.Staking.BondedRatio

	////gov
	metricData.Network.Gov.TotalProposalCount = restData.Gov.TotalProposalCount
        metricData.Network.Gov.VotingProposalCount = restData.Gov.VotingProposalCount


	/* validator */
	metricData.Validator.Moniker = restData.Validator.Description.Moniker
        metricData.Validator.VotingPower = utils.StringToFloat64(restData.Validatorsets[consPubKey][1])
	metricData.Validator.MinSelfDelegation = utils.StringToFloat64(restData.Validator.MinSelfDelegation)
	metricData.Validator.JailStatus = utils.BoolToFloat64(restData.Validator.Jailed)

	//// address
	metricData.Validator.Address.Operator = operAddr
	metricData.Validator.Address.Account = utils.GetAccAddrFromOperAddr(operAddr, log)
	metricData.Validator.Address.ConsensusHex = utils.Bech32AddrToHexAddr(consAddr, log)

	//// proposer
//	metricData.Validator.Proposer.Ranking = utils.StringToFloat64(restData.Validatorsets[consPubKey][3])
//	metricData.Validator.Proposer.Status = restData.Commit.ValidatorProposingStatus

	//// delegation
	metricData.Validator.Delegation.Shares = utils.StringToFloat64(restData.Validator.DelegatorShares)
	metricData.Validator.Delegation.Ratio = metricData.Validator.Delegation.Shares / metricData.Network.Staking.BondedTokens
	metricData.Validator.Delegation.DelegatorCount = restData.Delegations.DelegationCount
	metricData.Validator.Delegation.Self = restData.Delegations.SelfDelegation

	//// commission
	metricData.Validator.Commission.Rate = utils.StringToFloat64(restData.Validator.Commission.Commission_rates.Rate)
	metricData.Validator.Commission.MaxRate = utils.StringToFloat64(restData.Validator.Commission.Commission_rates.Max_rate)
	metricData.Validator.Commission.MaxChangeRate = utils.StringToFloat64(restData.Validator.Commission.Commission_rates.Max_change_rate)

	//// account
	metricData.Validator.Account.Balances = restData.Balances
	metricData.Validator.Account.Commission = restData.Commission
	metricData.Validator.Account.Rewards = restData.Rewards

	//// commit
//	metricData.Validator.Commit.VoteType = restData.Commit.VoteType
        metricData.Validator.Commit.PrecommitStatus = restData.Commit.ValidatorPrecommitStatus

	//// oracle
	// Terra
	metricData.Validator.Oracle.Miss = restData.Oracle_terra
	// Band
	metricData.Validator.Oracle.Active = restData.Oracle_band

}

func GetMetric() *metric {

	return &metricData
}

