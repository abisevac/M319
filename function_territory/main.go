package main

import "fmt"

// TEIL 1: Basis-Auftrieb
func berechneBasisAuftrieb(windgeschwindigkeit int, anflugwinkel int) int {
    return windgeschwindigkeit + anflugwinkel
}

// TEIL 2: Energieverbrauch
func berechneEnergieVerbrauch(gewicht int, strecke int) int {
    return (gewicht * strecke) / 10
}

// TEIL 3: Flugstatus
func checkFlugStatus(auftrieb int, verbrauch int) (int, string) {
    diff := auftrieb - verbrauch

    if diff > 0 {
        return diff, "Stabil"
    }
    return diff, "Absturzgefahr"
}

// TEIL 4: Kreischen
func kreischen() {
    fmt.Println("KKKRREEEIIIISSSCCCHHHHH")
}

func main() {
    fmt.Println("--- Analyse der Variable-Falken startet ---")

    // Auftrieb
    speed := 45
    winkel := 12
    auftrieb := berechneBasisAuftrieb(speed, winkel)
    fmt.Printf("Berechneter Auftrieb: %d Einheiten\n", auftrieb)

    // Verbrauch
    gewicht := 80
    distanz := 5
    verbrauch := berechneEnergieVerbrauch(gewicht, distanz)
    fmt.Printf("Voraussichtlicher Verbrauch: %d Einheiten\n", verbrauch)

    // Check
    reserve, status := checkFlugStatus(auftrieb, verbrauch)

    fmt.Println("-------------------------------------------")
    fmt.Printf("Analyse-Ergebnis: %s (Reserve: %d)\n", status, reserve)

    if status == "Stabil" && reserve > 0 {
        fmt.Println("Mission abgeschlossen: Der Falke hält seine Bahn!")
        kreischen()
        kreischen()
        kreischen()
    } else {
        fmt.Println("Fehler: Die mathematische Kapselung ist noch instabil.")
    }
}
