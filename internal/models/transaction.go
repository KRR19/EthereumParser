package models

type JsonRpcResponse struct {
	JsonRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Result  `json:"result"`
}

type Result struct {
	Hash                    string         `json:"hash"`
	ParentHash              string         `json:"parentHash"`
	Sha3Uncles              string         `json:"sha3Uncles"`
	Miner                   string         `json:"miner"`
	StateRoot               string         `json:"stateRoot"`
	TransactionsRoot        string         `json:"transactionsRoot"`
	ReceiptsRoot            string         `json:"receiptsRoot"`
	LogsBloom               string         `json:"logsBloom"`
	Difficulty              string         `json:"difficulty"`
	Number                  string         `json:"number"`
	GasLimit                string         `json:"gasLimit"`
	GasUsed                 string         `json:"gasUsed"`
	Timestamp               string         `json:"timestamp"`
	ExtraData               string         `json:"extraData"`
	MixHash                 string         `json:"mixHash"`
	Nonce                   string         `json:"nonce"`
	BaseFeePerGas           string         `json:"baseFeePerGas"`
	WithdrawalsRoot         string         `json:"withdrawalsRoot"`
	BlobGasUsed             string         `json:"blobGasUsed"`
	ExcessBlobGas           string         `json:"excessBlobGas"`
	ParentBeaconBlockRoot   string         `json:"parentBeaconBlockRoot"`
	TotalDifficulty         string         `json:"totalDifficulty"`
	Size                    string         `json:"size"`
	Uncles                  []string       `json:"uncles"`
	Transactions            []Transaction  `json:"transactions"`
	Withdrawals             []Withdrawal   `json:"withdrawals"`
}

type Transaction struct {
	Type                  string   `json:"type"`
	ChainID               string   `json:"chainId"`
	Nonce                 string   `json:"nonce"`
	Gas                   string   `json:"gas"`
	MaxFeePerGas          string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas  string   `json:"maxPriorityFeePerGas"`
	To                    string   `json:"to"`
	Value                 string   `json:"value"`
	AccessList            []string `json:"accessList"`
	Input                 string   `json:"input"`
	R                     string   `json:"r"`
	S                     string   `json:"s"`
	YParity               string   `json:"yParity"`
	V                     string   `json:"v"`
	Hash                  string   `json:"hash"`
	BlockHash             string   `json:"blockHash"`
	BlockNumber           string   `json:"blockNumber"`
	TransactionIndex      string   `json:"transactionIndex"`
	From                  string   `json:"from"`
	GasPrice              string   `json:"gasPrice"`
}

type Withdrawal struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}
