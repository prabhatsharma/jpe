package v1alpha1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var logger MyLogger

var json2Logger JSON2Logger

func (logm *JSON2Logger) LogJSON(description string, val interface{}) {

	finalMessage := make(map[string]interface{})

	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			// fmt.Println(key.Interface(), strct.Interface())
			finalMessage[key.Interface().(string)] = strct.Interface()
		}
	}
	finalMessage["timestamp"] = time.Now()
	finalMessage["description"] = description

	logMessage, _ := json.Marshal(finalMessage)

	fmt.Println(string(logMessage))
}

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

type JSON2Logger struct {
	Timestamp   time.Time              `json:"timestamp"`
	Description string                 `json:"description"`
	Message     map[string]interface{} `json:"message"`
}

func PrintPolicies() {
	logger.LogStuff("Current PolicyList: ", AdmissionPolicies)
}
