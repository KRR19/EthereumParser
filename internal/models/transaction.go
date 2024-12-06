package models

type Transaction struct {
	Type                 string
	ChainID              string
	Nonce                string
	Gas                  string
	MaxFeePerGas         string
	MaxPriorityFeePerGas string
	To                   string
	Value                string
	AccessList           []string
	Input                string
	R                    string
	S                    string
	YParity              string
	V                    string
	Hash                 string
	BlockHash            string
	BlockNumber          string
	TransactionIndex     string
	From                 string
	GasPrice             string
}
