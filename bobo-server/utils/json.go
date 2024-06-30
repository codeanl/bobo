package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Json = new(_json)

type _json struct{}

// data -> jsonStr
func (*_json) Marshal(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		Logger.Error("utils/json.go -> Marshal: ", zap.Error(err))
		panic(err)
	}
	return string(data)
}

// jsonStr -> data
func (*_json) Unmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		Logger.Error("utils/json.go -> Marshal: ", zap.Error(err))
		panic(err)
	}
}
