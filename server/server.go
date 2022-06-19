package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {

    LoadStore()

    e := echo.New()
    e.Use(middleware.CORS())
    
    e.GET("/", func(c echo.Context) error{
        return c.String(http.StatusOK, "Poker Evaluator API")
    })
    e.POST("/rank-hand", RankHandView)
    e.POST("/rank-table", RankTableView)

    e.Logger.Fatal(e.Start(":1323"))

}
