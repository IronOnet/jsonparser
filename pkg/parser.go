package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// IsFileValidJson returns 0 if the file is
// a valid json file and 1 if not.
func IsValidJSON(filepath string) int {
	if _, err := isValidJSON(filepath); err != nil{
		return 0
	}
	return 1
}

func isValidJSON(filepath string) (bool, error){
	// Open file 
	file, err := os.ReadFile(filepath)
	if err != nil{
		return false, fmt.Errorf("error, could not open file: %s", err)
	}

	jsonString := string(file)
	if _, err := parseJSON(jsonString); err != nil{
		return false, fmt.Errorf("error, could not parse file: %s", err)
	}

	return true, nil
}


type JSONNode struct {
	Type     string
	Value    string
	Children []*JSONNode
}

func parseJSON(jsonStr string) (*JSONNode, error) {
	var jsonData any
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return nil, err
	}
	return buildParseTree(jsonData), nil
}

func buildParseTree(jsonData any) *JSONNode {
	node := &JSONNode{}

	switch jsonData := jsonData.(type) {
	case map[string]any:
		node.Type = "object"
		for key, value := range jsonData {
			childNode := &JSONNode{
				Type:  "pair",
				Value: key,
			}
			childNode.Children = append(childNode.Children, buildParseTree(value))
			node.Children = append(node.Children, childNode)
		}
	case []interface{}:
		node.Type = "array"
		for _, value := range jsonData {
			node.Children = append(node.Children, buildParseTree(value))
		}
	case string:
		node.Type = "string"
		node.Value = jsonData
	case float64:
		node.Type = "number"
		node.Value = fmt.Sprintf("%f", jsonData)
	case bool:
		node.Type = "boolean"
		node.Value = fmt.Sprintf("%t", jsonData)
	case nil:
		node.Type = "null"
		node.Value = "null"
	}

	return node
}