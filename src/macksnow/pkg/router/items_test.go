package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestItemsShowRouter(t *testing.T) {
	e := echo.New()
	itemsShowRouter(e)

	// 正しい:idでリクエストした場合
	validReq := httptest.NewRequest("GET", "/v1/items/1", nil)
	validRec := httptest.NewRecorder()
	e.ServeHTTP(validRec, validReq)
	assert.Equal(t, http.StatusOK, validRec.Code)

	// 正しくない:idでリクエストした場合
	invalidReq := httptest.NewRequest("GET", "/v1/items/hoge", nil)
	invalidRec := httptest.NewRecorder()
	e.ServeHTTP(invalidRec, invalidReq)

	body, err := ioutil.ReadAll(invalidRec.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, invalidRec.Code)
	assert.Equal(t, "{\"message\":\"hoge is not integer.\"}\n", string(body))
}
