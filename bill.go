package main

type bill struct { //definimos el nombre de la estructura y sus parametros
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill { //constructor de la estructura
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b bill) format() string {

	return b.name
}
