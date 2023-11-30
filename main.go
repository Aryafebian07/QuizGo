package main

import "fmt"

type Mobil struct {
	Roda       [4]string
	PintuKiri  Pintu
	PintuKanan Pintu
}

type Pintu struct {
	SuaraKetuk string
	SuaraBuka  string
}

func (p *Pintu) Ketuk() {
	fmt.Println(p.SuaraKetuk)
}

func (p *Pintu) Buka() {
	fmt.Println(p.SuaraBuka)
}

func main() {
	mobil := Mobil{
		Roda: [4]string{"karet", "karet", "karet", "karet"},
		PintuKiri: Pintu{
			SuaraKetuk: "Beep",
			SuaraBuka:  "Knock",
		},
		PintuKanan: Pintu{
			SuaraKetuk: "Knock",
			SuaraBuka:  "Beep",
		},
	}

	fmt.Println("Keadaan Awal:")
	fmt.Println("Suara Ketuk Pintu Kiri:", mobil.PintuKiri.SuaraKetuk)
	fmt.Println("Suara Ketuk Pintu Kanan:", mobil.PintuKanan.SuaraKetuk)
	fmt.Println("Roda:", mobil.Roda)

	mobil.Roda = [4]string{"kayu", "kayu", "besi", "karet"}

	fmt.Println("\nSetelah Penggantian Roda:")
	fmt.Println("Roda:", mobil.Roda)

	fmt.Println("\nMengetuk dan Membuka Pintu:")
	mobil.PintuKiri.Ketuk()
	mobil.PintuKiri.Buka()

	mobil.PintuKanan.Ketuk()
	mobil.PintuKanan.Buka()
}
