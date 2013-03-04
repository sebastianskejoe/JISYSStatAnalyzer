package JISYSStatAnalyzer

import (
    "fmt"
)

/***********
 * YELLOW  *
 **********/
type YellowCard struct {
    event
    Player Player
}

func (y *YellowCard) Set(args []string) {
    y.Player.set(args[5:7])
}

func (y *YellowCard) String() string {
    str := fmt.Sprintf("YELLOW CARD(%s): %s", y.Team(), y.Player)
    return str
}

/***********
 * 2 MINUTE*
 **********/
type TwoMinutes struct {
    event
    Player Player
}

func (y *TwoMinutes) Set(args []string) {
    y.Player.set(args[5:7])
}

func (y *TwoMinutes) String() string {
    str := fmt.Sprintf("TWO MINUTES(%s): %s", y.Team(), y.Player)
    return str
}
