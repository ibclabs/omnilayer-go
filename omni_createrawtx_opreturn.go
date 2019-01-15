package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
