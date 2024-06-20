package filemanager

import (
	"delivery-app/offers"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadOfferCodesJson(filePath string) (map[string]offers.Offer, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open the codes file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	data := make(map[string]offers.Offer)

	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}

func WriteOfferCodesJson(data map[string]offers.Offer, filePath string) error {
	updatedJson, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the codes: %w", err)
	}
	err = os.WriteFile(filePath, updatedJson, 0644)
	if err != nil {
		return fmt.Errorf("write to offer code files failed: %w", err)
	}
	return nil
}
