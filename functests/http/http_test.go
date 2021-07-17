package http_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/gomega"
)

func TestActionDoubling(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	t.Run("no params", func(t *testing.T) {
		res := mustDoRequest(http.NewRequest(http.MethodGet, ts.URL+"/actions/doubling", nil))

		g := NewGomegaWithT(t)
		g.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
		g.Expect(mustRead(res.Body)).Should(matchJSON(H{"error": "parameter 'num' is required"}))
	})

	t.Run("num param is string", func(t *testing.T) {
		res := mustDoRequest(http.NewRequest(http.MethodGet, ts.URL+"/actions/doubling?num=foo", nil))

		g := NewGomegaWithT(t)
		g.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
		g.Expect(mustRead(res.Body)).Should(matchJSON(H{"error": "parameter 'num' must be a number"}))
	})

	t.Run("num param is float number, dot separator", func(t *testing.T) {
		res := mustDoRequest(http.NewRequest(http.MethodGet, ts.URL+"/actions/doubling?num=1.5", nil))

		g := NewGomegaWithT(t)
		g.Expect(res.StatusCode).Should(Equal(http.StatusOK))
		g.Expect(mustRead(res.Body)).Should(matchJSON(H{"result": 3}))
	})

	t.Run("num param is positive integer number", func(t *testing.T) {
		res := mustDoRequest(http.NewRequest(http.MethodGet, ts.URL+"/actions/doubling?num=1", nil))

		g := NewGomegaWithT(t)
		g.Expect(res.StatusCode).Should(Equal(http.StatusOK))
		g.Expect(mustRead(res.Body)).Should(matchJSON(H{"result": 2}))
	})

	t.Run("num param is negative integer number", func(t *testing.T) {
		res := mustDoRequest(http.NewRequest(http.MethodGet, ts.URL+"/actions/doubling?num=-1", nil))

		g := NewGomegaWithT(t)
		g.Expect(res.StatusCode).Should(Equal(http.StatusOK))
		g.Expect(mustRead(res.Body)).Should(matchJSON(H{"result": -2}))
	})
}
