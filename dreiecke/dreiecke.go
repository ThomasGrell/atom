package dreiecke
// Autor: St. Schmidt
// Datum: 03.11.2020
// Zweck: ADT Dreieck (im Raum) --> perspektivische Darstellung

type Punkt [3]float64 //ein "Tripel" aus drei Koordinaten:x, y, z

type Farbe [3]uint8  //ein "Tripel" aus den drei Farbwerten: rot, grün, blau
                     //gemäß dem RGB-24-Bit-Farbmodell

type Vektor [3]float64 // ein "Tripel" aus drei Komponenten: x, y, z

// Ein Dreieck ist definiert durch seine drei Eckpunkte A, B und C,
// wobei die Eckpunkte jeweils drei Koordinaten haben und das Dreieck
// damit als Teil einer Ebene im Raum liegt. Die Strecken AB, BC und CA
// stellen den Rand des Dreiecks dar, der die Dreieckfläche einschließt,
// die eine festgelegte Farbe besitzt.

// Vor.: -
// Erg.: ein neues Dreieck mit den Eckpunkte a,b und c und der 
//       Flächenfarbe f 
// New(a,b,c Punkt, f Farbe) *data //*data erfüllt das Interface Dreieck

type Dreieck interface {

	// Vor.: keine
	// Eff.: Das Dreieck ist im gfx-Grafikfenster dargestellt, wenn all
	//       seine Ecken im Sichtbereich liegen. Dabei erfolgt eine
	//       perspektive "Verzerrung" in negativer z-Richtung.  (UNGENAU!!)
	Grafik ()
	
	// Vor.: -
	// Erg.: Ein Dreieck ist geliefert, das sich aus dem Dreieck ergibt, 
	//       ween man es um den Winkel alpha im mathematisch positiven
	//       Sinne um die x-Achse gedreht hat.
	DrehenUmX (alpha float64) interface {}
	
	// Vor.: -
	// Erg.: Ein Dreieck ist geliefert, das sich aus dem Dreieck ergibt, 
	//       ween man es um den Winkel alpha im mathematisch positiven
	//       Sinne um die y-Achse gedreht hat.
	DrehenUmY (alpha float64) interface {}
	
	// Vor.: -
	// Erg.: Ein Dreieck ist geliefert, das sich aus dem Dreieck ergibt, 
	//       ween man es um den Winkel alpha im mathematisch positiven
	//       Sinne um die z-Achse gedreht hat.
	DrehenUmZ (alpha float64) interface {}
	
	// Vor.: -
	// Erg.: Ein Dreieck ist geliefert, das sich aus dem Dreieck ergibt, 
	//       wenn man es um dx Einheite in X-Richtung, dy Einheiten
	//       in Y-Richtung und um dz Einheiten in Z-Richtung verschoben hat.
	Verschieben (dx,dy,dz float64) interface {}
	
	// Vor.: -
	// Erg.: Eine tiefe Kopie der Dreiecksinstanz ist geliefert. Der statische
	//       Typ ist interface {}, der dahinterliegende dyn. DT ist vom 
	//       Typ der aufrufenden Instanz.
	Kopie () interface{}
	
	// Vor.: x erfüllt das Interface Dreieck und ist vom gleichen dynamischen
	//       DT wie die aufrufende Instanz.
	// Erg.: Der Wert -true- ist geliefert, gdw. die aufrufende Instanz
	//       kleiner war als x. Das ist genau dann der Fall, wenn das Dreieck
	//       "weiter vorn bzgl. der z-Achse im Raum liegt".  ---------------------UNGENAU
	Kleiner (x interface{}) bool 
	
}

