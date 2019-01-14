package omni_client

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
	return futureOmniGetTransaction(c.do(omnijson.OmniGettransactionCommand{
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

func (c *Client) CreateRawTransaction(parameters map[string]uint32) (*omnijson.CreateRawTransactionResult, error) {
	params := make([]omnijson.CreateRawTransactionParameter, 0, len(parameters))

	for tx, vout := range parameters {
		params = append(params, omnijson.CreateRawTransactionParameter{
			Tx:   tx,
			Vout: vout,
		})
	}

	return futureCreateRawTransaction(c.do(omnijson.CreateRawTransactionCommand{
		Parameters: params,
	})).Receive()
}
