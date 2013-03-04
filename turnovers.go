package JISYSStatAnalyzer

import (
    "fmt"
)

/***********
 *  FOUL   *
 **********/
type Foul struct {
    event
    Player Player
}

func (f *Foul) Set(args []string) {
    f.typ = FoulType
    f.Player.set(args[5:7])
}

func (f *Foul) String() string {
    str := fmt.Sprintf("FOUL(%s): %s", f.Team(), f.Player)
    return str
}

/***********
 *  LOST   *
 **********/
type LostBall struct {
    event
    Player Player
    Steal bool
    Stealer Player
}

func (lb *LostBall) Set(args []string) {
    lb.typ = LostBallType
    lb.Player.set(args[5:7])
    if args[7] == "Bolderobring" {
        lb.Steal = true
        lb.Stealer.set(args[8:10])
    }
}

func (lb *LostBall) String() string {
    str := fmt.Sprintf("LOST BALL(%s): %s", lb.Team(), lb.Player)
    if lb.Steal {
        str = fmt.Sprintf("%s STEAL %s", str, lb.Stealer)
    }
    return str
}

/***********
 * MISSED P*
 **********/
type MissedPass struct {
    event
    Player Player
    Steal bool
    Stealer Player
}

func (mp *MissedPass) Set(args []string) {
    mp.typ = MissedPassType
    mp.Player.set(args[5:7])
    if args[7] == "Bolderobring" {
        mp.Steal = true
        mp.Stealer.set(args[8:10])
    }
}

func (mp *MissedPass) String() string {
    str := fmt.Sprintf("MISSED PASS(%s): %s", mp.Team(), mp.Player)
    if mp.Steal {
        str = fmt.Sprintf("%s STEAL %s", str, mp.Stealer)
    }
    return str
}
