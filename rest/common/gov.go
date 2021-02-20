package rest

import (
//	"fmt"
	"strings"
	"go.uber.org/zap"
	"encoding/json"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

type govInfo struct {
	TotalProposalCount	float64
	VotingProposalCount	float64
}


type gov struct {
//	Height	stringa
	Proposals	[]proposal
	Pagination	struct {
		Total	string
	}
}

type proposal struct {
	Status	string
}

func getGovInfo(log *zap.Logger) govInfo {

	var g gov
	var gi govInfo

	votingCount := 0

        res, _ := runRESTCommand("/cosmos/gov/v1beta1/proposals")
        json.Unmarshal(res, &g)
	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("\t", zap.Bool("Success", true), zap.String("Total Proposal Count", g.Pagination.Total),)
        }

	for _, value := range g.Proposals {
		if value.Status == "PROPOSAL_STATUS_VOTING_PERIOD" {
			votingCount++
		}
	}

	gi.TotalProposalCount = utils.StringToFloat64(g.Pagination.Total)
	gi.VotingProposalCount = float64(votingCount)

	return gi
}
