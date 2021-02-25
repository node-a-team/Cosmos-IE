package rest

import (
        "fmt"
	"encoding/json"
//	"sort"
//	"strconv"
	"go.uber.org/zap"
	"strings"
//	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

type validatorsets struct {
/*
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
*/
	Validators []struct {
		Address	string
		Pub_key	struct {
			Type	string	`json:"@type"`
			Key	string
		}
		Voting_power	string
	}
}

func getValidatorsets(currentBlockHeight int64, log *zap.Logger) map[string][]string {

	var vSets validatorsets
	var vSetsResult map[string][]string = make(map[string][]string)

	res, _ := runRESTCommand("/cosmos/base/tendermint/v1beta1/validatorsets/" +fmt.Sprint(currentBlockHeight) +"?pagination.limit=1000")
	json.Unmarshal(res, &vSets)

	// log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
                log.Info("", zap.Bool("Success", true), zap.String("Number of loaded validators", fmt.Sprint(len(vSets.Validators))),)
        }




//	fmt.Println("res", res)
//	fmt.Println("res", string(res))



	for _, value := range vSets.Validators {
		// address, voting_power, proposer_priority, proposer_ranking
		vSetsResult[value.Pub_key.Key] = []string{value.Address, value.Voting_power}



//		vSetsResult[value.ConsPubKey] = []string{value.ConsAddr, value.VotingPower, value.ProposerPriority, "0"}
	}

	return  vSetsResult
//	return  Sort(vSetsResult)
}
/*
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
*/
