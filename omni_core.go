package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

type futureOmniCreatePayloadSimpleSend chan *response

func (f futureOmniCreatePayloadSimpleSend) Receive() (*omnijson.OmniCreatePayloadSimpleSendResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniCreatePayloadSimpleSendResult)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type futureOmniCreateRawTxChange chan *response

func (f futureOmniCreateRawTxChange) Receive() (*omnijson.OmniCreateRawTxChangeResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniCreateRawTxChangeResult)

	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

type futureOmniCreateRawTxOpReturn chan *response

func (f futureOmniCreateRawTxOpReturn) Receive() (*omnijson.OmniCreateRawTxOpReturnResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniCreateRawTxOpReturnResult)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type futureOmniCreateRawTxReference chan *response

func (f futureOmniCreateRawTxReference) Receive() (*omnijson.OmniCreateRawTxReferenceResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniCreateRawTxReferenceResult)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type futureGetInfo chan *response

func (f futureGetInfo) Receive() (*omnijson.OmniGetInfoResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniGetInfoResult)

	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

type futureOmniGetTransaction chan *response

func (f futureOmniGetTransaction) Receive() (*omnijson.OmniGettransactionResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := new(omnijson.OmniGettransactionResult)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type futureOmniListBlockTransactions chan *response

func (f futureOmniListBlockTransactions) Receive() ([]string, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
