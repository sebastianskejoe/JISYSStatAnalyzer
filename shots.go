package JISYSStatAnalyzer

import (
    "fmt"
)

/***********
 * BLOCKED *
 **********/
type BlockedShot struct {
    event
    Shooter Player
    Position Position
    Rebound bool
    Rebounder Player
    RebounderSet bool
}

func (bs *BlockedShot) Set(args []string) {
    bs.typ = BlockedShotType
    bs.Position = GetPosition(args[4])
    bs.Shooter.set(args[5:7])
    if args[7] == "Retur" {
        bs.Rebound = true
        if args[8] != "" {
            bs.RebounderSet = true
            bs.Rebounder.set(args[8:10])
        }
    }
}

func (bs *BlockedShot) String() string {
    str := fmt.Sprintf("BLOCKED SHOT(%s): %s @ %s", bs.Team(), bs.Shooter, bs.Position)
    if bs.Rebound {
        str = fmt.Sprintf("%s REBOUNDED", str)
        if bs.RebounderSet {
            str = fmt.Sprintf("%s %s", str, bs.Rebounder)
        }
    }
    return str
}

/************
 *   GOAL   *
 ***********/
type Goal struct {
    event
    Shooter Player
    Position Position
    Assisted bool
    Assist Player
    Keeper Player
}

func (g *Goal) Set(args []string) {
    g.typ = GoalType
    g.Position = GetPosition(args[4])
    g.Shooter.set(args[5:7])
    if args[7] == "Assist" {
        g.Assisted = true
        g.Assist.set(args[8:10])
    }
    g.Keeper.set(args[10:12])
}

func (g *Goal) String() string {
    str := fmt.Sprintf("GOAL(%s): %s @ %s", g.Team(), g.Shooter, g.Position)
    if g.Assisted {
        str = fmt.Sprintf("%s, Assist by %s", str, g.Assist)
    }
    str = fmt.Sprintf("%s, Keeper: %s", str, g.Keeper)
    return str
}

/***********
 * SAVED   *
 **********/
type SavedShot struct {
    event
    Shooter Player
    Keeper Player
    Position Position
    Rebound bool
    Rebounder Player
    RebounderSet bool
}

func (ss *SavedShot) Set(args []string) {
    ss.typ = SavedShotType
    ss.Position = GetPosition(args[4])
    ss.Shooter.set(args[5:7])
    if args[7] == "Retur" {
        ss.Rebound = true
        if args[8] != "" {
            ss.RebounderSet = true
            ss.Rebounder.set(args[8:10])
        }
    }
    ss.Keeper.set(args[10:12]);
}

func (ss *SavedShot) String() string {
    str := fmt.Sprintf("SAVED SHOT(%s): %s @ %s, Keeper: %s", ss.Team(), ss.Shooter, ss.Position, ss.Keeper)
    if ss.Rebound {
        str = fmt.Sprintf("%s REBOUNDED", str)
        if ss.RebounderSet {
            str = fmt.Sprintf("%s %s", str, ss.Rebounder)
        }
    }
    return str
}

/***********
 * MISSED  *
 **********/
type MissedShot struct {
    event
    Shooter Player
    Keeper Player
    Position Position
}

func (ss *MissedShot) Set(args []string) {
    ss.typ = MissedShotType
    ss.Position = GetPosition(args[4])
    ss.Shooter.set(args[5:7])
    ss.Keeper.set(args[10:12]);
}

func (ss *MissedShot) String() string {
    str := fmt.Sprintf("MISSED SHOT(%s): %s @ %s, Keeper: %s", ss.Team(), ss.Shooter, ss.Position, ss.Keeper)
    return str
}

/***********
 *   POST  *
 **********/
type PostShot struct {
    event
    Shooter Player
    Keeper Player
    Position Position
    Rebound bool
    Rebounder Player
    RebounderSet bool
}

func (ss *PostShot) Set(args []string) {
    ss.typ = PostShotType
    ss.Position = GetPosition(args[4])
    ss.Shooter.set(args[5:7])
    ss.Keeper.set(args[10:12]);
    if args[7] == "Retur" {
        ss.Rebound = true
        if args[8] != "" {
            ss.RebounderSet = true
            ss.Rebounder.set(args[8:10])
        }
    }
}

func (ss *PostShot) String() string {
    str := fmt.Sprintf("SHOT ON POST(%s): %s @ %s, Keeper: %s", ss.Team(), ss.Shooter, ss.Position, ss.Keeper)
    if ss.Rebound {
        str = fmt.Sprintf("%s REBOUND", str)
        if ss.RebounderSet {
            str = fmt.Sprintf("%s %s", str, ss.Rebounder)
        }
    }
    return str
}

/************
 * PENALTY G*
 ***********/
type PenaltyGoal struct {
    event
    Shooter Player
    Keeper Player
}

func (g *PenaltyGoal) Set(args []string) {
    g.typ = PenaltyType
    g.Shooter.set(args[5:7])
    g.Keeper.set(args[10:12])
}

func (g *PenaltyGoal) String() string {
    str := fmt.Sprintf("PENALTY GOAL(%s): %s, Keeper: %s", g.Team(), g.Shooter, g.Keeper)
    return str
}

/************
 * PENALTY S*
 ***********/
type PenaltySaved struct {
    event
    Shooter Player
    Keeper Player
    Rebound bool
    Rebounder Player
    RebounderSet bool
}

func (ss *PenaltySaved) Set(args []string) {
    ss.typ = PenaltySavedType
    ss.Shooter.set(args[5:7])
    if args[7] == "Retur" {
        ss.Rebound = true
        if args[8] != "" {
            ss.RebounderSet = true
            ss.Rebounder.set(args[8:10])
        }
    }
    ss.Keeper.set(args[10:12]);
}

func (ss *PenaltySaved) String() string {
    str := fmt.Sprintf("PENALTY SAVED(%s): %s, Keeper: %s", ss.Team(), ss.Shooter, ss.Keeper)
    if ss.Rebound {
        str = fmt.Sprintf("%s REBOUNDED", str)
        if ss.RebounderSet {
            str = fmt.Sprintf("%s %s", str, ss.Rebounder)
        }
    }
    return str
}
