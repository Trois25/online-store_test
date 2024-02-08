package main

import (
	"fmt"
	"store/app/configs"
	"store/app/database"
	"store/app/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	cfg := configs.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	database.InitMigration(dbMysql)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	route.New(e, dbMysql)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
