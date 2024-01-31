package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

func RunJsonTest() (string, error) {
	jsonData := `[
        {
            "org": "puzzle",
            "project": "helm",
            "created_at": "2019-10-12T07:20:50.52Z",
            "gitURL": "https://github.com/puzzle/dagger-module-helm"
        },
        {
            "org": "dzzle",
            "project": "helm",
            "created_at": "2019-10-12T07:20:50.52Z",
            "gitURL": "https://github.com/puzzle/dagger-module-helm"
        },
        {
            "org": "fluent-ci-templates",
            "project": "wasmer-pipeline",
            "created_at": "2019-11-05T15:45:30.12Z",
            "gitURL": "https://github.com/fluent-ci-templates/wasmer-pipeline"
        }
    ]`

	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	var resultBuilder strings.Builder
	resultBuilder.WriteString("[\n")

	for i, entry := range data {
		// Marshal each entry to JSON without indentation
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			return "", fmt.Errorf("error marshalling entry to JSON: %v", err)
		}

		entryStr := strings.ReplaceAll(string(entryJSON), "\n", "")
		resultBuilder.WriteString("   ")
		resultBuilder.WriteString(entryStr)

		if i < len(data)-1 {
			resultBuilder.WriteString(",\n")
		} else {
			resultBuilder.WriteString("\n")
		}
	}

	resultBuilder.WriteString("]")

	return resultBuilder.String(), nil
}
