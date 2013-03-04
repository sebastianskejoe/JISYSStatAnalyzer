package JISYSStatAnalyzer

import (
    "strings"
    "strconv"
    "fmt"
)

type EventInterface interface {
    String() string
    SetTeam(t string)
    SetTime(t string)
    Set([]string)
    Team() string
    Type() int
}

type event struct {
    typ int
    team string
    minute int
    second int
}

func (e *event) SetTeam(t string) {
    e.team = t;
}

func (e *event) Team() string {
    return e.team
}

func (e *event) Type() int {
    return e.typ
}

func (e *event) SetTime(t string) {
    parts := strings.Split(t, ".")
    e.minute,_ = strconv.Atoi(parts[0])
    e.second,_ = strconv.Atoi(parts[1])
}

type Player struct {
    Name string
    Number int
}

func (p *Player) set(parts []string) {
    p.Name = parts[1]
    p.Number,_ = strconv.Atoi(parts[0])
}

func (p Player) String() string {
    return fmt.Sprintf("#%d %s", p.Number, p.Name)
}

type Position int

const (
    VF = Position(iota)
    VB
    PM
    HB
    HF
    ST
    M7
    F1
    F2
    G1
    G2
    G3
    G4
    A1
    A2
)

func GetPosition(str string) Position {
    switch str {
    case "VF", "V6":
        return VF
    case "VB", "V9":
        return VB
    case "PM", "M9":
        return PM
    case "HB", "H9":
        return HB
    case "HF", "H6":
        return HF
    case "ST", "M6":
        return ST
    case "1.f":
        return F1
    case "2.f":
        return F2
    case "Gbr(1)":
        return G1
    case "Gbr(2)":
        return G2
    case "Gbr(3)":
        return G3
    case "Gbr(4)":
        return G4
    case "1:a":
        return A1
    case "2:a":
        return A2
    }
    panic(str)
}

func (p Position) String() string {
    switch p {
    case VF:
        return "Venstre fløj"
    case VB:
        return "Venstre back"
    case PM:
        return "Playmaker"
    case HB:
        return "Højre back"
    case HF:
        return "Højre fløj"
    case ST:
        return "Streg"
    case F1:
        return "1. fasekontra"
    case F2:
        return "2. fasekontra"
    case G1:
        return "Gennembrud Gbr(1)"
    case G2:
        return "Gennembrud mellem 1'er og 2'er"
    case G3:
        return "Gennembrud mellem 2'er og 3'er"
    case G4:
        return "Gennembrud Gbr(4)"
    case A1:
        return "A? 1. fasekontra"
    case A2:
        return "A? 2. fasekontra"
    }
    panic("Never gets here")
}

const (
    SavedShotType = iota
    BlockedShotType
    MissedShotType
    PostShotType
    GoalType
    PenaltyType
    PenaltySavedType
    FoulType
    LostBallType
    MissedPassType
)

func MakeEvent(line string) EventInterface {
    parts := strings.Split(line, ",")
    var ev EventInterface
    switch parts[3] {
    case "Skud blokeret":
        ev = new(BlockedShot)
    case "Mål":
        ev = new(Goal)
    case "Skud reddet":
        ev = new(SavedShot)
    case "Skud forbi":
        ev = new(MissedShot)
    case "Skud på stolpe":
        ev = new(PostShot)
    case "Mål på straffekast":
        ev = new(PenaltyGoal)
    case "Straffekast reddet":
        ev = new(PenaltySaved)
    case "Regelfejl":
        ev = new(Foul)
    case "Tabt bold":
        ev = new(LostBall)
    case "Missede aflevering":
        ev = new(MissedPass)
    case "Gult kort":
        ev = new(YellowCard)
    case "Udvisning":
        ev = new(TwoMinutes)
    default:
//        fmt.Println(parts[3])
        return nil
    }

    ev.SetTime(parts[1])
    ev.SetTeam(parts[2])
    ev.Set(parts)

    return ev
}
