package statistic

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	statisticManager "api-tests-template/internal/managers/statistic"
	base "api-tests-template/tests"
)

type GetStatisticTestSuite struct {
	suite.Suite
}

func TestGetStatisticSuiteRun(t *testing.T) {
	suite.Run(t, &GetStatisticTestSuite{})
}

func (s *GetStatisticTestSuite) SetupSuite() {
	base.SetupSuite()
}

// получение статистики по валидному item ID
func (s *GetStatisticTestSuite) TestGetStatisticByValidID() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Получение статистики по валидному ID", func() {
		base.Precondition("Отправляем GET /statistic/:id с существующим ID")

		response := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		statistic := response[0]

		assert.NotZero(s.T(), statistic.Likes, "Поле likes не должно быть нулевым")
		assert.NotZero(s.T(), statistic.ViewCount, "Поле viewCount не должно быть нулевым")
		assert.NotZero(s.T(), statistic.Contacts, "Поле contacts не должно быть нулевым")

		assert.GreaterOrEqual(s.T(), statistic.Likes, 0, "Likes должно быть >= 0")
		assert.GreaterOrEqual(s.T(), statistic.ViewCount, 0, "ViewCount должно быть >= 0")
		assert.GreaterOrEqual(s.T(), statistic.Contacts, 0, "Contacts должно быть >= 0")

		s.T().Logf("Статистика для ID %s: likes=%d, viewCount=%d, contacts=%d",
			validID, statistic.Likes, statistic.ViewCount, statistic.Contacts)
	})
}

func (s *GetStatisticTestSuite) TestGetStatisticStructure() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Проверка структуры ответа статистики", func() {
		base.Precondition("Отправляем GET /statistic/:id и проверяем структуру JSON")

		response := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		statistic := response[0]

		assert.NotNil(s.T(), statistic.Likes, "поле likes обязательно")
		assert.NotNil(s.T(), statistic.ViewCount, "поле viewCount обязательно")
		assert.NotNil(s.T(), statistic.Contacts, "поле contacts обязательно")

		assert.IsType(s.T(), 0, statistic.Likes, "likes должно быть integer")
		assert.IsType(s.T(), 0, statistic.ViewCount, "viewCount должно быть integer")
		assert.IsType(s.T(), 0, statistic.Contacts, "contacts должно быть integer")

		s.T().Logf("Структура статистики корректна для ID %s", validID)
	})
}

func (s *GetStatisticTestSuite) TestGetStatisticByNonExistentID() {
	nonExistentID := "51c64a9e-02b3-4793-1bf8-f23ed77cd6d4"

	s.Run("Получение статистики по несуществующему ID", func() {
		base.Precondition("Отправляем GET /statistic/:id с несуществующим ID")

		response := statisticManager.GetStatisticByIDWithError(s.T(), nonExistentID, http.StatusNotFound)

		assert.NotNil(s.T(), response, "Ответ не должен быть nil")

		assert.Equal(s.T(), "404", response["status"], "Статус должен быть 404")

		result, ok := response["result"].(map[string]interface{})
		assert.True(s.T(), ok, "Поле result должно быть объектом")
		message, ok := result["message"].(string)
		assert.True(s.T(), ok, "Поле message должно быть строкой")
		assert.Equal(s.T(), "statistic "+nonExistentID+" not found", message)

		s.T().Logf("Получена ошибка: %s", message)
	})
}

func (s *GetStatisticTestSuite) TestGetStatisticWithInvalidID() {
	invalidID := "1111"

	s.Run("Передача невалидного ID (не UUID)", func() {
		base.Precondition("Отправляем GET /statistic/:id с невалидным UUID")

		response := statisticManager.GetStatisticByIDWithError(s.T(), invalidID, http.StatusBadRequest)

		assert.NotNil(s.T(), response, "Ответ не должен быть nil")

		assert.Equal(s.T(), "400", response["status"], "Статус должен быть 400")

		result, ok := response["result"].(map[string]interface{})
		assert.True(s.T(), ok, "Поле result должно быть объектом")
		message, ok := result["message"].(string)
		assert.True(s.T(), ok, "Поле message должно быть строкой")
		assert.Equal(s.T(), "передан некорректный идентификатор объявления", message)

		s.T().Logf("Получена ошибка: %s", message)
	})
}

func (s *GetStatisticTestSuite) TestGetStatisticWithEmptyID() {
	emptyID := ""

	s.Run("Передача пустого ID", func() {
		base.Precondition("Отправляем GET /statistic/ с пустым ID")

		response := statisticManager.GetStatisticByIDWithError(s.T(), emptyID, http.StatusNotFound)

		assert.NotNil(s.T(), response, "Ответ не должен быть nil")

		message, ok := response["message"].(string)
		assert.True(s.T(), ok, "Поле message должно быть строкой")
		assert.Equal(s.T(), "route /api/1/statistic/ not found", message)

		code, ok := response["code"].(float64)
		assert.True(s.T(), ok, "Поле code должно быть числом")
		assert.Equal(s.T(), float64(400), code)

		s.T().Logf("Получена ошибка: %s (code=%d)", message, int(code))
	})
}

// идемпотентность запроса статистики
func (s *GetStatisticTestSuite) TestGetStatisticIdempotency() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Идемпотентность GET запроса статистики", func() {
		base.Precondition("Повторные GET /statistic/:id должны возвращать одинаковый результат")

		response1 := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)
		response2 := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)
		response3 := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)

		assert.Equal(s.T(), response1, response2, "Первый и второй ответы должны быть одинаковыми")
		assert.Equal(s.T(), response2, response3, "Второй и третий ответы должны быть одинаковыми")

		assert.NotEmpty(s.T(), response1, "Ответ не должен быть пустым")

		if len(response1) > 0 {
			stat1 := response1[0]
			stat2 := response2[0]

			assert.Equal(s.T(), stat1.Likes, stat2.Likes, "Likes не должны измениться")
			assert.Equal(s.T(), stat1.ViewCount, stat2.ViewCount, "ViewCount не должны измениться")
			assert.Equal(s.T(), stat1.Contacts, stat2.Contacts, "Contacts не должны измениться")
		}

		s.T().Logf("GET запросы статистики идемпотентны")
	})
}

// проверка корректности значений статистики
func (s *GetStatisticTestSuite) TestGetStatisticValuesConsistency() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Проверка консистентности значений статистики", func() {
		base.Precondition("Проверяем, что статистика возвращает ожидаемые значения")

		response := statisticManager.GetStatisticByID(s.T(), validID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым")

		statistic := response[0]

		// проверяем конкретные значения
		// знаем, что: likes=1, viewCount=1, contacts=1
		assert.Equal(s.T(), 1, statistic.Likes, "Likes должно быть 1")
		assert.Equal(s.T(), 1, statistic.ViewCount, "ViewCount должно быть 1")
		assert.Equal(s.T(), 1, statistic.Contacts, "Contacts должно быть 1")

		s.T().Logf("Статистика соответствует ожидаемым значениям: %+v", statistic)
	})
}
