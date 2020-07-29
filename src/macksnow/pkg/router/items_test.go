package router

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestItemsIndexRouter(t *testing.T) {
	e := echo.New()
	itemsIndexRouter(e)

	// リクエストした場合
	req := httptest.NewRequest("GET", "/v1/items", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

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

func TestItemsCreateRouter(t *testing.T) {
	e := echo.New()
	itemsCreateRouter(e)

	// 正しいパラメータでリクエストをした場合
	json := `{"name":"hoge","summary":"fuga","uri":"piyo"}`
	params := bytes.NewBuffer([]byte(json))

	validReq := httptest.NewRequest("POST", "/v1/items", params)
	validReq.Header.Set("Content-Type", "application/json")
	validRec := httptest.NewRecorder()

	e.ServeHTTP(validRec, validReq)
	assert.Equal(t, http.StatusOK, validRec.Code)

	// 構文エラーのパラメータでリクエストをした場合
	json = `{"name""hoge","summary":"fuga","uri":"piyo"}`
	params = bytes.NewBuffer([]byte(json))

	invalidReq := httptest.NewRequest("POST", "/v1/items", params)
	invalidReq.Header.Set("Content-Type", "application/json")
	invalidRec := httptest.NewRecorder()
	e.ServeHTTP(invalidRec, invalidReq)

	body, err := ioutil.ReadAll(invalidRec.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, invalidRec.Code)
	assert.Equal(t, "{\"message\":\"Invalid parameters.\"}\n", string(body))

	// 必須パラメータが不足したリクエストをした場合
	json = `{"name":"","summary":"fuga","uri":"piyo"}`
	params = bytes.NewBuffer([]byte(json))

	incompleteReq := httptest.NewRequest("POST", "/v1/items", params)
	incompleteReq.Header.Set("Content-Type", "application/json")
	incompleteRec := httptest.NewRecorder()
	e.ServeHTTP(incompleteRec, incompleteReq)

	body, err = ioutil.ReadAll(incompleteRec.Result().Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, incompleteRec.Code)
	assert.Equal(t, `{"message":"Required parameters is empty. Key: 'item.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`+"\n", string(body))
}
