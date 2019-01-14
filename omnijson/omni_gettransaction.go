package omnijson

/*
Result:
{
  "txid" : "hash",                  (string) the hex-encoded hash of the transaction
  "sendingaddress" : "address",     (string) the Bitcoin address of the sender
  "referenceaddress" : "address",   (string) a Bitcoin address used as reference (if any)
  "ismine" : true|false,            (boolean) whether the transaction involes an address in the wallet
  "confirmations" : nnnnnnnnnn,     (number) the number of transaction confirmations
  "fee" : "n.nnnnnnnn",             (string) the transaction fee in bitcoins
  "blocktime" : nnnnnnnnnn,         (number) the timestamp of the block that contains the transaction
  "valid" : true|false,             (boolean) whether the transaction is valid
  "invalidreason" : "reason",     (string) if a transaction is invalid, the reason
  "version" : n,                    (number) the transaction version
  "type_int" : n,                   (number) the transaction type as number
  "type" : "type",                  (string) the transaction type as string
  [...]                             (mixed) other transaction type specific properties

  "propertyid": 2,
  "divisible": boolean
  "blockhash": "00000000000000ca5f93a53a5400ee909a560f4f5255ae72816ace0d6a9c3b11",
  "positioninblock": 65,
  "block": 1451927
}
*/

type OmniGettransactionResult struct {
	ID              string `json:"txid"`
	From            string `json:"sendingaddress"`
	To              string `json:"referenceaddress"`
	Fee             string `json:"fee"`
	InvalidReason   string `json:"invalidreason"`
	Type            string `json:"type"`
	BlockHash       string `json:"blockhash"`
	Confirmations   uint32 `json:"confirmations"`
	BlockTimestamp  int32  `json:"blocktime"`
	Version         int32  `json:"version"`
	TypeInt         int32  `json:"type_int"`
	PropertyID      int32  `json:"propertyid"`
	PositionInBlock int32  `json:"positioninblock"`
	Block           int32  `json:"block"`
	Mine            bool   `json:"ismine"`
	Valid           bool   `json:"valid"`
	Divisible       bool   `json:"divisible"`
}

type OmniGettransactionCommand struct {
	Hash string
}

func (OmniGettransactionCommand) Method() string {
	return "omni_gettransaction"
}

func (OmniGettransactionCommand) ID() string {
	return "1"
}

func (cmd OmniGettransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Hash}
}
