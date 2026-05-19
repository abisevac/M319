package main

import (
    "fmt"
)

func main() {
    // 1. SYNTAXFEHLER FIX
    fmt.Println("Initialisiere Heilungs-Protokoll...")

    // 2. LAUFZEITFEHLER FIX
    var heilungsKristall *int

    if heilungsKristall == nil {
        fmt.Println("Kein Kristall vorhanden. Energie = 0")
    } else {
        fmt.Printf("Energie-Level des Kristalls: %d\n", *heilungsKristall)
    }

    // 3. LOGIKFEHLER FIX
    energieFluss := 40
    istHeilungErfolgreich := false

    if energieFluss >= 50 {
        istHeilungErfolgreich = true
    }

    if istHeilungErfolgreich {
        fmt.Println("✅ Die Aura breitet sich aus! Die Highlands erblühen.")
    } else {
        fmt.Println("❌ Die Energie reicht nicht aus. Die Highlands bleiben grau.")
    }

    // 4. SCHLEIFEN-FEHLER FIX
    baeumeGepflanzt := 0
    fmt.Println("Beginne Aufforstung...")

    for baeumeGepflanzt < 5 {
        baeumeGepflanzt++
        fmt.Printf("Baum Nr. %d gepflanzt.\n", baeumeGepflanzt)
    }

    fmt.Println("Mission abgeschlossen: Der Weltencode ist rein!")
}
