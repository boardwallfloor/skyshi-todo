package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var ErrIdNotFound = errors.New("id not found")

func ErrResourceNotFound(id int, resource string) string {
	var resName string
	if len(resource) < 1 {
		resName = resource
	} else {
		resName = strings.ToUpper(resource[0:1]) + strings.ToLower(resource[1:])
	}
	return fmt.Sprintf("%s with ID %d Not Found", resName, id)
}

func parseTime(t []uint8) time.Time {
	tString := string(t)
	date, err := time.Parse("2006-01-02 15:04:05", tString)
	if err != nil {
		log.Fatal(err)
	}
	return date.UTC()
}

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
			"data":    map[string]interface{}{},
		}
	}

	respJson, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
