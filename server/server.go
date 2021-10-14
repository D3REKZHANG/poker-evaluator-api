package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "local/views"
    "local/business"
)

func main() {

    business.LoadStore()

    e := echo.New()
    
    e.GET("/", func(c echo.Context) error{
        return c.String(http.StatusOK, "Poker Evaluator API")
    })
    e.POST("/rankHand", views.RankHand)
    e.POST("/rankTable", views.RankTable)

    e.Logger.Fatal(e.Start(":1323"))

}
