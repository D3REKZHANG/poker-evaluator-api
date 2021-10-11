package main

import (
    "net/http"
    "os"
    "bufio"
    "fmt"
    "sort"
    "strconv"
    "github.com/labstack/echo/v4"
    "./helpers"
)

type Hand struct {
    Cards [5]string `form:"cards" json:"cards"`
}

type Table struct {
    River [5]string `form:"river" json:"river"`
    Holes [][2]string `form:"holes" json:"holes"`
}

type ScoreRes struct {
    Score int `form:"score" json:"score"`
}

var s map[string]int

func rankHand(hand [5]string) int {
    key := ""
    prevSuit := hand[0][1]
    flush := true
    for i, card := range hand {
        if flush && card[1] != prevSuit {
            flush = false
        }
        prevSuit = card[1]
        
        key += string(card[0]) // potential speed issues
    }

    if flush {
        key += "f"
    }

    helpers.SortString(key)

    return store[key]
}


func _rankHand(c echo.Context) (err error) {
    h := new(Hand)
    if err := c.Bind(h); err != nil {
        return err
    }
    result := ScoreRes{
        Score: rankHand(h.cards),
    }
    return c.JSON(http.StatusOK, result)
}

func load_store() map[string]int {
    
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

    s := make(map[string]int)
    e := echo.New()

    store := load_store()

    fmt.Print(store)

    e.GET("/", func(c echo.Context) error{
        return c.String(http.StatusOK, "Poker Evaluator API")
    })
    e.POST("/rankHand", _rankHand)

    e.Logger.Fatal(e.Start(":1323"))

}
