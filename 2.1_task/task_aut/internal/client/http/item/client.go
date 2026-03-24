package item

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

func HttpPostItem(t *testing.T, body string) *http.Response {
	return apiRunner.GetRunner().Create().
		Post(path.CreateItemPath).
		Header("Content-Type", "application/json").
		Body(body).
		Expect(t).
		End().Response
}

func HttpGetItem(t *testing.T, id string) *http.Response {
	urlPath := fmt.Sprintf(path.GetItemByIDPath, id)

	return apiRunner.GetRunner().Create().
		Get(urlPath).
		Header("Accept", "application/json").
		Expect(t).
		End().Response
}
