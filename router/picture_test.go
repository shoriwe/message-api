package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/common/http_utils"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testDownloadPicture(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Valid", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		token := r.Controller.JWT.New(session.NewSession(u.UUID, d.FirebaseToken))
		expect.
			GET(PictureRoute+u.UUID.String()).
			WithHeader("Authorization", http_utils.ToBearer(token)).
			Expect().
			Status(http.StatusOK).
			Body().
			IsEqual(models.TestingImage)
	})
	t.Run("FAKE UUID", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		token := r.Controller.JWT.New(session.NewSession(u.UUID, d.FirebaseToken))
		expect.
			GET(PictureRoute+"FAKE_UUID").
			WithHeader("Authorization", http_utils.ToBearer(token)).
			Expect().
			Status(http.StatusUnauthorized)
	})
}

func TestDownloadPicture(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testDownloadPicture(tt, r, expect)
	})
}
