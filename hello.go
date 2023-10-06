package main

import "fmt"

func main() {
	flat1 := flat{
		area:       150,
		furniture:  []string{"Диван", "Комод", "Кровать", "Стул", "Стол"},
		appliances: []string{"Чайник", "Тостер", "Утюг", "Мультиварка", "Стиралка"},
		family:     []string{"Мама", "Папа", "Я", "Кошка"},
	}
	flat1.seeFlat()
}

type flat struct {
	area       float64
	furniture  []string
	appliances []string
	family     []string
}

type furnitureStruct struct {
	name string
}
type appliancesStruct struct {
	name string
}
type familyStruct struct {
	name string
}

func (f flat) seeFlat() {
	fmt.Println("Площадь квартиры:", f.area, "кв.м")
	fmt.Println("Из мебели есть: ")
	for index, furnitureStruct := range f.furniture {
		fmt.Println(index+1, furnitureStruct)
	}
	fmt.Print("Есть бытовая техника, такая как: \n")
	for index, appliancesStruct := range f.appliances {
		fmt.Println(index+1, appliancesStruct)
	}
	fmt.Println("В квартире живут: ")
	for index, familyStruct := range f.family {
		fmt.Println(index+1, familyStruct)
	}
}
