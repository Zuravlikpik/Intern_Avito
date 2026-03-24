package seller

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	sellerClient "api-tests-template/internal/client/http/seller"
	itemModels "api-tests-template/internal/managers/item/models"
)

func GetItemsBySellerID(t *testing.T, sellerID int, expectedStatus int) []itemModels.ItemResponse {
	resp := sellerClient.HttpGetItemsBySeller(t, sellerID)
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp != nil {
		assert.Equal(t, expectedStatus, resp.StatusCode)
	}

	var responseBody []byte
	var err error
	if resp != nil && resp.Body != nil {
		responseBody, err = io.ReadAll(resp.Body)
		assert.NoError(t, err)
	}

	t.Logf("Response: %s", string(responseBody))

	var result []itemModels.ItemResponse
	if len(responseBody) > 0 {
		err = json.Unmarshal(responseBody, &result)
		assert.NoError(t, err)
	}

	return result
}

// функция для обработки ошибок (когда возвращается не массив, а объект с ошибкой)
func GetItemsBySellerIDError(t *testing.T, sellerID int, expectedStatus int) map[string]interface{} {
	resp := sellerClient.HttpGetItemsBySeller(t, sellerID)
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp != nil {
		assert.Equal(t, expectedStatus, resp.StatusCode)
	}

	var responseBody []byte
	var err error
	if resp != nil && resp.Body != nil {
		responseBody, err = io.ReadAll(resp.Body)
		assert.NoError(t, err)
	}

	t.Logf("Response: %s", string(responseBody))

	var result map[string]interface{}
	if len(responseBody) > 0 {
		err = json.Unmarshal(responseBody, &result)
		assert.NoError(t, err)
	}

	return result
}
