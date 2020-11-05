package rest

import (
	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"
)

type govInfo struct {
	TotalProposalCount	float64
	VotingProposalCount	float64
}


type gov struct {
	Height	string
	Result	[]proposal
}

type proposal struct {
	Content struct {
		Type			string
		Value struct {
			Title		string
			Description	string
		}
	}

	Id		string
	Proposal_status	string
	Final_tally_result struct {
		Yes		string
		Abstain		string
		No		string
		No_with_veto	string
	}

	Submit_time		string
	Deposit_end_time	string
	Total_deposit		string
	Voting_start_time	string
	Voting_end_time		string

}

func getGovInfo(log *zap.Logger) govInfo {

	var g gov
	var gi govInfo

	votingCount := 0

        res, _ := runRESTCommand("/gov/proposals")
        json.Unmarshal(res, &g)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("\t", zap.Bool("Success", true), zap.String("Total Proposal Count", fmt.Sprint(len(g.Result))),)
        }

	for _, value := range g.Result {
		if value.Proposal_status == "VotingPeriod" {
			votingCount++
		}
	}

	gi.TotalProposalCount = float64(len(g.Result))
	gi.VotingProposalCount = float64(votingCount)

	return gi
}
