package object

import "fmt"

type Kendaraan interface {
	Akselerasi()
}

type Sepeda struct {
	suara  string
	Rantai string
}

func NewSepeda() (s Kendaraan) {
	s = &Sepeda{
		suara:  "Swoosh",
		Rantai: "Normal",
	}
	return
}

func (s *Sepeda) Akselerasi() {
	fmt.Println(s.suara)
	s.Rantai = "Perlu perbaikan"
}

type Mobil struct {
	suara  string
	Bensin string
}

func NewMobil() (m Kendaraan) {
	m = &Mobil{
		suara:  "Vroom",
		Bensin: "Penuh",
	}
	return
}

func (m *Mobil) Akselerasi() {
	fmt.Println(m.suara)
	m.Bensin = "Kosong"
}
