package statistic

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

func HttpGetStatisticByID(t *testing.T, id string) *http.Response {
	urlPath := fmt.Sprintf(path.GetStatisticByID, id)

	return apiRunner.GetRunner().Create().
		Get(urlPath).
		Header("Accept", "application/json").
		Expect(t).
		End().Response
}
