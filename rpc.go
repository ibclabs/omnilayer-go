package omnilayer

import "github.com/ibclabs/omnilayer-go/omnijson"

func (c *Client) GetBlockChainInfo() (*omnijson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(omnijson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) OmniListBlockTransactions(block int64) ([]string, error) {
	return futureOmniListBlockTransactions(c.do(omnijson.OmniListBlockTransactionsCommand{
		Block: block,
	})).Receive()
}

func (c *Client) GetInfo() (*omnijson.OmniGetInfoResult, error) {
	return futureGetInfo(c.do(omnijson.OmniGetInfoCommand{})).Receive()
}

func (c *Client) OmniGetTransaction(hash string) (*omnijson.OmniGettransactionResult, error) {
	return futureOmniGetTransaction(c.do(omnijson.OmniGetTransactionCommand{
		Hash: hash,
	})).Receive()
}

func (c *Client) ListUnspent(addresses []string, min int) ([]omnijson.ListUnspentResult, error) {
	return futureListUnspent(c.do(omnijson.ListUnspentCommand{
		Addresses: addresses,
		Min:       min,
	})).Receive()
}

func (c *Client) OmniCreatePayloadSimpleSend(
	property int32, amount string,
) (*omnijson.OmniCreatePayloadSimpleSendResult, error) {
	return futureOmniCreatePayloadSimpleSend(c.do(omnijson.OmniCreatePayloadSimpleSendCommand{
		Property: property,
		Amount:   amount,
	})).Receive()
}

func (c *Client) CreateRawTransaction(params []omnijson.CreateRawTransactionParameter) (*omnijson.CreateRawTransactionResult, error) {
	return futureCreateRawTransaction(c.do(omnijson.CreateRawTransactionCommand{
		Parameters: params,
	})).Receive()
}

func (c *Client) OmniCreateRawTxOpReturn(raw, payload string) (*omnijson.OmniCreateRawTxOpReturnResult, error) {
	return futureOmniCreateRawTxOpReturn(c.do(omnijson.OmniCreateRawTxOpReturnCommand{
		Raw:     raw,
		Payload: payload,
	})).Receive()
}

func (c *Client) OmniCreateRawTxReference(raw, destination, amount string) (*omnijson.OmniCreateRawTxReferenceResult, error) {
	return futureOmniCreateRawTxReference(c.do(omnijson.OmniCreateRawTxReferenceCommand{
		Raw:         raw,
		Destination: destination,
		Amount:      amount,
	})).Receive()
}

func (c *Client) OmniCreateRawTxChange(
	raw, destination string, fee float64,
	prevs []omnijson.OmniCreateRawTxChangeParameter) (*omnijson.OmniCreateRawTxChangeResult, error) {
	return futureOmniCreateRawTxChange(c.do(omnijson.OmniCreateRawTxChangeCommand{
		Raw:         raw,
		Destination: destination,
		Fee:         fee,
	})).Receive()
}

func (c *Client) ImportAddress(address string, rescan bool) error {
	return futureImportAddress(c.do(omnijson.ImportAddressCommand{Adress: address, Rescan: rescan})).Receive()
}
