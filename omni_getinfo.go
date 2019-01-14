package omni_client

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
