package main

import (
	"fmt"
	"strings"
)

func recursivelyCheckTrace (txTrace TxTrace, checkFunc func(TxTrace,*TxTrace) bool,resultTrace *TxTrace) bool {
	if (checkFunc(txTrace,resultTrace)) {return true}
	for i := 0; i < len(txTrace.Calls); i++ {
		// txTrace.Calls[i]
		if recursivelyCheckTrace(txTrace.Calls[i],checkFunc,resultTrace) {return true}
	}
	return false
}

// var ERC20PredicateProxy = strings.ToLower("0x40ec5B33f54e0E8A33A975908C5BA1c14e5BbbDf")
var RootChainManagerProxy = strings.ToLower("0xA0c68C638235ee32657e8f720a23ceC1bFc77C77")
var DepositForMethodId = strings.ToLower("0xe3dec8fb")
var ExitMethodId = strings.ToLower("0x3805550f")
func checkPOSTrace (trace TxTrace,resultTrace *TxTrace) bool {
	if (trace.To == RootChainManagerProxy && 
		(strings.HasPrefix(trace.Input,DepositForMethodId) || strings.HasPrefix(trace.Input,ExitMethodId))) {
		*resultTrace = trace
		return true
	}
	return false
}

type decodedTransfer struct {
	method string;
	to string
	amount string
}
func decodeTransferInput(inpt string) decodedTransfer {
	return decodedTransfer{"0x"+inpt[0:10],"0x"+inpt[10:74],"0x"+inpt[74:138]}
}

func checkPolygonPOSDeposit (txHash string,txData TxData,txTrace TxTrace) error {
	var trace TxTrace;
	if (recursivelyCheckTrace(txTrace,checkPOSTrace,&trace)) {
		if strings.HasPrefix(trace.Input,DepositForMethodId) {
			// fmt.Println("Deposit",trace.TransactionHash)
		} else if strings.HasPrefix(trace.Input,ExitMethodId) {
			// fmt.Println("Exit",trace.TransactionHash,trace.Calls)
			// fmt.Println(trace.TransactionHash)

			// example of exit is https://etherscan.io/tx/0x2e86f1b55a8f750fc15d4fd58622337530fecbd70265902a7afc3a3134b92a63#internal
			// this is the trace that gets the amount deposited

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
				fmt.Println("Transaction failed (probably).")
				return nil
			}
			
			swapTrace := trace.Calls[0].Calls[1].Calls[0].Calls[0]
			
			if (swapTrace.Input == "0x") {
				fmt.Println("Ether exit",swapTrace.Value,"ETH", swapTrace.TransactionHash)
			} else {
				decoded := decodeTransferInput(swapTrace.Input)
				fmt.Println("ERC20 exit", decoded.amount, swapTrace.To, swapTrace.TransactionHash)
			}
		}
	}
	return nil
}

func CheckBridge(txHash string,txData TxData,txTrace TxTrace) {
	// fmt.Println(txHash)

	err := checkPolygonPOSDeposit(txHash,txData,txTrace)
	if err != nil {
		fmt.Println(err)
	}

}