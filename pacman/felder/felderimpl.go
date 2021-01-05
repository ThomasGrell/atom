package felder

import "gfx"

type data struct {
	posx, posy uint16
	breite uint16
	art rune          // 'W' Wand 'P' Pacman 'G' Ghost 'F' frei
}

func New (x,y, breite uint16, art rune) *data {
	var f *data = new(data)
	if art=='W' || art=='P' || art=='G' || art=='F' {
		(*f).posx=x
		(*f).posy=y
		(*f).breite=breite
		(*f).art=art
	} else {
		panic("Ungültige Feldart übergeben")
	}
	return f
}

func (f *data) GibFeldArt () rune {
	return (*f).art
}

func (f *data) SetzeFeldArt (art rune) {
	if art=='W' || art=='P' || art=='G' || art=='F' {
		(*f).art=art
	} else {
		panic("Ungültige Feldart übergeben")
	}
}

func (f *data) ZeichneFeld () {
	if gfx.FensterOffen() {
		switch (*f).art {
			case 'W':
				gfx.Stiftfarbe (0,0,155)
			case 'F':
				gfx.Stiftfarbe (0,0,0)
			case 'G':
				gfx.Stiftfarbe(200,200,200)
			case 'P':
				gfx.Stiftfarbe(200,200,0)
		}
		gfx.Vollrechteck ((*f).posx,(*f).posy,(*f).breite,(*f).breite)
	} else {
		panic("Grafikfenster ist nicht offen")
	}
}

func (f *data) GibPosition () (uint16,uint16) {
	return (*f).posx,(*f).posy
}

func (f *data) SetzePosition (posx,posy uint16) {
	(*f).posx=posx
	(*f).posy=posy
}
