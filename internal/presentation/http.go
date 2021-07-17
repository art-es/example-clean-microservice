package presentation

import (
	"net/http"
	"strconv"

	"github.com/art-es/example-clean-microservice/internal/service"
	. "github.com/art-es/example-clean-microservice/internal/util/http"
)

type HTTPPresenter struct {
	CalculationService *service.CalculationService
}

// ActionDoubling is an endpoint for calculating a double number of the specified number.
// Param 'num' is required, must be a float number, the decimal separator is the dot ".".
func (p HTTPPresenter) ActionDoubling(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		MethodNotAllowed(w)
		return
	}

	_ = r.ParseForm() // When a method is GET r.ParseForm never throws an error.

	rawnum := r.FormValue("num")
	if rawnum == "" {
		Error(w, http.StatusBadRequest, "parameter 'num' is required")
		return
	}

	num, err := strconv.ParseFloat(rawnum, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, "parameter 'num' must be a number")
		return
	}

	result := p.CalculationService.Doubling(num)

	Respond(w, http.StatusOK, map[string]interface{}{
		"result": result,
	})
}
