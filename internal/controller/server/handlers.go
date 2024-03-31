package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func startHandler(c echo.Context) error {
	log.Info("start handler")

	return nil
}

func stopHandler(c echo.Context) error {
	log.Info("stop handler")

	return nil
}

func leftHandler(c echo.Context) error {
	log.Info("left handler")

	return nil
}

func rightHandler(c echo.Context) error {
	log.Info("right handler")

	return nil
}

func forwardHandler(c echo.Context) error {
	log.Info("forward handler")

	return nil
}

func backHandler(c echo.Context) error {
	log.Info("back handler")

	return nil
}
