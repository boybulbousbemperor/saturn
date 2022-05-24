package handlers

import (
	"net/http"
	"testing"

	models "github.com/boybulbousbemperor/go-saturn/src/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type error_handle_gppd_i interface {
	handle_error() error
}

type offset_error_boundary_t struct {
	offset uint64
}

func test_off_bounds_error(t *testing.T) {
	var offs offset_error_boundary_t
	require.Equal(t, uint64(1), offs.offset)
}

/* func (offs *offset_error_boundary_t) procedure_call_status() *status.Status {
	stat := status.New(
		404,
		fmt.Sprintf("Error's offset out of bounds: %d", offs.offset),
	)

	message := fmt.Sprintf("Offset boundary for error: %d", offs.offset)

	details := &errdetails.LocalizedMessage{
		Locale:  "en-US",
		Message: message,
	}

	std, err := stat.WithDetails(details)
	if err != nil {
		return stat
	}

	return std
} */

type ParamtersAPI interface {
	Params(models.APIParameters) struct{}
}

type criterion_t struct {
	params *models.APIParameters
	ctx    *gin.Context
}

var apiparams = &[]models.APIParameters{
	{Controller: "home", Action: "index", ID: "id"},
}

func (crit *criterion_t) get_criterion() {
	cont := crit.params.Controller
	act := crit.params.Action
	id := crit.params.ID

	crit.ctx.IndentedJSON(http.StatusOK, apiparams)

	for _, a := range *apiparams { // uncopy lock or transform `a` with a mutex?
		if a.Controller == cont {
			crit.ctx.IndentedJSON(http.StatusOK, &a)
		}

		if a.Action == act {
			crit.ctx.IndentedJSON(http.StatusOK, &a)
		}

		if a.ID == id {
			crit.ctx.IndentedJSON(http.StatusOK, &a)
		}
	}
}

func (crit *criterion_t) PostContent() {
	var newViews []models.APIParameters

	if err := crit.ctx.BindJSON(&newViews); err != nil {
		return
	}

	apiparams = &newViews
	crit.ctx.IndentedJSON(http.StatusCreated, newViews)
}

/* func StreamConsumer(req *pb.APIParameters, stream pb.ResourcingHandler) error */
