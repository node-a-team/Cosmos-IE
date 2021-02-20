package cmd

import (
//        cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos/getData/rest"
//        iris "github.com/node-a-team/Cosmos-IE/chains/iris/getData/rest"
//        kava "github.com/node-a-team/Cosmos-IE/chains/kava/getData/rest"
//    	  emoney "github.com/node-a-team/Cosmos-IE/chains/emoney/getData/rest"
//        band "github.com/node-a-team/Cosmos-IE/chains/bandprotocol/getData/rest"

//	  terra "github.com/node-a-team/Cosmos-IE/rest/chains/terra"

        common "github.com/node-a-team/Cosmos-IE/rest/common"

//	rpc_iov "github.com/node-a-team/Cosmos-IE/chains/iov/getData/rpc"
)

func set_config() {

	common.Addr = restAddr
        common.OperAddr = operAddr

//	switch chain {
//        case "cosmos":
//                cosmos.Addr = restAddr
//                cosmos.OperAddr = operAddr
/*        case "terra":
		terra.FeeDenom = feeDenom
*/
/*	case "iris":
                iris.Addr = restAddr
                iris.OperAddr = operAddr
*/
/*        case "kava":
                kava.Addr = restAddr
                kava.OperAddr = operAddr
*/
/*        case "iov":
                rpc_iov.Addr = rpcAddr
                rpc_iov.ConsHexAddr = consHexAddr
*/
/*        case "emoney":
                .Addr = restAddr
                emoney.OperAddr = operAddr
*/
/*	case "band":
                band.Addr = restAddr
                band.OperAddr = operAddr
*/
//        }
}
