package rest

import (
        "fmt"
	"encoding/json"
	"sort"
	"strconv"
	"go.uber.org/zap"
	"strings"
)

type validatorsets struct {

	Height	string	`json:"height"`

	Result	struct {
		Block_Height	string	`json:"block_height"`
		Validators	[]struct {
			ConsAddr			string	`json:"address"`
			ConsPubKey			string	`json:"pub_key"`
			ProposerPriority	string	`json:"proposer_priority"`
			VotingPower		string	`json:"voting_power"`
		}

	}
}

func getValidatorsets(currentBlockHeight int64, log *zap.Logger) map[string][]string {

	var vSets validatorsets
	var vSetsResult map[string][]string = make(map[string][]string)

	res, _ := runRESTCommand("/validatorsets/" +fmt.Sprint(currentBlockHeight))
	json.Unmarshal(res, &vSets)

	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("REST-Server", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("REST-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Validatorsets"),)
        }



	for _, value := range vSets.Result.Validators {
		// address, voting_power, proposer_priority, proposer_ranking
		vSetsResult[value.ConsPubKey] = []string{value.ConsAddr, value.VotingPower, value.ProposerPriority, "0"}
	}

	return  Sort(vSetsResult)
}

func Sort(mapValue map[string][]string) map[string][]string {

	keys := []string{}
	newMapValue := mapValue

	for key := range mapValue {
		keys = append(keys, key)
	}

	// Sort by proposer_priority
	sort.Slice(keys, func(i, j int) bool {
		a, _ := strconv.Atoi(mapValue[keys[i]][2])
		b, _ := strconv.Atoi(mapValue[keys[j]][2])
		return a > b
	})

	for i, key := range keys {
		// proposer_ranking
		newMapValue[key][3] = strconv.Itoa(i + 1)
	}
	return newMapValue
}
