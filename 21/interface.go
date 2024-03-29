package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}
type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung
	bangunDatar = &kubus{4}
	fmt.Println("===== persegi")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())
	fmt.Println("volume :", bangunDatar.volume())

}
