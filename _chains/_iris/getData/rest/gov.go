package rest

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
)

type govInfo struct {
	TotalProposalCount  float64
	VotingProposalCount float64
}

type gov struct {
	proposal
}

type proposal struct {
	Type  string
	Value struct {
		BasicProposal struct {
			Proposal_id     string
			Proposal_status string
		}
	}
}

func getGovInfo(log *zap.Logger) govInfo {

	var g []gov
	var gi govInfo

	votingCount := 0

	res, _ := runRESTCommand("/gov/proposals")
	json.Unmarshal(res, &g)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Governance"))
	}

	for _, value := range g {
		if value.Value.BasicProposal.Proposal_status == "VotingPeriod" {
			votingCount++
		}
	}

	gi.TotalProposalCount = float64(len(g))
	gi.VotingProposalCount = float64(votingCount)

	return gi
}
