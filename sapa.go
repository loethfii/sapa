package sapa

func Sapa(nama) string {
	return `Hello` + nama
}

type TanyaNama interface {
	ohkamu()
}

type Nama struct {
	nama string
}

func (n Nama) ohkamu() string {
	return n.nama
}
