package business

import (
    "os"
    "strconv"
    "bufio"
    "fmt"
    "local/helpers"
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
    key := ""
    prevSuit := hand[0][1]
    flush := true
    for _, card := range hand {
        if flush && card[1] != prevSuit {
            flush = false
        }
        prevSuit = card[1]

        key += string(card[0]) // potential speed issues
    }

    if flush {
        key += "f"
    }

    key = helpers.SortString(key)

    return store[key]
}
