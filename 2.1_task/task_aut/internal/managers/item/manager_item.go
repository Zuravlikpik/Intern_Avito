package item

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	itemClient "api-tests-template/internal/client/http/item"
	itemModels "api-tests-template/internal/managers/item/models"
)

func CreateItem(t *testing.T, request itemModels.CreateItemRequest, expectedStatus int) itemModels.CreateItemResponse {
	bodyBytes, err := json.Marshal(request)
	assert.NoError(t, err)

	resp := itemClient.HttpPostItem(t, string(bodyBytes))
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp != nil {
		assert.Equal(t, expectedStatus, resp.StatusCode)
	}

	var responseBody []byte
	if resp != nil && resp.Body != nil {
		responseBody, err = io.ReadAll(resp.Body)
		assert.NoError(t, err)
	}

	t.Logf("Response: %s", string(responseBody))

	var result itemModels.CreateItemResponse
	err = json.Unmarshal(responseBody, &result)
	assert.NoError(t, err)

	return result
}

func GetItemByID(t *testing.T, id string, expectedStatus int) []itemModels.ItemResponse {
	// Изменяем вызов: убираем обработку ошибки, так как теперь её нет
	resp := itemClient.HttpGetItem(t, id)
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp != nil {
		assert.Equal(t, expectedStatus, resp.StatusCode)
	}

	assert.Equal(t, expectedStatus, resp.StatusCode)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	t.Logf("Response: %s", string(bodyBytes))

	// т.к. API всегда возвращает массив
	var result []itemModels.ItemResponse
	err = json.Unmarshal(bodyBytes, &result)
	assert.NoError(t, err)

	return result
}
