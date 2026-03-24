package item

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	itemManager "api-tests-template/internal/managers/item"
	itemModels "api-tests-template/internal/managers/item/models"

	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
}

// смоук
func (s *TestSuite) TestCreateItemPositive() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок"),
		Price:    itemModels.IntPtr(12),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с валидными данными", func() {
		base.Precondition("Отправляем POST /item с валидными данными")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление: %s", response.Status)
	})
}

func (s *TestSuite) TestCreateItemWithoutName() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr(""),
		Price:    itemModels.IntPtr(12),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления без обязательного поля name", func() {
		base.Precondition("Отправляем POST /item без обязательного поля name")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле name обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

func (s *TestSuite) TestCreateItemWithLongName() {
	longName := ""
	for i := 0; i < 500; i++ {
		longName += "a"
	}

	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr(longName),
		Price:    itemModels.IntPtr(12),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с очень длинным именем", func() {
		base.Precondition("Отправляем POST /item с очень длинным значением поля name")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Contains(s.T(), response.Status, "Сохранили объявление",
			"Ожидалось успешное создание объявления")

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление с длинным именем: %s", response.Status)
	})
}

func (s *TestSuite) TestCreateItemWithEmojiInName() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок 🐱🔥✨你好 👋 مرحبا"),
		Price:    itemModels.IntPtr(12),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с эмодзи в имени", func() {
		base.Precondition("Отправляем POST /item с эмодзи в поле name")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Contains(s.T(), response.Status, "Сохранили объявление",
			"Ожидалось успешное создание объявления с эмодзи")

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление с эмодзи: %s", response.Status)
	})
}

func (s *TestSuite) TestCreateItemWithNullPrice() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок"),
		Price:    nil,
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с null в поле price", func() {
		base.Precondition("Отправляем POST /item с price = null")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле price обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

func (s *TestSuite) TestCreateItemWithZeroPrice() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок бесплатный"),
		Price:    itemModels.IntPtr(0),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с price = 0", func() {
		base.Precondition("Отправляем POST /item с price = 0")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле price обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

// падает (фактич. рез. несоответствует ожидаемому)
func (s *TestSuite) TestCreateItemWithNegativePrice() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок с отрицательной цееной"),
		Price:    itemModels.IntPtr(0),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с price < 0", func() {
		base.Precondition("Отправляем POST /item с price < 0")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле price не может быть отрицательным", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

func (s *TestSuite) TestCreateItemWithHugePrice() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок"),
		Price:    itemModels.IntPtr(999999999999999999),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с очень большой ценой", func() {
		base.Precondition("Отправляем POST /item с price = 999999999999999999")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Contains(s.T(), response.Status, "Сохранили объявление",
			"Ожидалось успешное создание объявления")

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление: %s", response.Status)
	})
}

func (s *TestSuite) TestCreateItemWithNullLikes() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     nil,
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с likes = null", func() {
		base.Precondition("Отправляем POST /item с likes = null")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле likes обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

func (s *TestSuite) TestCreateItemWithNullViewCount() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: nil,
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с ViewCount = null", func() {
		base.Precondition("Отправляем POST /item с ViewCount = null")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле viewCount обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

func (s *TestSuite) TestCreateItemWithNullContacts() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  nil,
		},
	}

	s.Run("Создание объявления с Contacts = null", func() {
		base.Precondition("Отправляем POST /item с Contacts = null")

		response := itemManager.CreateItem(s.T(), request, http.StatusBadRequest)

		assert.Equal(s.T(), "400", response.Status)

		assert.Equal(s.T(), "поле contacts обязательно", response.Result.Message)

		s.T().Logf("Ответ API: %s", response.Result.Message)
	})
}

// падает
func (s *TestSuite) TestCreateItemWithZerroLikes() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(0),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с валидными данными", func() {
		base.Precondition("Отправляем POST /item с валидными данными")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление: %s", response.Status)
	})
}

// падает
func (s *TestSuite) TestCreateItemWithZerroViewCount() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(0),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Создание объявления с валидными данными", func() {
		base.Precondition("Отправляем POST /item с валидными данными")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление: %s", response.Status)
	})
}

// падает
func (s *TestSuite) TestCreateItemWithZeroContacts() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский"),
		Price:    itemModels.IntPtr(10),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(0),
		},
	}

	s.Run("Создание объявления с валидными данными", func() {
		base.Precondition("Отправляем POST /item с валидными данными")

		response := itemManager.CreateItem(s.T(), request, http.StatusOK)

		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response.Status)

		s.T().Logf("Создано объявление: %s", response.Status)
	})
}

func (s *TestSuite) TestCreateItemIdempotency() {
	request := itemModels.CreateItemRequest{
		SellerID: itemModels.IntPtr(151962),
		Name:     itemModels.StringPtr("Персидский котёнок"),
		Price:    itemModels.IntPtr(12),
		Statistics: itemModels.Statistics{
			Likes:     itemModels.IntPtr(1),
			ViewCount: itemModels.IntPtr(1),
			Contacts:  itemModels.IntPtr(1),
		},
	}

	s.Run("Идемпотентность создания объявления", func() {
		base.Precondition("Отправляем POST /item дважды с одинаковыми данными")

		// Первый запрос
		response1 := itemManager.CreateItem(s.T(), request, http.StatusOK)
		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response1.Status)
		s.T().Logf("Первое объявление создано: %s", response1.Status)

		// Второй запрос
		response2 := itemManager.CreateItem(s.T(), request, http.StatusOK)
		assert.Regexp(s.T(), `Сохранили объявление - [a-f0-9\-]+`, response2.Status)
		s.T().Logf("Второе объявление создано: %s", response2.Status)

		// Проверяем, что UUID разные
		uuid1 := strings.Split(response1.Status, " - ")[1]
		uuid2 := strings.Split(response2.Status, " - ")[1]
		assert.NotEqual(s.T(), uuid1, uuid2, "UUID двух созданных объявлений должны быть разными")
	})
}
