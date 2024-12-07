package models

type Transaction struct {
	Type                 string        `json:"type"`
	ChainID              string        `json:"chainId"`
	Nonce                string        `json:"nonce"`
	Gas                  string        `json:"gas"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	To                   string        `json:"to"`
	Value                string        `json:"value"`
	AccessList           []interface{} `json:"accessList"`
	Input                string        `json:"input"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	YParity              string        `json:"yParity"`
	V                    string        `json:"v"`
	Hash                 string        `json:"hash"`
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	TransactionIndex     string        `json:"transactionIndex"`
	From                 string        `json:"from"`
	GasPrice             string        `json:"gasPrice"`
}
