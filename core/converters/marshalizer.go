package converters

import (
	"encoding/json"
)

func Marshal(obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func Unmarshal(bytes []byte, obj interface{}) error {
	err := json.Unmarshal(bytes, obj)
	if err != nil {
		return err
	}
	return nil
}
