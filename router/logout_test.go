package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/common/http_utils"
	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func testLogout(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Authorized", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, r.Controller.DB.Create(u).Error)
		_, token, lErr := r.Controller.Login(&controller.Login{Email: *u.Email, Password: *u.Password, FirebaseToken: random.String()})
		assert.Nil(tt, lErr)
		expect.
			POST(LogoutRoute).
			WithHeader("Authorization", http_utils.ToBearer(token)).
			Expect().
			Status(http.StatusOK)
	})
}

func TestLogout(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testLogout(tt, r, expect)
	})
}
