package middleware

import (
	"context"
	"net/http"
	"simulator/ent/record"
	"simulator/internal/controller/client"
	"simulator/internal/controller/db"

	"github.com/labstack/echo/v4"
)

func DBLogger(db *db.DB, client *client.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			_, err := db.Record.
				Create().
				SetHTTPMethod(record.HTTPMethod(req.Method)).
				SetIPAddress(req.RemoteAddr).
				SetInstruction(req.RequestURI[1:]).
				Save(context.Background())
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// err = client.Send(api.NewMessage(r.IPAddress, r.Instruction))
			// if err != nil {
			// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			// }

			return c.String(http.StatusOK, "OK")
		}
	}
}
