package bridges

import (
	"fmt"
	"strings"

	"github.com/gregpr07/bridge-decoder/utils"
	"gorm.io/gorm"
)

// var ERC20PredicateProxy = strings.ToLower("0x40ec5B33f54e0E8A33A975908C5BA1c14e5BbbDf")
var WithdrawMethodId = strings.ToLower("0x2e1a7d4d")
var chain = 137

func checkTrace(trace utils.TxTrace, resultTrace *utils.TxTrace) bool {
	if strings.HasPrefix(trace.Input, WithdrawMethodId) {
		*resultTrace = trace
		return true
	}
	return false
}

func CheckPolygonPOS_Polygon(db *gorm.DB, txHash string, txTrace utils.TxTrace) error {
	var trace utils.TxTrace
	if utils.RecursivelyCheckTrace(txTrace, checkTrace, &trace) {
		if strings.HasPrefix(trace.Input, DepositEtherForMethodId) {
			//? this is deposit of ether

			// To test if I can get internal calls as well
			// if txData.To != RootChainManagerProxy {
			// 	fmt.Println("\n\nNOT TO CHAIN MANAGER\n\n")
			// }

			valid := false
			if len(trace.Calls) > 0 {
				if len(trace.Calls[0].Calls) > 1 {
					valid = true
				}
			}

			decoded, _ := DecodeDepositEtherFor(trace.Input)

			var amount string
			if valid {
				transferTrace := trace.Calls[0].Calls[2]
				amount = transferTrace.Value
			} else {
				amount = "0x"
			}

			// fmt.Println("Deposit ether", transferTrace.Value, "eth", "TO", decoded.to, utils.MakeEtherscanLink(trace.TransactionHash))
			fmt.Println("Deposit ether")
			depositedEther := PolygonPOSBridgeTx{
				Chain:           chain,
				Successful:      valid,
				Type:            "Deposit",
				Method:          "DepositEtherFor",
				OriginFrom:      trace.From,
				OriginTo:        trace.To,
				Bridge:          trace.To,
				From:            trace.From,
				To:              decoded.to,
				Amount:          amount,
				BaseToken:       utils.ZERO_ADDRESS,
				Input:           trace.Input,
				GasUsed:         trace.GasUsed,
				BlockNumber:     trace.BlockNumber,
				TransactionHash: trace.TransactionHash,
			}
			dbResult := db.Create(&depositedEther)
			if dbResult.Error != nil {
				fmt.Println("Error creating deposit ether")
			}

		} else if strings.HasPrefix(trace.Input, DepositForMethodId) {
			//? this is deposit of erc20

			valid := false
			if len(trace.Calls) > 0 {
				if len(trace.Calls[0].Calls) > 0 {
					if len(trace.Calls[0].Calls[0].Calls) > 0 {
						if len(trace.Calls[0].Calls[0].Calls[0].Calls) > 0 {
							valid = true
						}
					}
				}
			}

			// if !valid {
			// 	transferTrace := trace.Calls[0].Calls[0].Calls[0].Calls[0]

			// } else {
			// 	//fmt.Println("DepositFor failed (probably).", utils.MakeEtherscanLink(trace.TransactionHash))
			// }

			decoded, _ := DecodeDepositFor(trace.Input)

			// fmt.Println("Deposit erc20", decoded.amount, decoded.rootToken, utils.MakeEtherscanLink(trace.TransactionHash))
			fmt.Println("Deposit ERC20")
			depositedEther := PolygonPOSBridgeTx{
				Chain:           chain,
				Successful:      valid,
				Type:            "Deposit",
				Method:          "DepositFor",
				OriginFrom:      trace.From,
				OriginTo:        trace.To,
				Bridge:          trace.To,
				To:              decoded.user,
				Amount:          decoded.amount,
				BaseToken:       decoded.rootToken,
				Input:           trace.Input,
				GasUsed:         trace.GasUsed,
				BlockNumber:     trace.BlockNumber,
				TransactionHash: trace.TransactionHash,
			}
			dbResult := db.Create(&depositedEther)
			if dbResult.Error != nil {
				fmt.Println("Error creating deposit ether")
			}

		} else if strings.HasPrefix(trace.Input, WithdrawMethodId) {
			//? withdraw, burn of erc20 or ether
			// fmt.Println("Exit",trace.TransactionHash,trace.Calls)
			// fmt.Println(trace.TransactionHash)

			// example of exit is https://etherscan.io/tx/0x2e86f1b55a8f750fc15d4fd58622337530fecbd70265902a7afc3a3134b92a63#internal
			// this is the trace that gets the amount deposited

			fmt.Println("hash: ", txHash)
			return nil

			valid := false
			if len(trace.Calls) > 0 {
				if len(trace.Calls[0].Calls) > 0 {
					if len(trace.Calls[0].Calls[1].Calls) > 0 {
						if len(trace.Calls[0].Calls[1].Calls[0].Calls) > 0 {
							valid = true
						}
					}
				}
			}
			if !valid {
				fmt.Println("Exit failed (probably).", utils.MakeEtherscanLink(trace.TransactionHash))
				return nil
			}

			transferTrace := trace.Calls[0].Calls[1].Calls[0].Calls[0]

			depositedEther := PolygonPOSBridgeTx{
				Chain:      chain,
				Successful: valid,
				Type:       "Exit",
				Method:     "",
				OriginFrom: trace.From,
				OriginTo:   trace.To,
				Bridge:     trace.To,
				//? probably redundant but still
				From:            transferTrace.To,
				To:              "",
				Amount:          "",
				BaseToken:       "",
				Input:           trace.Input,
				GasUsed:         trace.GasUsed,
				BlockNumber:     trace.BlockNumber,
				TransactionHash: trace.TransactionHash,
			}
			if transferTrace.Input == "0x" {
				// fmt.Println("Ether exit",transferTrace.Value,"ETH", utils.MakeEtherscanLink(transferTrace.TransactionHash))
				fmt.Println("Ether Exit")
				depositedEther.Method = "ExitEther"
				depositedEther.To = transferTrace.To
				depositedEther.Amount = transferTrace.Value
				depositedEther.BaseToken = utils.ZERO_ADDRESS
			} else {
				// fmt.Println(trace.Calls[0].Calls[1].Calls[0].Calls[0])
				decoded, _ := DecodeTransferInput(transferTrace.Input)
				fmt.Println(decoded.to)
				// fmt.Println("ERC20 exit", decoded.amount, transferTrace.To, utils.MakeEtherscanLink(transferTrace.TransactionHash))
				depositedEther.Method = "ExitERC20"
				depositedEther.To = decoded.to
				depositedEther.Amount = decoded.amount
				depositedEther.BaseToken = transferTrace.To
			}

			dbResult := db.Create(&depositedEther)
			if dbResult.Error != nil {
				fmt.Println("Error creating deposit ether")
			}
		}
	}
	return nil
}
