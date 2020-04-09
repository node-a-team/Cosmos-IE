package rpc

import (
	"fmt"
	"go.uber.org/zap"

        tmclient "github.com/tendermint/tendermint/rpc/client"
)


type RPCData struct {
        Commit	commitInfo
}

var (
        Addr string

        Client *tmclient.HTTP
)

func newRPCData() *RPCData {

        rd := &RPCData {
		//
        }

        return rd
}

func GetData(blockHeight int64, consHexAddr string, log *zap.Logger) *RPCData {

	rd := newRPCData()

	var commitHeight int64 = blockHeight -1

	commitData, err := Client.Commit(&commitHeight)
	// log
        if err != nil {
                // handle error
                log.Fatal("RPC-Server", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
                log.Info("RPC-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Get Data", "Commit Data"),)
        }

	rd.Commit = getCommit(commitData, consHexAddr)


        return rd
}

func OpenSocket(log *zap.Logger) {

        Client = tmclient.NewHTTP("tcp://"+Addr, "/websocket")

        err := Client.Start()

	if err != nil {
                // handle error
                log.Fatal("RPC-Server", zap.Bool("Success", false), zap.String("err", fmt.Sprintf("%s", err)),)
        } else {
                log.Info("RPC-Server", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Open Socket", "tcp://" +Addr +"/websocket"),)
        }


        defer Client.Stop()

}

func BlockHeight() (res int64) {

	info, _:= Client.ABCIInfo()

	res = info.Response.LastBlockHeight

	return res
}


