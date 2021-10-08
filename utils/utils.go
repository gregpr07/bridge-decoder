package utils

import (
	"strings"
)

func MakeEtherscanLink(hash string) string {
	return strings.Join([]string{"https://etherscan.io/tx/", hash}, "")
}

// this function recursively goes through trace and checks if matching function returns true
// matching function just means the correct function call to certain (bridge) smart contract
func RecursivelyCheckTrace(txTrace TxTrace, checkFunc func(TxTrace, *TxTrace) bool, resultTrace *TxTrace) bool {
	if checkFunc(txTrace, resultTrace) {
		return true
	}
	for i := 0; i < len(txTrace.Calls); i++ {
		// txTrace.Calls[i]
		if RecursivelyCheckTrace(txTrace.Calls[i], checkFunc, resultTrace) {
			return true
		}
	}
	return false
}
