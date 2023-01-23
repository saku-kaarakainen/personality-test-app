package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// APITestCase represents the data needed to describe an API test case.
type APITestCase struct {
	Name         string
	Method, URL  string
	Body         string
	Header       http.Header
	WantStatus   int
	WantResponse string
}

// Endpoint tests an HTTP endpoint using the given APITestCase spec.
func Endpoint(t *testing.T, router *gin.Engine, tc APITestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
		if tc.Header != nil {
			req.Header = tc.Header
		}
		res := httptest.NewRecorder()
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}
		router.ServeHTTP(res, req)
		assert.Equal(t, tc.WantStatus, res.Code, "status mismatch. Response: "+res.Body.String())
		if tc.WantResponse == "" {
			return
		}

		pattern := strings.Trim(tc.WantResponse, "*")
		str := res.Body.String()
		if pattern != tc.WantResponse {
			assert.Contains(t, str, pattern, "response mismatch (assert.Contains). str: "+str)
		} else {
			assert.JSONEq(t, tc.WantResponse, str, "response mismatch (assert.JSONEq). str: "+str)
		}

	})
}
