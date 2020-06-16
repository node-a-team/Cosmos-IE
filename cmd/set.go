package cmd

import (
        rest_cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos/getData/rest"
        rest_terra "github.com/node-a-team/Cosmos-IE/chains/terra/getData/rest"
        rest_iris "github.com/node-a-team/Cosmos-IE/chains/iris/getData/rest"
        rest_kava "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rest"
	rest_emoney "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rest"
        rest_band "github.com/node-a-team/Cosmos-IE/chains/bandprotocol/getData/rest"

	rpc_iov "github.com/node-a-team/Cosmos-IE/chains/iov/getData/rpc"
)

func set_config() {

	switch chain {
        case "cosmos":
                rest_cosmos.Addr = restAddr
                rest_cosmos.OperAddr = operAddr
        case "terra":
                rest_terra.Addr = restAddr
                rest_terra.OperAddr = operAddr
        case "iris":
                rest_iris.Addr = restAddr
                rest_iris.OperAddr = operAddr
        case "kava":
                rest_kava.Addr = restAddr
                rest_kava.OperAddr = operAddr
        case "iov":
                rpc_iov.Addr = rpcAddr
                rpc_iov.ConsHexAddr = consHexAddr
        case "emoney":
                rest_emoney.Addr = restAddr
                rest_emoney.OperAddr = operAddr
	case "band":
                rest_band.Addr = restAddr
                rest_band.OperAddr = operAddr
        }
}
