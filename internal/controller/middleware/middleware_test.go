package middleware

import (
	"net/http"
	"net/http/httptest"
	"simulator/configs"
	"simulator/internal/controller/client"
	"simulator/internal/controller/db"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDBLogger(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/start", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a real db.DB instance for testing with an in-memory SQLite database
	testConfig := configs.Config{
		Database: configs.DatabaseConfig{
			Driver: "sqlite3",
			Source: "file:ent?mode=memory&cache=shared&_fk=1",
		},
	}
	realDB, err := db.NewDB(testConfig)
	assert.NoError(t, err)
	defer realDB.Client.Close()

	// Create a dummy client.Client instance to satisfy the DBLogger signature
	dummyClient, _ := client.NewClient(configs.Config{})

	h := DBLogger(realDB, dummyClient)(func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	err = h(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}