package sapa

func Sapa() string {
	return `Hello`
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
