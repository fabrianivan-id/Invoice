package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"esb-test/src/app"
	"esb-test/src/middleware/response"

	"github.com/go-chi/chi/v5"
	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Chdir("../../../")

	app.Init(context.Background())

	exitVal := m.Run()

	os.Exit(exitVal)
}

func CheckBodyResponse(t *testing.T, actualResponse []byte, expected interface{}) response.Response {
	var body response.Response
	err := json.Unmarshal(actualResponse, &body)

	assert.Nil(t, err, "Error when trying to unmarshal response")
	assert.NotNil(t, body.Data, "Your response data is should not be nil or empty")

	actualBytes, err := json.Marshal(body.Data)
	assert.Nil(t, err, "Error when trying to Marshal response data")

	expectedBytes, err := json.Marshal(expected)
	assert.Nil(t, err, "Error when trying to Marshal expected data")

	_, diffStr := jsondiff.Compare(expectedBytes, actualBytes, &jsondiff.Options{
		SkipMatches:      true,
		ChangedSeparator: " expected value is ",
	})

	assert.Empty(t, diffStr, "Your response data and expected data is difference. Check your expected and actual data again!")

	return body
}

func CheckMiradaResponse(t *testing.T, actual interface{}, expected interface{}) {
	actualBytes, err := json.Marshal(actual)
	assert.Nil(t, err, "Error when trying to Marshal response data")

	expectedBytes, err := json.Marshal(expected)
	assert.Nil(t, err, "Error when trying to Marshal expected data")

	_, diffStr := jsondiff.Compare(expectedBytes, actualBytes, &jsondiff.Options{
		SkipMatches:      true,
		ChangedSeparator: " expected value is ",
	})

	assert.Empty(t, diffStr, "Your response data and expected data is difference. Check your expected and actual data again!")
}

func MarshalToReader(value interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonData), nil
}

func AddParameters(r *http.Request, params map[string]string) *http.Request {
	ctx := chi.NewRouteContext()
	for k, v := range params {
		ctx.URLParams.Add(k, v)
	}

	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
}
