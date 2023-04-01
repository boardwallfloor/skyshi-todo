package main

import (
	"encoding/json"
)

func wrapResp(data interface{}, status string, message string) ([]byte, error) {
	var response map[string]interface{}
	if data != nil {
		response = map[string]interface{}{
			"status":  status,
			"message": message,
			"data":    data,
		}
	} else {
		response = map[string]interface{}{
			"status":  status,
			"message": message,
		}
	}

	respJson, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
