package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/common/http_utils"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testUsers(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Valid", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		// Init users
		for i := 0; i < 10; i++ {
			u2 := models.RandomUser()
			assert.Nil(tt, r.Controller.DB.Create(u2).Error)
		}
		//
		expect.
			POST(UsersRoute).
			WithHeader("Authorization", http_utils.ToBearer(r.Controller.JWT.New(session.NewSession(u.UUID, d.FirebaseToken)))).
			WithJSON(controller.Filter[models.User]{}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array().
			Length().IsEqual(10)
	})
	t.Run("Invalid JSON", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, r.Controller.DB.Create(d).Error)
		expect.
			POST(UsersRoute).
			WithHeader("Authorization", http_utils.ToBearer(r.Controller.JWT.New(session.NewSession(u.UUID, d.FirebaseToken)))).
			WithBytes([]byte("[")).
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestUsers(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testUsers(tt, r, expect)
	})
}
