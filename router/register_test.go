package router

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/shoriwe/message-api/models"
)

func testRegister(t *testing.T, r *Router, expect *httpexpect.Expect) {
	t.Run("Valid", func(tt *testing.T) {
		u := models.RandomUser()
		expect.
			PUT(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated)
	})
	t.Run("Invalid JSON", func(tt *testing.T) {
		expect.
			PUT(RegisterRoute).
			WithBytes([]byte("[")).
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("Repeated User", func(tt *testing.T) {
		u := models.RandomUser()
		expect.
			PUT(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusCreated)
		expect.
			PUT(RegisterRoute).
			WithJSON(u).
			Expect().
			Status(http.StatusInternalServerError)
	})
}

func TestRegist(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		serverClose, r, expect := NewTest(tt)
		defer serverClose()
		defer r.Close()
		testRegister(tt, r, expect)
	})
}
