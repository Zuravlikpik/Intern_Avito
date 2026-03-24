package item

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	itemManager "api-tests-template/internal/managers/item"
	base "api-tests-template/tests"
)

type GetItemTestSuite struct {
	suite.Suite
}

func TestGetItemSuiteRun(t *testing.T) {
	suite.Run(t, &GetItemTestSuite{})
}

func (s *GetItemTestSuite) SetupSuite() {
	base.SetupSuite()
}

// успешное получение объявления по валидному ID
func (s *GetItemTestSuite) TestGetItemByValidID() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Получение объявления по валидному ID", func() {
		base.Precondition("Отправляем GET /item/:id с существующим ID")

		response := itemManager.GetItemByID(s.T(), validID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		item := response[0]

		assert.Equal(s.T(), validID, item.ID)
		assert.NotEmpty(s.T(), item.Name)
		assert.NotZero(s.T(), item.Price)
		assert.NotZero(s.T(), item.SellerID)
		assert.NotNil(s.T(), item.Statistics)
		assert.NotNil(s.T(), item.Statistics.Likes)
		assert.NotNil(s.T(), item.Statistics.ViewCount)
		assert.NotNil(s.T(), item.Statistics.Contacts)

		s.T().Logf("Ответ API: %+v", item)
	})
}

func (s *GetItemTestSuite) TestGetItemStructure() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Проверка структуры ответа объявления", func() {
		base.Precondition("Отправляем GET /item/:id и проверяем поля JSON")

		response := itemManager.GetItemByID(s.T(), validID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		item := response[0]

		assert.NotEmpty(s.T(), item.ID)
		assert.NotEmpty(s.T(), item.Name)
		assert.NotZero(s.T(), item.Price)
		assert.NotZero(s.T(), item.SellerID)
		assert.NotEmpty(s.T(), item.CreatedAt)
		assert.NotNil(s.T(), item.Statistics)
		assert.NotNil(s.T(), item.Statistics.Likes)
		assert.NotNil(s.T(), item.Statistics.ViewCount)
		assert.NotNil(s.T(), item.Statistics.Contacts)
	})
}

// идемпотентность GET
func (s *GetItemTestSuite) TestGetItemIdempotency() {
	validID := "51c64a9e-02b3-4793-9bf8-f23ed77cd6d4"

	s.Run("Идемпотентность GET запроса", func() {
		base.Precondition("Повторные GET /item/:id должны возвращать одинаковый результат")

		response1 := itemManager.GetItemByID(s.T(), validID, http.StatusOK)
		response2 := itemManager.GetItemByID(s.T(), validID, http.StatusOK)
		response3 := itemManager.GetItemByID(s.T(), validID, http.StatusOK)

		// Проверяем, что все три ответа одинаковые
		assert.Equal(s.T(), response1, response2)
		assert.Equal(s.T(), response2, response3)

		s.T().Logf("GET запросы возвращают одинаковые данные: %+v", response1)
	})
}
