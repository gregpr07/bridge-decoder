package utils

type TxData struct {
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Input       string `json:"input"`
	Nonce       string `json:"nonce"`
	PublicKey   string `json:"publicKey"`
	Value       string `json:"value"`
}

type TxTrace struct {
	Type            string `json:"type"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	Gas             string `json:"gas"`
	Input           string `json:"input"`
	Output          string `json:"output"`
	GasUsed         string `json:"gasUsed"`
	BlockNumber     int    `json:"blockNumber"`
	TransactionHash string `json:"transactionHash"`
	Calls           []TxTrace
}

const SCHEMA = "bridges"
const ZERO_ADDRESS = "0x"
