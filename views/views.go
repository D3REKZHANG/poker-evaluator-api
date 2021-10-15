package views
import (
    "net/http"
    "github.com/labstack/echo/v4"
    "local/business"
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
    h := new(Hand)
    if err := c.Bind(h); err != nil {
        return err
    }

    result := ScoreRes{
        Score: business.RankHand(h.Cards),
    }
    return c.JSON(http.StatusOK, result)
}

func RankTable(c echo.Context) (err error) {
    t := new(Table)
    if err := c.Bind(t); err != nil {
        return err
    }
    winner, score, hand := business.RankTable(t.River, t.Holes)
    result := TableRes{ 
        Winner: winner,
        Hand: hand,
        Score: score,
    }
    return c.JSON(http.StatusOK, result)
}