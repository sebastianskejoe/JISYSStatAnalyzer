package JISYSStatAnalyzer

type Stats struct {
    Teams map[string][]*GameStats
    CurrentGame int
    GameCount int
}

type GameStats struct {
    Shots []EventInterface
    Events []EventInterface
    Attacks int
    Goals int
    GoalsAgainst int
    Missed int
    Blocked int
    Saved int
    Saves int // own saves
    Turnovers int
    Opponent string
}

func NewStats() *Stats {
    s := new(Stats)
    s.CurrentGame = -1
    s.Teams = make(map[string][]*GameStats)
    return s
}

func NewGameStats(opp string) *GameStats {
    gs := new(GameStats)

    gs.Shots = make([]EventInterface, 0)
    gs.Events = make([]EventInterface, 0)
    gs.Opponent = opp

    return gs
}

func (s *Stats) AddGame(t1, t2 string) {
    s.CurrentGame++
    s.GameCount++
    s.Teams[t1] = append(s.Teams[t1], NewGameStats(t2))
    s.Teams[t2] = append(s.Teams[t2], NewGameStats(t1))
}

func (s *Stats) AddShot(ev EventInterface) {
    gs := s.getCurrent(ev.Team())
    op := s.getCurrent(gs.Opponent)
    gs.Shots = append(gs.Shots, ev)
    gs.Attacks++

    switch ev.Type() {
    case GoalType, PenaltyType:
        gs.Goals++
        op.GoalsAgainst++
    case SavedShotType, PenaltySavedType:
        gs.Saved++
        op.Saves++
    case BlockedShotType:
        gs.Blocked++
    case MissedShotType, PostShotType:
        gs.Missed++
    }
}

func (s *Stats) AddTurnover(ev EventInterface) {
    gs := s.getCurrent(ev.Team())
    gs.Turnovers++
    gs.Attacks++
}

func (s *Stats) AddEvent(ev EventInterface) {
    gs := s.getCurrent(ev.Team())

    switch ev.Type() {
    case GoalType, PenaltyType, SavedShotType, PenaltySavedType, BlockedShotType, MissedShotType, PostShotType:
        s.AddShot(ev)
    case FoulType, LostBallType, MissedPassType:
        s.AddTurnover(ev)
    default:
    }
    gs.Events = append(gs.Events, ev)
}

func (s *Stats) getCurrent(team string) *GameStats {
    if len(s.Teams[team]) < 1 {
        panic("No games added yet!")
    }
    return s.Teams[team][len(s.Teams[team])-1]
}
