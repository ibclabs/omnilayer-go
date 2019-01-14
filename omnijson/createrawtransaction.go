package omnijson

/*
Result:
"transaction"            (string) hex string of the transaction
*/

type CreateRawTransactionResult string

type CreateRawTransactionCommand struct {
	Parameters []CreateRawTransactionParameter
}

type CreateRawTransactionParameter struct {
	Tx   string
	Vout uint32
}

type createrawtransactionOutput struct {
	Address string
	Data    string
}

func (CreateRawTransactionCommand) Method() string {
	return "createrawtransaction"
}

func (CreateRawTransactionCommand) ID() string {
	return "1"
}

func (cmd CreateRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Parameters, createrawtransactionOutput{Address: "", Data: ""}}
}
