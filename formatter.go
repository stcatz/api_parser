package api_parser

import (
  "gopkg.in/yaml.v1"
  "encoding/json"
)

func ConstructJson(b []byte) (map[string]interface{}, error){
	var f map[string]interface{}
	err := json.Unmarshal(b, &f)
	return f, err
}

func ConstructYaml(b []byte) (map[string]interface{}, error){
	var f map[string]interface{}
	err := yaml.Unmarshal(b, &f)
	return f, err
}
