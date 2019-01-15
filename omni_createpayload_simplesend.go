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
