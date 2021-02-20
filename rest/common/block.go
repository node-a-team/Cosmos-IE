package rest

import (
        "strings"
        "encoding/json"
        "go.uber.org/zap"
)

type Blocks struct {

    Block struct {
        Header struct {
                ChainID string `json:"chain_id"`
                Height  string
                Proposer_address string
        }

        Last_commit struct {
                Signatures []struct {
                        Block_id_flag           string
                        Validator_address       string
                }
        }
    }
}

func GetBlocks(log *zap.Logger) Blocks {


        var b Blocks

        res, _ := runRESTCommand("/blocks/latest")
        json.Unmarshal(res, &b)

        // log
        if strings.Contains(string(res), "not found") {
                // handle error
                log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res),))
        } else {
		// log.Info("Common Info", zap.Bool("Success", true), zap.String("Block Info", b),)
        }

        return b
}

