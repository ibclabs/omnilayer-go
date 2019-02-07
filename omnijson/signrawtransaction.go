package omnijson

type SignRawTransactionResult = string

type SignRawTransactionCommand struct {
	Hex      string
	Previous []Previous
	Keys     []string
	Type     string
}

type Previous struct {
	TxID         string  `json:"txid"`
	Vout         uint32  `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	RedeemScript string  `json:"redeemScript"`
	Value        float64 `json:"value"`
}

func (SignRawTransactionCommand) Method() string {
	return "signrawtransaction"
}

func (SignRawTransactionCommand) ID() string {
	return "1"
}

func (cmd SignRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Hex, cmd.Previous, cmd.Keys, cmd.Type}
}
