package main

//import "fmt"
import "./felder"
import "gfx"
import "./zufallszahlen"
import "time"

// jswbqwjbcbjskn

var spielfeld [12][20](felder.Feld)
var fensterbreite uint16 = 650
var fensterhoehe uint16 = 410
var abstandVomRand uint16 = 25
var richtung uint8 // 0 nach rechts, 1 nach unten, 2 nach links, 3 nach oben
var posxPac, posyPac, posxGeist, posyGeist uint16
//var zufall int64

//Bin da (Rayk, Ben, Sebastian)


var level [12][20]rune = [12][20]rune{
	{'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
	{'W', 'P', 'F', 'W', 'F', 'F', 'F', 'F', 'F', 'W', 'F', 'F', 'F', 'W', 'F', 'F', 'F', 'W', 'F', 'W'},
	{'W', 'F', 'F', 'W', 'F', 'W', 'W', 'W', 'F', 'W', 'F', 'W', 'F', 'W', 'F', 'W', 'W', 'W', 'F', 'W'},
	{'W', 'F', 'F', 'W', 'F', 'F', 'F', 'F', 'F', 'W', 'F', 'W', 'F', 'W', 'F', 'F', 'F', 'F', 'F', 'W'},
	{'W', 'F', 'F', 'F', 'F', 'W', 'W', 'W', 'F', 'W', 'F', 'W', 'F', 'W', 'F', 'F', 'G', 'F', 'F', 'W'},
	{'W', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W', 'G', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W'},
	{'W', 'F', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'F', 'W', 'W', 'W', 'W', 'W', 'W', 'F', 'F', 'W'},
	{'W', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W', 'F', 'W', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W'},
	{'W', 'F', 'W', 'W', 'W', 'W', 'W', 'W', 'F', 'W', 'F', 'W', 'F', 'W', 'W', 'W', 'W', 'W', 'F', 'W'},
	{'W', 'F', 'F', 'F', 'F', 'F', 'F', 'W', 'F', 'W', 'F', 'W', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W'},
	{'W', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'W', 'F', 'F', 'W', 'F', 'F', 'W'},
	{'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'}}


func erstelleSpielfeld () {

	var feld felder.Feld
	var i, j uint16
	for i=0 ; i<12; i++ {
		for j=0; j<20; j++ {
			feld = felder.New(25+j*30,25+i*30,30,level[i][j])
			spielfeld[i][j]=feld
		}
	}
	//feld = felder.New(20,20,30,'W')
	//func New (x,y, breite uint16, art rune) *data { .. }

}

func zeichneFelder () {
	//
	// To Do: hier soll je nach Feldart das entsprechende Feld gezeichnet werden
	// Man beachte: Das Feld spielfeld ist hier bereits gefüllt.
	//
	for i := 0; i<12; i++ {
		for j := 0; j<20; j++ {
			(spielfeld[i][j]).ZeichneFeld()
		}
	}
}



func bewegePacMan () {
	//var pacman felder.Feld
	var gefunden bool

	for i := 0; i < 12; i++ {
		for j := 0; j < 20; j++ {
			if (spielfeld[i][j]).GibFeldArt() == 'P' {

				if (spielfeld[i][j+1]).GibFeldArt() == 'F' && richtung==0 {
					(spielfeld[i][j+1]).SetzeFeldArt('P')
					(spielfeld[i][j]).SetzeFeldArt('F')
				}
				if (spielfeld[i+1][j]).GibFeldArt() == 'F' && richtung==1 {
					(spielfeld[i+1][j]).SetzeFeldArt('P')
					(spielfeld[i][j]).SetzeFeldArt('F')
				}
				if (spielfeld[i][j-1]).GibFeldArt() == 'F' && richtung==2 {
					(spielfeld[i][j-1]).SetzeFeldArt('P')
					(spielfeld[i][j]).SetzeFeldArt('F')
				}
				if (spielfeld[i-1][j]).GibFeldArt() == 'F' && richtung==3 {
					(spielfeld[i-1][j]).SetzeFeldArt('P')
					(spielfeld[i][j]).SetzeFeldArt('F')
				}
				gefunden = true
				break
			}
		}
		if gefunden {
			break
		}
	}


				//Wenn Taste gedrückt wird, bewege Pacman in die Richtung
	// To Do: Je nach "richtung" wird das PacManFeld in spielfeld einFled nach oben,
	// nach unten, nach rechts oder nach links gesetzt
	//
}



func nachbarfelder(x, y int) []felder.Feld {
	var nachbarn []felder.Feld
	if x < 19 { // rechts
		nachbarn = append(nachbarn, spielfeld[y][x+1])
	}
	if y < 11 { // unten
		nachbarn = append(nachbarn, spielfeld[y+1][x])
	}
	if x > 0 { // links
		nachbarn = append(nachbarn, spielfeld[y][x-1])
	}
	if y > 0 { // oben
		nachbarn = append(nachbarn, spielfeld[y-1][x])
	}
	return nachbarn
}

func bewegeGeister() {
	var nachbarn, freieNachbarn []felder.Feld
	var naechstesFeld felder.Feld

	for i := 0; i < 12; i++ {
		for j := 0; j < 20; j++ {
			if spielfeld[i][j].GibFeldArt() == 'G' {
				nachbarn = nachbarfelder(j, i)
				freieNachbarn = nil
				for _, nachbar := range(nachbarn) {
					if nachbar.GibFeldArt() == 'F' {
						freieNachbarn = append(freieNachbarn, nachbar)
					}
				}

				naechstesFeld = freieNachbarn[zufallszahlen.Zufallszahl(0, int64(len(freieNachbarn)-1))]
				spielfeld[i][j].SetzeFeldArt('F')
				naechstesFeld.SetzeFeldArt('G')
			}
		}
	}
}

func leseRichtung() {
	for gfx.FensterOffen() {
		taste,status,_:=gfx.TastaturLesen1()
		if status==1 {
			switch taste {
				case 273:
					richtung=3
				case 274:
					richtung=1
				case 275:
					richtung=0
				case 276:
					richtung=2
			}
		}
	}
}


func main() {
	zufallszahlen.Randomisieren()

	gfx.Fenster(fensterbreite, fensterhoehe)

	//fmt.Println(zufallszahlen.Zufallszahl(1,4))

	erstelleSpielfeld()

	zeichneFelder()

	/*
	tast,_,_:=gfx.TastaturLesen1() // Oben 273
	fmt.Println(tast)
	tast,_,_=gfx.TastaturLesen1() // Unten 274
	fmt.Println(tast)
	tast,_,_=gfx.TastaturLesen1() // Rechts 275
	fmt.Println(tast)
	tast,_,_=gfx.TastaturLesen1() // Links 276
	fmt.Println(tast)
	*/
	//taste,status,tiefe := TastaturpufferLesen1 ()

	go leseRichtung()

	for gfx.FensterOffen() {
		//leseRichtung()
		bewegeGeister()
		bewegePacMan()
		zeichneFelder()
		time.Sleep(time.Second)

	}

	//gfx.TastaturLesen1()

}
