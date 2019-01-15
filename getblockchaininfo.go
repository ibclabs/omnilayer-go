package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
