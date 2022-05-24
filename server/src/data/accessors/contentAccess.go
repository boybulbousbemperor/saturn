package accessors

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boybulbousbemperor/go-saturn/src/factory"
	models "github.com/boybulbousbemperor/go-saturn/src/models/services"
	"github.com/stretchr/testify/assert"
)

type ContentKeymap struct {
	guid map[string]int
	key  map[string]string
}

func Create() *ContentKeymap {
	return &ContentKeymap{map[string]int{}, map[string]string{}}
}

func (ckm *ContentKeymap) ReplaceGUID(toreplace string, replacingwith int) {
	ckm.guid[toreplace] = replacingwith
}

func (ckm *ContentKeymap) ReplaceKey(toreplace string, replacingwith string) {
	ckm.key[toreplace] = replacingwith
}

func TestGppdApiContent(t *testing.T) {
	cont := new(factory.ContentRepository)
	serv := models.ContentService{
		Repository: *cont,
	}

	t.Run("Must return HTTP Response Code: 200, in \"api/\"", func(t *testing.T) {
		reqs, err := http.NewRequest(http.MethodGet, "api/", nil)
		resp := httptest.NewRecorder()

		if err != nil {
			log.Fatalln(err)
		}

		serv.EncodingServiceGPPD(resp, reqs)

		assert.Equal(t, resp.Code, http.StatusOK, "Response Code and HTTP Status of OK: should be the same.")
	})
}
