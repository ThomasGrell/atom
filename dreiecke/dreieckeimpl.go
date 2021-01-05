package dreiecke

import ( "gfx" ; "math" ; "reflect" )

type data struct {
	a,b,c  Punkt
	winkel float64
	farbe  Farbe
}

//Hilfsfunktionen
func skalarprod (a,b Vektor) float64{
	return a[0]*b[0]+a[1]*b[1]+a[2]*b[2]
}

func vektorprod (a,b Vektor) (c Vektor) {
	c[0] = a[1]*b[2] - b[1]*a[2]
	c[1] = a[2]*b[0] - b[2]*a[0]
	c[2] = a[0]*b[1] - b[0]*a[1]
	return
}

func New (a,b,c Punkt, f Farbe) *data {
	var ab, ac Vektor
	var d *data
	d = new (data)
	(*d).a = a
	(*d).b = b
	(*d).c = c
	(*d).farbe = f
	for i:=0; i < 3; i++ {
		ab[i]= (*d).b[i]-(*d).a[i]
		ac[i]= (*d).c[i]-(*d).a[i]
	}
	n:= vektorprod(ab, ac) //ein Normalenvektor der Dreiecksfläche
	nz:= [3]float64 {0,0,1} //Einheits-Normalenvektor in z-Richtung
	(*d).winkel = math.Acos (skalarprod(n, nz)/math.Sqrt(skalarprod(n,n)))
	return d
}

//Hilfsfunktion
func dreheZ (alpha float64, a Punkt) (p Punkt) {
	var winkel float64 = alpha*math.Pi/180
	p[0] = math.Cos(winkel)*a[0]+ math.Sin(winkel)*a[1]
	p[1] = -math.Sin(winkel)*a[0] + math.Cos(winkel)*a[1]
	p[2] = a[2]
	return
}

func dreheX (alpha float64, a Punkt) (p Punkt) {
	var winkel float64 = alpha*math.Pi/180
	p[0] = a[0]
	p[1] = math.Cos(winkel)*a[1]- math.Sin(winkel)*a[2]
	p[2] = math.Sin(winkel)*a[1]+ math.Cos(winkel)*a[2]
	return
}

func dreheY (alpha float64, a Punkt) (p Punkt) {
	var winkel float64 = alpha*math.Pi/180
	p[0] = math.Cos(winkel)*a[0]+ math.Sin(winkel)*a[2]
	p[1] = a[1]
	p[2] = -math.Sin(winkel)*a[0]+ math.Cos(winkel)*a[2]
	//fmt.Println ("Drehe:",a,"auf",p)
	return
}

func (d *data) DrehenUmZ (alpha float64) interface {} {
	var a_,b_,c_ Punkt
	a_ = dreheZ (alpha, (*d).a)
	b_ = dreheZ (alpha, (*d).b)
	c_ = dreheZ (alpha, (*d).c)
	e:= New (a_,b_,c_,(*d).farbe)
	return e
}

func (d *data) DrehenUmX (alpha float64) interface {} {
	var a_,b_,c_ Punkt
	a_ = dreheX (alpha, (*d).a)
	b_ = dreheX (alpha, (*d).b)
	c_ = dreheX (alpha, (*d).c)
	e:= New (a_,b_,c_,(*d).farbe)
	return e
}

func (d *data) DrehenUmY (alpha float64) interface {} {
	var a_,b_,c_ Punkt
	a_ = dreheY (alpha, (*d).a)
	b_ = dreheY (alpha, (*d).b)
	c_ = dreheY (alpha, (*d).c)
	e:= New (a_,b_,c_,(*d).farbe)
	return e
}

func (d *data) Verschieben (dx,dy,dz float64) interface {} {
	var a_,b_,c_ Punkt
	a_[0],a_[1],a_[2] = (*d).a[0]+dx,(*d).a[1]+dy,(*d).a[2]+dz
	b_[0],b_[1],b_[2] = (*d).b[0]+dx,(*d).b[1]+dy,(*d).b[2]+dz
	c_[0],c_[1],c_[2] = (*d).c[0]+dx,(*d).c[1]+dy,(*d).c[2]+dz
	return New (a_,b_,c_,(*d).farbe)
}

func (d *data) Grafik () {
	var faktorA, faktorB, faktorC float64 = 100,100,100
	gfx.Stiftfarbe (uint8(math.Abs(math.Cos ((*d).winkel)) * float64((*d).farbe[0])),
	                uint8(math.Abs(math.Cos ((*d).winkel)) * float64((*d).farbe[1])),
	                uint8(math.Abs(math.Cos ((*d).winkel)) * float64((*d).farbe[2])))
	if abstand:= (*d).a[2]; abstand!=0 {
		faktorA= 320.0/abstand
	}
	if abstand:= (*d).b[2]; abstand!=0 {
		faktorB= 320.0/abstand
	}
	if abstand:= (*d).c[2]; abstand!=0 {
		faktorC= 320.0/abstand
	}
	// Randkorrekturen bzgl. des Fensters
	x0:=faktorA*d.a[0]+320
	if x0 < 0 { x0 = 0 }
	if x0 > 639 {x0 = 639 }
	y0:= 240-faktorA*d.a[1]
	if y0<0 {y0 = 0}
	if y0>479 {y0 = 479}
	x1:=faktorB*d.b[0]+320
	if x1 < 0 { x1 = 0 }
	if x1 > 639 {x1 = 639 }
	y1:= 240-faktorB*d.b[1]
	if y1<0 {y1 = 0}
	if y1>479 {y1 = 479}
	x2:=faktorC*d.c[0]+320
	if x2 < 0 { x2 = 0 }
	if x2 > 639 {x2 = 639 }
	y2:= 240-faktorC*d.c[1]
	if y2<0 {y2 = 0}
	if y2>479 {y2 = 479}
	if d.a[2]<0 && d.b[2]<0 && d.c[2]<0{
		gfx.Volldreieck (uint16(x0),uint16(y0),uint16(x1),uint16(y1),uint16(x2),uint16(y2))
	}
}

func (d *data) Kopie () interface{} {
	var dk *data
	dk = new (data)
	(*dk)  = (*d)
	return dk
}

//Hilfsfunktionen
func istkleiner (a,b [9]float64) bool {
	for i:=0 ; i<9; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

func einordnen (a,b,c Punkt) (f [9]float64) {
	var reihenfolge [3]uint
	if a[2] <= b[2] && b[2] <= c[2] {
		reihenfolge = [3]uint {0,1,2}
	} else if  a[2] <= c[2] && c[2] <= b[2] {
		reihenfolge = [3]uint {0,2,1}
	} else if b[2] <= a[2] && a[2] <= c[2] {
		reihenfolge = [3]uint {1,0,2}
	} else if b[2] <= c[2] && c[2] <= a[2] {
		reihenfolge = [3]uint {1,2,0}
	} else if c[2] <= a[2] && a[2] <= b[2] {
		reihenfolge = [3]uint {2,0,1}
	} else {
		reihenfolge = [3]uint {2,1,0}
	}
	for i:=0; i<3; i++ {
		for j:=0 ;j<3; j++ {
			switch reihenfolge[i] {
				case 0:
				f[3*i+j] = a[j]
				case 1:
				f[3*i+j] = b[j]
				case 2:
				f[3*i+j] = c[j]
			}
		}
	}
	return
}
			
func umordnen (a [9]float64) (f [9]float64){
	f = [9]float64 { a[2],a[5],a[8],a[0],a[3],a[6],a[1],a[4],a[7] }
	return
}

func (d *data) Kleiner (d2 interface{}) bool {
	var x,y [9]float64
	if reflect.TypeOf(d) != reflect.TypeOf(d2) {
		panic ("Fehler: Vergleich verschiedener Typen ist nicht möglich!")
	}
	e:=d2.(*data)
	//fmt.Println ("Einordnen der eigenen Punkte:",einordnen ((*d).a,(*d).b,(*d).c))
	x = umordnen (einordnen ((*d).a,(*d).b,(*d).c))
	//fmt.Println ("Umgeordnet:",x)
	
	//fmt.Println ("Einordnen der anderen Punkte:",einordnen ((*e).a,(*e).b,(*e).c))
	y = umordnen (einordnen ((*e).a,(*e).b,(*e).c))
	//fmt.Println ("Umgeordnet:",y)
	return istkleiner (x,y)
}

func init () {
	gfx.Fenster (640,480)
}
