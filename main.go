package main

import (
	"fmt"
	"math"
	"time"
)

func Toplama() {
	fmt.Println("İlk sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("İkinci sayiyi giriniz : ")
	var num2 int
	fmt.Scanln(&num2)
	result := num1 + num2
	fmt.Println("İşleminizin sonucu : ", result)
}

func Cikarma() {
	fmt.Println("İlk sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("İkinci sayiyi giriniz : ")
	var num2 int
	fmt.Scanln(&num2)
	result := num1 - num2
	fmt.Println("İşleminizin sonucu : ", result)
}

func Carpma() {
	fmt.Println("İlk sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("İkinci sayiyi giriniz : ")
	var num2 int
	fmt.Scanln(&num2)
	result := num1 * num2
	fmt.Println("İşleminizin sonucu : ", result)
}

func Bolme() {
	fmt.Println("İlk sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("İkinci sayiyi giriniz : ")
	var num2 int
	fmt.Scanln(&num2)
	result := num1 / num2
	fmt.Println("İşleminizin sonucu : ", result)
}

func usAlma() {
	fmt.Println("İlk sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("İkinci sayiyi giriniz : ")
	var num2 int
	fmt.Scanln(&num2)
	result := math.Pow(float64(num1), float64(num2))
	fmt.Println("İşleminizin sonucu : ", result)
}

func Karekok() {
	fmt.Println("Karekökünü almak istediğiniz sayiyi giriniz : ")
	var num1 int
	fmt.Scanln(&num1)
	result := math.Pow(float64(num1), 0.5)
	fmt.Println("İşleminizin sonucu : ", result)
}

func main() {
	for {
		fmt.Println("Hesap Makinesi")
		time.Sleep(2 * time.Second)
		fmt.Println("Yapmak İstediğiniz İşlemi Seçiniz 1)Toplama 2)Çıkarma 3) Çarpma 4)Bölme 5)Üs Alma 6)Karekök Hesaplama")
		var secim string
		fmt.Scanln(&secim)
		if secim == "1" {
			Toplama()
			time.Sleep(2 * time.Second)
		} else if secim == "2" {
			Cikarma()
			time.Sleep(2 * time.Second)
		} else if secim == "3" {
			Carpma()
			time.Sleep(2 * time.Second)
		} else if secim == "4" {
			Bolme()
			time.Sleep(2 * time.Second)
		} else if secim == "5" {
			usAlma()
			time.Sleep(2 * time.Second)
		} else if secim == "6" {
			Karekok()
			time.Sleep(2 * time.Second)
		} else if secim == "Q" {
			break
		} else {
			fmt.Println("Geçerli bir değer giriniz.")
			time.Sleep(2 * time.Second)
		}
	}
}
