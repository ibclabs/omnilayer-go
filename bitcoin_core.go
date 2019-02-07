package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

type futureCreateRawTransaction chan *response

func (f futureCreateRawTransaction) Receive() (omnijson.CreateRawTransactionResult, error) {
	var result omnijson.CreateRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetBlockChainInfo chan *response

func (f futureGetBlockChainInfo) Receive() (omnijson.GetBlockChainInfoResult, error) {
	var result omnijson.GetBlockChainInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureListUnspent chan *response

func (f futureListUnspent) Receive() (omnijson.ListUnspentResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(omnijson.ListUnspentResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureImportAddress chan *response

func (f futureImportAddress) Receive() error {
	_, err := receive(f)
	return err
}

type futureSendRawTransaction chan *response

func (f futureSendRawTransaction) Receive() (omnijson.SendRawTransactionResult, error) {
	var res omnijson.SendRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	return res, err
}
