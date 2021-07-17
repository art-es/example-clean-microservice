package http_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/art-es/example-clean-microservice/internal/infrastructure/httphandler"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func newTestServer() *httptest.Server {
	mux, err := httphandler.NewServeMux()
	if err != nil {
		log.Panicf("[PANIC] newTestServer: %v\n", err)
	}
	return httptest.NewServer(mux)
}

func mustDoRequest(req *http.Request, err error) *http.Response {
	if err != nil {
		log.Panicf("[PANIC] mustDoRequest: %v\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("[PANIC] mustDoRequest: [%s %s] http.Do: %v\n", req.Method, req.URL, err)
	}
	return res
}

func mustRead(r io.ReadCloser) []byte {
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		log.Panicf("[PANIC] mustRead: %v\n", err)
	}
	return data
}

func mustJSONEncode(i interface{}) []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Panicf("[PANIC] mustJSONEncode: %v\n", err)
	}
	return data
}

type H map[string]interface{}

func matchJSON(h H) types.GomegaMatcher {
	return gomega.MatchJSON(mustJSONEncode(h))
}
