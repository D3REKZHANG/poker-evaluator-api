package business

import (
    "os"
    "strconv"
    "bufio"
    "fmt"
    "local/helpers"
    "strings"
)

var store map[string]int

func LoadStore(){
    s := make(map[string]int)
    file, err := os.Open("../store.txt")

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
    store = s
}

func RankHand(hand [5]string) int {
    var key strings.Builder
    prevSuit := hand[0][1]
    flush := true
    for _, card := range hand {
        if flush && card[1] != prevSuit {
            flush = false
        }
        prevSuit = card[1]

        key.WriteByte(card[0])
    }

    if flush {
        key.WriteByte('f')
    }

    sKey := helpers.SortString(key.String())

    return store[sKey]
}

func RankPlayer(river [5]string, hole [2]string) (int, [5]string){
    bestScore := RankHand(river)
    bestHand := river

    // 3 river cards 2 hole card
    for a:=0;a<5;a++{
        for b:=a+1;b<5;b++{
            for c:=b+1;c<5;c++{
                hand := [5]string{river[a],river[b], river[c], hole[0], hole[1]}
                score := RankHand(hand)
                if score > bestScore{
                    bestScore = score
                    bestHand = hand
                }
            }
        }
    }

    // 4 river cards 1 hole card
    for a:=0;a<5;a++{
        for b:=a+1;b<5;b++{
            for c:=b+1;c<5;c++{
                for d:=c+1;d<5;d++{
                    for h:=0;h<2;h++ {
                        hand := [5]string{river[a],river[b], river[c], river[d], hole[h]}
                        score := RankHand(hand)
                        if score > bestScore{
                            bestScore = score
                            bestHand = hand
                        }
                    }
                }
            }
        }
    }

    return bestScore, bestHand
}

func RankTable(river [5]string, holes [][2]string) (int, int, [5]string){
    var bestHand [5]string
    winner := 0
    bestScore := 0

    // For each player:
    for i, h := range holes {
        score, hand := RankPlayer(river, h)
        if score > bestScore{
            bestScore = score
            bestHand = hand
            winner = i
        }
    }
    return winner, bestScore, bestHand
}
