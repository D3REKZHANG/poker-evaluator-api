package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type Hand struct {
    Cards string `json:"cards"`
}

func rankHand(c echo.Context) error {
    u := &Hand{
        Cards: "2h",
    }
    return c.JSON(http.StatusOK, u)
}

func main() {

    e := echo.New()
    e.GET("/", func(c echo.Context) error{
        return c.String(http.StatusOK, "Poker Evaluator API")
    })
    e.GET("/rankHand", rankHand)
    e.Logger.Fatal(e.Start(":1323"))

}
