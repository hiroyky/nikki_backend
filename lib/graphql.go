package lib

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func EncodeGraphQLID(modelName string, id int) string {
	if strings.Contains(modelName, "-") == true {
		panic(fmt.Sprintf("Invalid modelName: %s", modelName))
	}

	format := fmt.Sprintf("%s-%d", modelName, id)
	return base64.StdEncoding.EncodeToString([]byte(format))
}

func DecodeGraphQLID(graphID string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(graphID)
	if err != nil {
		return 0, err
	}
	decodedGraphID := string(decoded)
	parts := strings.Split(decodedGraphID, "-")
	if len(parts) != 2 {
		return 0, fmt.Errorf("GraphQL ID: %s is invalid ID", decodedGraphID)
	}

	return strconv.Atoi(parts[1])
}
