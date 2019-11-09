package account_test

import (
	"gin-framework/app"
	router "gin-framework/app/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// init routes
	router.InitAccountRouter(app.Engine)
	os.Exit(m.Run())
}
func TestGetAccount(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/account/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "got account 1", rep.Body.String())
}

func TestCreateAccount(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/account", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "creating account", rep.Body.String())
}

func TestUpdateAccount(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/account/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "updating account 1", rep.Body.String())
}

func TestDeleteAccount(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/account/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "deleting account 1", rep.Body.String())
}
