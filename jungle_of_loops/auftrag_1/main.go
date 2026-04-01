package main

import "fmt"

func main() {

    /*
     * Fragment 1: Die unendliche Bewässerung
     * Die Ranken werden mit magischem Wasser versorgt,
     * bis das definierte Limit erreicht ist.
     */
    wasser := 0
    limit := 10

    fmt.Println("--- Start: Bewässerung ---")
    for wasser < limit {
        fmt.Println("Giesse magisches Wasser...")
        fmt.Println("Limit:", limit)
        wasser++
    }
    fmt.Println("Bewässerung abgeschlossen!\n")

    /*
     * Fragment 2: Die fehlerhafte Licht Zufuhr
     * Die Lichtenergie wird korrekt reduziert,
     * bis sie vollständig verbraucht ist.
     */
    lichtEnergie := 100

    fmt.Println("--- Start: Licht Zufuhr ---")
    for lichtEnergie > 0 {
        fmt.Println("Lichtstrahl wird fokussiert...")
        lichtEnergie--
        fmt.Println("Lichtenergie:", lichtEnergie)
    }
    fmt.Println("Lichtenergie verbraucht!\n")

    /*
     * Fragment 3: Der blockierte Wächter Check
     * Nach mehreren Versuchen werden die Ranken gezähmt.
     */
    istGezähmt := false
    versuche := 0

    fmt.Println("--- Start: Wächter Check ---")
    for !istGezähmt {
        fmt.Printf("Versuch %d: Besänftige die Ranken...\n", versuche)
        versuche++

        if versuche >= 10 {
            istGezähmt = true
        }
    }

    fmt.Println("Mission abgeschlossen: Die Ranken sind gezähmt!")
}
