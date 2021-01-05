//ja
package felder

type Feld interface {

	// Konstruktor New(...) erstellt ein struct *data auf das die folgenden
	// Methoden angewendet werden können:

// Vor.: -
// Erg.: Die Feldart ist geliefert: 'W' für Wand, 'P' für Pacman,
//  'G' für Ghost und 'F' für frei.
GibFeldArt () rune

// Vor.: Es muss entwerder 'W', 'G', 'P' oder 'F' übergeben werden
// Erg.: Das Feld ist nun von der Art "art"
SetzeFeldArt (art rune)

// Vor.: Das Grafikfenster ist geöffnet.
// Erg.: Das Feld ist je nach Feldart gezeichnet:
// Packman ist Gelb, Der Geist weiß, freie Felder schwarz und Wandfelder grau
ZeichneFeld ()

// Vor.: -
// Erg.: Die x- und yKoordinate der linken oberen Ecke des Feldes ist geliefert
GibPosition () (uint16,uint16)


// Vor.: -
// Erg.: Das Feld hat nun die Koordinaten (x|y), welche die Position
// der linken oberen Ecke des Feldes angeben
SetzePosition (x uint16,y uint16)

}
