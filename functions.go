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

var ERC20PredicateProxy = strings.ToLower("0x40ec5B33f54e0E8A33A975908C5BA1c14e5BbbDf")
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

func checkPolygonPOSDeposit (txHash string,txData TxData,txTrace TxTrace) error {
	var trace TxTrace;
	if (recursivelyCheckTrace(txTrace,checkPOSTrace,&trace)) {
		if strings.HasPrefix(trace.Input,DepositForMethodId) {
			// fmt.Println("Deposit",trace.TransactionHash)
		} else if strings.HasPrefix(trace.Input,ExitMethodId) {
			// fmt.Println("Exit",trace.TransactionHash,trace.Calls)
			fmt.Println(trace.TransactionHash)
			fmt.Println(trace.Calls[0].Calls[1].Calls[0].Calls[0])
		}

		// fmt.Println(trace.TransactionHash)
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