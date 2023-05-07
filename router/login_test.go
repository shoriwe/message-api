package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func testLogin(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Authorized", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.DB.Create(u).Error)
		token := random.String()
		expect.
			POST(LoginRoute).
			WithJSON(controller.Login{
				Email:         *u.Email,
				Password:      *u.Password,
				FirebaseToken: token,
			}).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("Invalid JSON", func(tt *testing.T) {
		expect.
			POST(LoginRoute).
			WithBytes([]byte("[")).
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("Unauthorized", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.DB.Create(u).Error)
		token := random.String()
		expect.
			POST(LoginRoute).
			WithJSON(controller.Login{
				Email:         "Invalid",
				Password:      *u.Password,
				FirebaseToken: token,
			}).
			Expect().
			Status(http.StatusUnauthorized)
	})
}

func TestLogin(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testLogin(tt, r, expect)
	})
}
