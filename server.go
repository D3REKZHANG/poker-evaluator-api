package main

import (
    "net/http"
    "os"
    "bufio"
    "fmt"
    "strconv"
    "github.com/labstack/echo/v4"
)

type Hand struct {
    Cards [2]string `json:"cards"`
}

func rankHand(c echo.Context) error {
    return c.JSON(http.StatusOK, &Hand{Cards: [2]string{"2h", "4s"}})
}

func load_store() map[string]int {
    s := make(map[string]int)
    
    file, err := os.Open("./store.txt")

    if err != nil {
        panic(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Split(bufio.ScanWords) 

    for scanner.Scan() {
        key := scanner.Text()
        scanner.Scan()
        val, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }
        s[key] = val
    }

    return s
}

func main() {

    e := echo.New()

    store := load_store()

    e.GET("/", func(c echo.Context) error{
        return c.String(http.StatusOK, "Poker Evaluator API")
    })
    e.POST("/rankHand", rankHand)

    e.Logger.Fatal(e.Start(":1323"))

}
