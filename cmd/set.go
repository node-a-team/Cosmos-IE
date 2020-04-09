package cmd

import (
//        "fmt"

	rpc_cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos/getData/rpc"
        rest_cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos/getData/rest"

        rpc_terra "github.com/node-a-team/Cosmos-IE/chains/terra/getData/rpc"
        rest_terra "github.com/node-a-team/Cosmos-IE/chains/terra/getData/rest"

        rpc_iris "github.com/node-a-team/Cosmos-IE/chains/iris/getData/rpc"
        rest_iris "github.com/node-a-team/Cosmos-IE/chains/iris/getData/rest"

        rpc_kava "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rpc"
        rest_kava "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rest"

        rpc_iov "github.com/node-a-team/Cosmos-IE/chains/iov/getData/rpc"

        rpc_emoney "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rpc"
        rest_emoney "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rest"

)

func set_config() {

	switch chain {
        case "cosmos":
                rpc_cosmos.Addr = rpcAddr
                rest_cosmos.Addr = restAddr
                rest_cosmos.OperAddr = operAddr
        case "terra":
                rpc_terra.Addr = rpcAddr
                rest_terra.Addr = restAddr
                rest_terra.OperAddr = operAddr
        case "iris":
                rpc_iris.Addr = rpcAddr
                rest_iris.Addr = restAddr
                rest_iris.OperAddr = operAddr
        case "kava":
                rpc_kava.Addr = rpcAddr
                rest_kava.Addr = restAddr
                rest_kava.OperAddr = operAddr
        case "iov":
                rpc_iov.Addr = rpcAddr
                rpc_iov.ConsHexAddr = consHexAddr
        case "emoney":
                rpc_emoney.Addr = rpcAddr
                rest_emoney.Addr = restAddr
                rest_emoney.OperAddr = operAddr
        }


}
