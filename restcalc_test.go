package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testResponse struct {
	Expr string `json:"expr"`
	Res  int    `json:"res"`
	Err  error  `json:"err"`
}

func TestEvaluateExpression(t *testing.T) {
	testCases := []struct {
		Expr string
		Res  int
		Err  error
	}{
		{
			Expr: "1  + 2 -  7",
			Res:  -4,
			Err:  nil,
		},
		{
			Expr: "1  +8-    2 ",
			Res:  7,
			Err:  nil,
		},
		{
			Expr: "1    ",
			Res:  1,
			Err:  nil,
		},
	}

	handler := http.HandlerFunc(evaluateExpression)

	for _, tc := range testCases {

		t.Run(tc.Expr, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:9000/evaluate/?expr=%s", tc.Expr), nil)

			handler.ServeHTTP(rec, req)

			var tr testResponse
			err := json.Unmarshal(rec.Body.Bytes(), &tr)

			if err != nil {
				log.Fatal(err)
			}

			assert.Equal(t, tc.Res, tr.Res, fmt.Sprintf("incorrect result. Expected %d. Got %d", tc.Res, tr.Res))
			assert.Equal(t, tc.Err, tr.Err, fmt.Sprintf("incorrect error. Expected %d. Got %d", tc.Err, tr.Err))
		})
	}
}
