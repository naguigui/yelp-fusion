package utility

import "encoding/json"

func StructToMap(obj interface{}) (m map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &m)
	return
}
