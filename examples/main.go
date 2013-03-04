package main

import (
//    "strings"
    "os"
    "fmt"
    "log"
    "bufio"
    "bytes"
    JSA "github.com/sebastianskejoe/JISYSStatAnalyzer"
    "path/filepath"

    "flag"
    "strings"
)

const (
    SHOT = iota
    TURNOVER
    OTHER
)

var csvPath = flag.String("csv", "", "Path to csv-files")
var format = flag.String("format", "O,S,M,S-M,S-M/S,G", "Output format")

func main() {
    flag.Parse()

    args := flag.Args()

    // Open file for reading
    if len(args) < 1 {
        fmt.Printf("Usage: %s [flags] <path>\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(0)
    }

    stats := JSA.NewStats()

    filenames, err := filepath.Glob(args[0])
    if err != nil {
        fmt.Println("Unknown filepath - ", err)
        return
    }

    if len(args) > len(filenames) {
        filenames = args
    }

    for _, filename := range filenames {
        file,err := os.Open(filename)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Analyzing %s ... ", filename)

        stats.AddGame(filename[0:3], filename[3:6])

        // Read lines
        reader := bufio.NewReader(file)
        line, err := reader.ReadBytes('\n')
        for ; err == nil ; line, err = reader.ReadBytes('\n') {
            // Convert line to utf8
            buf := new(bytes.Buffer)
            for _,c := range line[:len(line)-2] {
                buf.WriteRune(rune(c))
            }

            if buf.String() == "" {
                continue
            }

            ev := JSA.MakeEvent(buf.String())

            if ev == nil {
                continue
            }

            stats.AddEvent(ev)
        }

        fmt.Println("OK!")
    }

    for key := range stats.Teams {
        var output *os.File
        var sep string

        // Either print to stdout or a file
        if *csvPath == "" {
            output = os.Stdout
            sep = "\t"
        } else {
            path := fmt.Sprintf("%s/%s.csv", *csvPath, key)
            output, err = os.Create(path)
            if err != nil {
                panic(err)
            }
            sep = ","
        }

        // Print output
        parts := strings.Split(*format, ",")
        fmt.Fprintln(output,"Stats for",key)
        fmt.Fprintln(output,"========")
        fmt.Fprintln(output, strings.Replace(*format, ",", sep, -1))

        s := sortTeamStats(stats.Teams[key])
        for _, gs := range s {
            for _,p := range parts {
                fmt.Fprintf(output, "%s%s", solveExpr(p, gs), sep)
            }
            fmt.Fprintln(output)
        }
        fmt.Fprintln(output)

        if *csvPath != "" {
            output.Close()
        }
    }
}

func sortTeamStats(stats []*JSA.GameStats) []*JSA.GameStats {
    ret := make([]*JSA.GameStats, len(stats))
    ret[0] = stats[0]
outer:
    for i := 1 ; i < len(stats) ; i++ {
        for j := 0 ; j < len(ret) ; j++ {
            if ret[j] != nil && stats[i].Opponent < ret[j].Opponent {
                copy(ret[j+1:], ret[j:])
                ret[j] = stats[i]
                continue outer
            } else if ret[j] == nil {
                ret[j] = stats[i]
                continue outer
            }
        }
    }

    return ret
}

func solveExpr(expr string, stats *JSA.GameStats) string {
    val := float64(0)
    var tmp int
    op := uint8('+')
    last := uint8('-')
    frac := false

    for pos := 0 ; pos < len(expr) ; pos++ {
        switch expr[pos] {
        case 'O':
            if pos+1 != len(expr) {
                break
            }
            if pos != 0 {
                panic("Opponent must be first in expression")
            }
            return stats.Opponent
        case 'G':
            if last == 'O' {
                tmp = stats.GoalsAgainst
            } else {
                tmp = stats.Goals
            }
        case 'M':
            tmp = stats.Missed
        case 'S':
            tmp = len(stats.Shots)
        case 'A':
            tmp = stats.Attacks
        case 'T':
            tmp = stats.Turnovers
        case 'B':
            tmp = stats.Blocked
        case 'V': // Saves
            if last == 'O' {
                tmp = stats.Saved
            } else {
                tmp = stats.Saves
            }
        default:
            val = doOp(val, op, tmp)
            op = expr[pos]
            if op == '/' {
                frac = true
            }
        }
        last = expr[pos]
    }
    val = doOp(val, op, tmp)

    if frac {
        return fmt.Sprintf("%.3f", val)
    }
    return fmt.Sprintf("%.0f", val)
}

func doOp(a float64, op uint8, b int) float64 {
    bf := float64(b)
    switch op {
    case '+':
        return a+bf
    case '-':
        return a-bf
    case '*':
        return a*bf
    case '/':
        return a/bf
    }
    return 0
}
