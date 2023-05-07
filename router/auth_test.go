package router

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/http_utils"
	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testAuthenticate(t *testing.T, r *Router, expect *httpexpect.Expect) {
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
	t.Run("Unauthorized", func(tt *testing.T) {
		s := session.NewSession(uuid.NewV4(), "YAY")
		expect.
			POST(LogoutRoute).
			WithHeader("Authorization", http_utils.ToBearer(r.JWT.New(s))).
			Expect().
			Status(http.StatusUnauthorized)
	})
	t.Run("Invalid B64", func(tt *testing.T) {
		expect.
			POST(LogoutRoute).
			WithHeader("Authorization", http_utils.ToBearer("???")).
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("Invalid JWT", func(tt *testing.T) {
		expect.
			POST(LogoutRoute).
			WithHeader("Authorization", http_utils.ToBearer(base64.StdEncoding.EncodeToString([]byte("X")))).
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestAuthenticate(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testAuthenticate(tt, r, expect)
	})
}
