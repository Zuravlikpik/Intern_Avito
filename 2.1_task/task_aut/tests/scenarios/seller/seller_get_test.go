package seller

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sellerManager "api-tests-template/internal/managers/seller"
	base "api-tests-template/tests"
)

type GetSellerItemsTestSuite struct {
	suite.Suite
}

func TestGetSellerItemsSuiteRun(t *testing.T) {
	suite.Run(t, &GetSellerItemsTestSuite{})
}

func (s *GetSellerItemsTestSuite) SetupSuite() {
	base.SetupSuite()
}

// Получение всех объявлений по валидному sellerID
func (s *GetSellerItemsTestSuite) TestGetItemsByValidSellerID() {
	validSellerID := 378170

	s.Run("Получение объявлений по валидному sellerID", func() {
		base.Precondition("Отправляем GET /{sellerID}/item с существующим sellerID")

		response := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		for _, item := range response {
			assert.Equal(s.T(), validSellerID, item.SellerID, "sellerId должен совпадать")
			assert.NotEmpty(s.T(), item.ID, "ID объявления не должен быть пустым")
			assert.NotEmpty(s.T(), item.Name, "Название объявления не должно быть пустым")
			assert.NotZero(s.T(), item.Price, "Цена не должна быть нулевой")
			assert.NotEmpty(s.T(), item.CreatedAt, "Дата создания не должна быть пустой")
			assert.NotNil(s.T(), item.Statistics, "Statistics не должен быть nil")
		}

		s.T().Logf("Получено %d объявлений для sellerID %d", len(response), validSellerID)
	})
}

func (s *GetSellerItemsTestSuite) TestGetItemsStructure() {
	validSellerID := 378170

	s.Run("Проверка структуры ответа списка объявлений", func() {
		base.Precondition("Отправляем GET /{sellerID}/item и проверяем структуру JSON")

		response := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым массивом")

		for _, item := range response {
			assert.NotEmpty(s.T(), item.ID, "поле id обязательно")
			assert.NotZero(s.T(), item.SellerID, "поле sellerId обязательно")
			assert.NotEmpty(s.T(), item.Name, "поле name обязательно")
			assert.NotZero(s.T(), item.Price, "поле price обязательно")
			assert.NotEmpty(s.T(), item.CreatedAt, "поле createdAt обязательно")

			assert.NotNil(s.T(), item.Statistics, "поле statistics обязательно")
			assert.NotNil(s.T(), item.Statistics.Likes, "поле likes обязательно")
			assert.NotNil(s.T(), item.Statistics.ViewCount, "поле viewCount обязательно")
			assert.NotNil(s.T(), item.Statistics.Contacts, "поле contacts обязательно")
		}

		s.T().Logf("Структура ответа корректна для %d объявлений", len(response))
	})
}

// получение объявлений для sellerID без объявлений
func (s *GetSellerItemsTestSuite) TestGetItemsBySellerIDWithNoItems() {
	sellerIDWithNoItems := 516293

	s.Run("Получение объявлений для sellerID без объявлений", func() {
		base.Precondition("Отправляем GET /{sellerID}/item для sellerID без объявлений")

		response := sellerManager.GetItemsBySellerID(s.T(), sellerIDWithNoItems, http.StatusOK)

		assert.Empty(s.T(), response, "Должен возвращаться пустой массив")
		assert.Equal(s.T(), 0, len(response), "Длина массива должна быть 0")

		s.T().Logf("Для sellerID %d получен пустой массив", sellerIDWithNoItems)
	})
}

// идемпотентность GET запроса
func (s *GetSellerItemsTestSuite) TestGetItemsIdempotency() {
	validSellerID := 378170

	s.Run("Идемпотентность GET запроса по sellerID", func() {
		base.Precondition("Повторные GET /{sellerID}/item должны возвращать одинаковый результат")

		response1 := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)
		response2 := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)
		response3 := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)

		assert.Equal(s.T(), response1, response2, "Первый и второй ответы должны быть одинаковыми")
		assert.Equal(s.T(), response2, response3, "Второй и третий ответы должны быть одинаковыми")

		assert.NotEmpty(s.T(), response1, "Ответ не должен быть пустым")
		assert.Equal(s.T(), len(response1), len(response2), "Количество объявлений должно совпадать")

		s.T().Logf("GET запросы идемпотентны. Получено %d объявлений", len(response1))
	})
}

// Проверка, что все объявления действительно принадлежат продавцу
func (s *GetSellerItemsTestSuite) TestGetItemsAllBelongToSeller() {
	validSellerID := 378170

	s.Run("Проверка принадлежности всех объявлений указанному продавцу", func() {
		base.Precondition("Отправляем GET /{sellerID}/item и проверяем sellerId в каждом объявлении")

		response := sellerManager.GetItemsBySellerID(s.T(), validSellerID, http.StatusOK)

		assert.NotEmpty(s.T(), response, "Ответ не должен быть пустым")

		for i, item := range response {
			assert.Equal(s.T(), validSellerID, item.SellerID,
				"Объявление %d (ID: %s) принадлежит другому продавцу", i, item.ID)
		}

		s.T().Logf("Все %d объявлений принадлежат продавцу %d", len(response), validSellerID)
	})
}
