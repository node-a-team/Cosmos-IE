package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"

	utils "github.com/node-a-team/cosmos_metric/utils"
)

type param struct {
	paramResult
}

type paramResult struct {
	Type  string
	Value struct {
		Inflation string

		Asset_tax_rate          string
		Issue_token_base_fee    Coin
		Mint_token_fee_ratio    string
		Create_gateway_base_fee Coin
		Gateway_asset_fee_ratio string

		Fee string

		Gas_price_threshold string
		Tx_size             string

		Unbonding_time string
		Max_validators int64

		Community_tax         string
		Base_proposer_reward  string
		Bonus_proposer_reward string

		Max_evidence_age           string
		Signed_blocks_window       string
		Min_signed_per_window      string
		Double_sign_jail_duration  string
		Downtime_jail_duration     string
		Censorship_jail_duration   string
		Slash_fraction_double_sign string
		Slash_fraction_downtime    string
		Slash_fraction_censorship  string

		Max_request_timeout    string
		Min_deposit_multiple   string
		Service_fee_tax        string
		Slash_fraction         string
		Complaint_retrospect   string
		Arbitration_time_limit string
		Tx_size_limit          string

		Critical_deposit_period  string
		Critical_min_deposit     []Coin
		Critical_voting_period   string
		Critical_max_num         string
		Critical_threshold       string
		Critical_veto            string
		Critical_participation   string
		Critical_penalty         string
		Important_deposit_period string
		Important_min_deposit    []Coin
		Important_voting_period  string
		Important_max_num        string
		Important_threshold      string
		Important_veto           string
		Important_participation  string
		Important_penalty        string
		Normal_deposit_period    string
		Normal_min_deposit       []Coin
		Normal_voting_period     string
		Normal_max_num           string
		Normal_threshold         string
		Normal_veto              string
		Normal_participation     string
		Normal_penalty           string
		System_halt_period       string
	}
}

func getParam(log *zap.Logger) *[]param {

	var p []param

	res, _ := runRESTCommand("/params")
	json.Unmarshal(res, &p)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Inflation"))
	}

	return &p
}

func getInflation(log *zap.Logger) float64 {

	p := getParam(log)

	for _, value := range *p {
		if value.Type == "irishub/mint/Params" {
			return utils.StringToFloat64(value.Value.Inflation)
		}
	}

	return 0.0
}
