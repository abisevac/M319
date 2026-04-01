package main

import "fmt"

/* STORY KONTEXT:
   Um den reissenden Fluss ins Function Territory zu überqueren, müssen die Helden
   eine Brücke konstruieren. Sie nutzen IF ELSE für die Materialprüfung,
   SWITCH für die Pfeilerarten und LOOPS für die tausenden Haltekabel.
*/

// --- TEIL 1: Architektur Veranschaulichung ---

type Baumaterial struct {
    Name      string
    Qualitaet int // Wert von 0 bis 100
}

type BrueckenPfeiler struct {
    Baumaterial
    PfeilerTyp string
    IstStabil  bool
}

func main() {
    fmt.Println("=== BAU PROTOKOLL: BRÜCKE DER LOGIK ===")

    // --- TEIL 2: Initialisierung ---
    pfeilerKern := BrueckenPfeiler{
        Baumaterial: Baumaterial{
            Name:      "Granit Block",
            Qualitaet: 85,
        },
        PfeilerTyp: "Stein",
        IstStabil:  false,
    }

    var materialQualitaet int = pfeilerKern.Qualitaet
    var pfeilerTyp string = pfeilerKern.PfeilerTyp
    var kabelAnzahl int = 5

    // --- TEIL 3: IMPLEMENTIERUNG ---

    // AUFGABE A: Materialprüfung
    if materialQualitaet >= 80 {
        fmt.Println("[CHECK]: Materialprüfung bestanden. Fundament wird gesetzt.")
        pfeilerKern.IstStabil = true

        // AUFGABE B: Pfeiler Konstruktion
        switch pfeilerTyp {
        case "Stein":
            fmt.Println("Massiver Steinpfeiler verankert.")
        case "Holz":
            fmt.Println("Warnung: Holzpfeiler könnten bei Flut brechen!")
        default:
            fmt.Println("Unbekanntes Material! Stabilität nicht garantiert.")
        }

        // AUFGABE C: Haltekabel spannen
        fmt.Println("\n[AKTION]: Beginne mit dem Spannen der Kabel...")
        for i := 1; i <= kabelAnzahl; i++ {
            fmt.Printf("Haltekabel Nr. [%d] wird festgezogen...\n", i)
        }

    } else {
        fmt.Println("Baugenehmigung verweigert: Material zu schwach!")
    }

    // --- TEIL 4: Abschlussbericht ---
    fmt.Printf("\n--- Konstruktion der Einheit %s abgeschlossen ---\n", pfeilerKern.Name)
}
