package main

import "fmt"

// Danna Simaluisa

func main() {
	fmt.Println("Parcial 1 Go")

	var op int
	foriu := true

	for foriu {
		fmt.Println("Menú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Registrar una nueva venta\nOpción 2: Mostrar estadísticas\nOpción 3: Salir ")
		fmt.Scanln(&op)
		switch {
		case op == 1:
			fmt.Println("Productos disponibles\n Número    Producto    Precio en USD\n1    Arroz    1.25\n2    Leche    0.95\n3    Pan    0.50")
		var op int
		foriu := true
			for foriu {
				fmt.Println("Selecciona una opción: ")
				fmt.Scanln(&op)
				switch {
				case op == 1:
					nombre := Arroz
					precio := 1.25
					foriu = false
				case op == 2:
					nombre := Leche
					precio := 0.95
					foriu = false
				case op == 3:
					nombre := Pan
					precio := 0.50
					foriu = false
				}
			}
			RegistrarVenta(nombre, precio, cantidad)
		case op == 2:

		case op == 3:
			foriu = false
		}
	}

	prodven := []string []
	subtot := []float64 []

}

func RegistrarVenta(nombre string, precio float64, cantidad int) (subt float64, nom string) {
	

}
