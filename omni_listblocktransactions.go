package omnilayer

import "encoding/json"

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
