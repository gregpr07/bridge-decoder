package bridges

type DecodedTransfer struct {
	method string
	to     string
	amount string
}

func DecodeTransferInput(inpt string) (DecodedTransfer, error) {
	return DecodedTransfer{"0x" + inpt[0:10], "0x" + inpt[74-40:74], "0x" + inpt[74:138]}, nil
}

type DecodedDepositFor struct {
	method    string
	user      string
	rootToken string
	amount    string
}

func DecodeDepositFor(inpt string) (DecodedDepositFor, error) {
	return DecodedDepositFor{"0x" + inpt[0:10], "0x" + inpt[74-40:74], "0x" + inpt[138-40:138], "0x" + inpt[len(inpt)-64:]}, nil
}

type DecodedDepositEtherFor struct {
	method string
	to     string
}

func DecodeDepositEtherFor(inpt string) (DecodedDepositEtherFor, error) {
	return DecodedDepositEtherFor{"0x" + inpt[0:10], "0x" + inpt[74-40:74]}, nil
}
