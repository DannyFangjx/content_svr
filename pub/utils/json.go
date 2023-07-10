package utils

import (
	"content_svr/pub/errors"
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
	"strings"
)

func JsonUnmarshal(jsonStr string, dest interface{}) error {
	if jsonStr == "" {
		return nil
	}
	err := json.Unmarshal([]byte(jsonStr), dest)
	return errors.Wrap(err)
}

func JsonMarshal(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", errors.Wrap(err)
	}
	return string(bytes), nil
}

func JsonValidate(schema map[string]interface{}, source interface{}) (string, error) {
	goLoader := gojsonschema.NewGoLoader(schema)
	verify, err := gojsonschema.NewSchema(goLoader)
	if err != nil {
		return "", err
	}
	doc := gojsonschema.NewRawLoader(source)
	result, err := verify.Validate(doc)
	if err != nil {
		return "", err
	}
	var errMsgs []string
	if !result.Valid() {
		for _, err := range result.Errors() {
			errStr := err.String()
			errStr = strings.TrimSuffix(errStr, "/1")
			errStr = strings.TrimPrefix(errStr, "(root): ")
			errMsgs = append(errMsgs, errStr)
		}
	}
	return strings.Join(errMsgs, "; "), nil
}
