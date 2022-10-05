package helpers

import (
	"bytes"
	modelRef "go-rest-starter/src/api/models/req"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseExtraCientInfoWithEmptyValues(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	w.Header().Add("User-Agent", "")
	w.Header().Add("Referer", "")
	w.Header().Add("Host", "")
	testCtx, r := gin.CreateTestContext(w)

	testCtx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))

	r.ServeHTTP(w, testCtx.Request)

	var req modelRef.AuthInitReq

	ParseExtraClientInfo(testCtx, &req)

	assert.Equal("", req.IP)
	assert.Equal("", req.UserAgent)
	assert.Equal("", req.Referer)
	assert.Equal("", req.Host)
}

func TestParseExtraCientInfoWithTestValues(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	headers := make(http.Header)
	headers.Set("User-Agent", "Mozilla v103")
	headers.Set("Referer", "http://127.0.0.1:7500")
	headers.Set("Host", "http://127.0.0.1/")

	ctx.Request = &http.Request{
		Header: headers,
		Host:   "127.0.0.1",
	}

	var req modelRef.AuthInitReq

	ParseExtraClientInfo(ctx, &req)

	assert.Equal("", req.IP)
	assert.Equal("Mozilla v103", req.UserAgent)
	assert.Equal("http://127.0.0.1:7500", req.Referer)
	assert.Equal("http://127.0.0.1/", req.Host)
}
