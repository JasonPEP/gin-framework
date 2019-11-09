package book_test

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
	router.InitBookRouter(app.Engine)
	os.Exit(m.Run())
}

func TestCreateBook(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/book", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "creating book", rep.Body.String())
}

func TestDeleteBook(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/book/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "deleting book 1", rep.Body.String())
}

func TestGetBook(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/book/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "got book 1", rep.Body.String())
}

func TestUpdateBook(t *testing.T) {
	rep := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/book/1", nil)
	app.Engine.ServeHTTP(rep, req)

	assert.Equal(t, 200, rep.Code)
	assert.Equal(t, "updating book 1", rep.Body.String())
}
