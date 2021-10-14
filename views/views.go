package views
import (
    "net/http"
    "github.com/labstack/echo/v4"
    "local/business"
    "time"
    "fmt"
)

// DTOs ---------------------------------------------------
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

type TableRes struct {
    Winner int `form:"winner" json:"winner"`
    Hand [5]string `form:"hand" json:"hand"`
    Score int `form:"score" json:"score"`
}
// --------------------------------------------------------

func RankHand(c echo.Context) (err error) {
    start := time.Now()
    h := new(Hand)
    if err := c.Bind(h); err != nil {
        return err
    }
    var score int
    for i:=0;i<100000;i++{
        score = business.RankHand(h.Cards) 
    }
    fmt.Println(time.Since(start))
    result := ScoreRes{
        Score: score,
    }
    return c.JSON(http.StatusOK, result)
}

func RankTable(c echo.Context) (err error) {
    start := time.Now()
    t := new(Table)
    if err := c.Bind(t); err != nil {
        return err
    }
    var winner int
    var score int
    var hand [5]string
    for i:=0;i<100000;i++{
        winner, score, hand = business.RankTable(t.River, t.Holes)
    }
    fmt.Println(time.Since(start))
    result := TableRes{ 
        Winner: winner,
        Hand: hand,
        Score: score,
    }
    return c.JSON(http.StatusOK, result)
}
