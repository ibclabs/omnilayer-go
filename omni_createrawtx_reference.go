package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
