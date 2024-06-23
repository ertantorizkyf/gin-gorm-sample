package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleMiddlewareImplSuccess(t *testing.T) {
	InitTestEnv()
	router := InitTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/sample-middleware-impl", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphbmljZUBtYWlsLmNvbSIsImV4cCI6MTcxOTE1NTE3OSwiZnVsbF9uYW1lIjoiamFuaWNlIn0.O8z2VnBM_ODs0fwFuuGO3hbPkH0pZsJ-UZz7KoQ8mgw")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Equal(t, "HELLO WORLD!", respBody["data"])
}
