package test

import (
	"delivery-app/filemanager"
	"delivery-app/offers"
	"encoding/json"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestReadOfferCodesJson(t *testing.T) {
	// Create a temporary JSON file with sample data
	file, err := os.CreateTemp("", "offers_test.json")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	sampleData := map[string]offers.Offer{
		"offer1": {
			DistanceCriteria: offers.Criteria{Min: 0, Max: 10},
			WeightCriteria:   offers.Criteria{Min: 0, Max: 5},
			Discount:         10,
		},
		"offer2": {
			DistanceCriteria: offers.Criteria{Min: 10, Max: 20},
			WeightCriteria:   offers.Criteria{Min: 5, Max: 10},
			Discount:         20,
		},
	}

	byteValue, err := json.Marshal(sampleData)
	assert.NoError(t, err)

	_, err = file.Write(byteValue)
	assert.NoError(t, err)

	file.Close()

	result, err := filemanager.ReadOfferCodesJson(file.Name())
	assert.NoError(t, err)

	// Check the result is same as the sample data
	assert.Equal(t, sampleData, result)

	// Check error handling with a non-existent file
	_, err = filemanager.ReadOfferCodesJson("non_existent_file.json")
	assert.Error(t, err)
}

func TestWriteOfferCodesJson(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "offers_*.json")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	sampleData := map[string]offers.Offer{
		"offer1": {
			DistanceCriteria: offers.Criteria{Min: 0, Max: 10},
			WeightCriteria:   offers.Criteria{Min: 0, Max: 5},
			Discount:         10,
		},
		"offer2": {
			DistanceCriteria: offers.Criteria{Min: 10, Max: 20},
			WeightCriteria:   offers.Criteria{Min: 5, Max: 10},
			Discount:         20,
		},
	}

	err = filemanager.WriteOfferCodesJson(sampleData, file.Name())
	assert.NoError(t, err)

	// Read the file back
	content, err := os.ReadFile(file.Name())
	assert.NoError(t, err)

	// Unmarshal the content
	var result map[string]offers.Offer
	err = json.Unmarshal(content, &result)
	assert.NoError(t, err)

	// Check if the written data matches the sample data
	assert.Equal(t, sampleData, result)
}
