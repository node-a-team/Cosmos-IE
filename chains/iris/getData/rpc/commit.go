package rpc

import (
	"fmt"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type commitInfo struct {
	ChainId                  string
	VoteType                 float64 // [0]: false, [1]: prevote, [2]: precommit
	ValidatorPrecommitStatus float64 // [0]: false, [1]: true
	ValidatorProposingStatus float64 // [0]: false, [1]: true
}

func getCommit(commitData *ctypes.ResultCommit, consHexAddr string) commitInfo {

	var cInfo commitInfo

	blockProposer := fmt.Sprint(commitData.SignedHeader.Header.ProposerAddress)

	cInfo.ChainId = commitData.SignedHeader.Header.ChainID
	cInfo.VoteType, cInfo.ValidatorPrecommitStatus, cInfo.ValidatorProposingStatus = 0.0, 0.0, 0.0

	for _, v := range commitData.SignedHeader.Commit.Precommits {

		func() {
			defer func() {

				if r := recover(); r != nil {
					// precommit failure validator
				}
			}()

			if consHexAddr == fmt.Sprint(v.ValidatorAddress) {

				cInfo.VoteType = float64(v.Type)

				if v.Type == 2 {
					cInfo.ValidatorPrecommitStatus = 1.0
				}
			}

			if consHexAddr == blockProposer {
				cInfo.ValidatorProposingStatus = 1.0
			}
		}()

	}

	return cInfo
}
