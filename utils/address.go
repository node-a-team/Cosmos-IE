package utils

import (
        "fmt"
        "go.uber.org/zap"

        "github.com/tendermint/tendermint/libs/bech32"

        sdk "github.com/cosmos/cosmos-sdk/types"

)

// Bech32 Addr -> Hex Addr
func Bech32AddrToHexAddr(bech32str string, log *zap.Logger) string {
        _, bz, err := bech32.DecodeAndConvert(bech32str)
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "Bech32Addr To HexAddr"),)
        }

        return fmt.Sprintf("%X", bz)
}

func GetAccAddrFromOperAddr(operAddr string, log *zap.Logger) string {

        // Get HexAddress
        hexAddr, err := sdk.ValAddressFromBech32(operAddr)
        // log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "OperAddr To HexAddr"),)
        }

        accAddr, err := sdk.AccAddressFromHex(fmt.Sprint(hexAddr))
        // log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "HexAddr To AccAddr"),)
        }

        return accAddr.String()
}


func GetAccAddrFromOperAddr_localPrefixes(operAddr string, log *zap.Logger, bech32Prefixes []string) string {

        // Get HexAddress
        bz, err := sdk.GetFromBech32(operAddr, bech32Prefixes[2])
        // hexAddr := fmt.Sprintf("%X", bz)

        // log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
        } else {
                //                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "OperAddr To HexAddr"),)
        }

        accAddr, err := bech32.ConvertAndEncode(bech32Prefixes[0], bz)

        // log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
        } else {
                //                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "HexAddr To AccAddr"),)
        }

        return accAddr
}

