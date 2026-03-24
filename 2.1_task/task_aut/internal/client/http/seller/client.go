package seller

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

func HttpGetItemsBySeller(t *testing.T, sellerID int) *http.Response {
	urlPath := fmt.Sprintf(path.GetItemsBySeller, sellerID)

	return apiRunner.GetRunner().Create().
		Get(urlPath).
		Header("Accept", "application/json").
		Expect(t).
		End().Response
}
