package object

import "fmt"

type Kendaraan interface {
	Akselerasi()
}

type Sepeda struct {
	Suara  string
	Rantai string
}

func NewSepeda() (s Kendaraan) {
	s = &Sepeda{
		Suara:  "Swoosh",
		Rantai: "Normal",
	}
	return
}

func (s *Sepeda) Akselerasi() {
	fmt.Println(s.Suara)
	s.Rantai = "Perlu perbaikan"
}

type Mobil struct {
	Suara  string
	Bensin string
}

func NewMobil() (m Kendaraan) {
	m = &Mobil{
		Suara:  "Vroom",
		Bensin: "Penuh",
	}
	return
}

func (m *Mobil) Akselerasi() {
	fmt.Println(m.Suara)
	m.Bensin = "Kosong"
}
