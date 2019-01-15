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
