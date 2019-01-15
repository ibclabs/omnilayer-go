package omnilayer

import (
	"encoding/json"

	"github.com/ibclabs/omnilayer-go/omnijson"
)

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
