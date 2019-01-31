package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

type futureCreateRawTransaction chan *response

func (f futureCreateRawTransaction) Receive() (*omnijson.CreateRawTransactionResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.CreateRawTransactionResult)

	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

type futureGetBlockChainInfo chan *response

func (f futureGetBlockChainInfo) Receive() (*omnijson.GetBlockChainInfoResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.GetBlockChainInfoResult)

	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

type futureListUnspent chan *response

func (f futureListUnspent) Receive() ([]omnijson.ListUnspentResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make([]omnijson.ListUnspentResult, 0)

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type futureImportAddress chan *response

func (f futureImportAddress) Receive() error {
	_, err := receive(f)
	return err
}
