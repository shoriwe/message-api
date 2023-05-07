package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/common/http_utils"
	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testSendMessage(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Valid", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		u2 := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u2).Error)
		d2 := models.RandomDevice(u2)
		assert.Nil(tt, r.Controller.DB.Create(d2).Error)
		expect.
			PUT(MessageRoute).
			WithHeader("Authorization", http_utils.ToBearer(r.Controller.JWT.New(s))).
			WithJSON(models.Message{
				RecipientUUID: u2.UUID,
				Title:         random.String(),
				Body:          random.String(),
			}).
			Expect().
			Status(http.StatusCreated)
	})
	t.Run("Invalid JSON", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		expect.
			PUT(MessageRoute).
			WithHeader("Authorization", http_utils.ToBearer(r.Controller.JWT.New(s))).
			WithBytes([]byte("[")).
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testSendMessage(tt, r, expect)
	})
}
