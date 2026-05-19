package main

import (
    "fmt"
)

func main() {
    // 1. SYNTAXFEHLER FIX
    fmt.Println("Beschwörung des Debugger-Auges gestartet...")

    // 2. LAUFZEITFEHLER FIX
    var DebuggerAuge *string

    // Nil-Prüfung
    if DebuggerAuge == nil {
        fmt.Println("Debugger-Auge ist nicht verfügbar (nil)")
    } else {
        fmt.Println("Aktiviere Kraft von:", *DebuggerAuge)
    }

    // Division durch Null verhindern
    energieQuelle := 100
    aktiveGatter := 0

    if aktiveGatter != 0 {
        lastProGatter := energieQuelle / aktiveGatter
        fmt.Println("Lastverteilung:", lastProGatter)
    } else {
        fmt.Println("Keine aktiven Gatter → keine Berechnung möglich")
    }

    // 3. LOGIKFEHLER FIX
    fmt.Println("Aktiviere Schutzsiegel...")
    siegelZaehler := 1

    for siegelZaehler <= 3 {
        fmt.Printf("Siegel Nr. %d fixiert.\n", siegelZaehler)
        siegelZaehler++
    }

    fmt.Println("Die Kathedrale ist gereinigt!")
}
