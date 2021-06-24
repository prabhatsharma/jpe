package v1alpha1

import (
	"encoding/json"
	"fmt"
	"time"
)

var logger MyLogger

// var jsonLogger JSONLogger

func (logm *JSONLogger) LogJSON(description string, val map[string]interface{}) {
	val["timestamp"] = time.Now()
	val["description"] = description

	logMessage, _ := json.Marshal(val)

	fmt.Println(string(logMessage))
}

func (logm *MyLogger) LogStuff(description string, val ...interface{}) {

	logm.Timestamp = time.Now()
	logm.Message = val
	logm.Description = description

	logMessage, _ := json.Marshal(logm)

	fmt.Println(string(logMessage))
}

type MyLogger struct {
	Timestamp   time.Time     `json:"timestamp"`
	Description string        `json:"description"`
	Message     []interface{} `json:"message"`
}

type JSONLogger struct {
	Timestamp   time.Time     `json:"timestamp"`
	Description string        `json:"description"`
	Message     []interface{} `json:"message"`
}
