package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
