package statistic

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	statisticClient "api-tests-template/internal/client/http/statistic"
	statisticModels "api-tests-template/internal/managers/statistic/models"
)

// получаем статистику по ID объявления
func GetStatisticByID(t *testing.T, id string, expectedStatus int) []statisticModels.StatisticResponse {
	resp := statisticClient.HttpGetStatisticByID(t, id)
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

	var result []statisticModels.StatisticResponse
	if len(responseBody) > 0 {
		err = json.Unmarshal(responseBody, &result)
		assert.NoError(t, err)
	}

	return result
}

// получаем статистику с возможной ошибкой
func GetStatisticByIDWithError(t *testing.T, id string, expectedStatus int) map[string]interface{} {
	resp := statisticClient.HttpGetStatisticByID(t, id)
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
