package main

import "fmt"

type Multimedia interface {
	mostrar()
}
type imagen struct {
	titulo  string
	formato string
	canales string
}

type audio struct {
	titulo   string
	formato  string
	duracion float32
}

type video struct {
	titulo  string
	formato string
	frames  int
}

type ContenidoWeb struct {
	multimedias []Multimedia
}

func (a *audio) mostrar() {
	fmt.Println(a.titulo)
	fmt.Println(a.formato)
	fmt.Println(a.duracion)
}
func (i *imagen) mostrar() {
	fmt.Println(i.titulo)
	fmt.Println(i.formato)
	fmt.Println(i.canales)
}
func (v *video) mostrar() {
	fmt.Println(v.titulo)
	fmt.Println(v.formato)
	fmt.Println(v.frames)
}

func main() {
	var opc int
	var ciclo bool = true
	contenidos := ContenidoWeb{}
	for ciclo {
		fmt.Println("Capturar...")
		fmt.Println("1. Imagen")
		fmt.Println("2. Audio")
		fmt.Println("3. Video")
		fmt.Println("4. Mostar")
		fmt.Println("0. Salir")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			i01 := imagen{}
			fmt.Println("Titulo")
			fmt.Scanln(&i01.titulo)
			fmt.Println("Formato")
			fmt.Scanln(&i01.formato)
			fmt.Println("Canales")
			fmt.Scanln(&i01.canales)
			contenidos.multimedias = append(contenidos.multimedias, &i01)
		case 2:
			a01 := audio{}
			fmt.Println("Titulo")
			fmt.Scanln(&a01.titulo)
			fmt.Println("Formato")
			fmt.Scanln(&a01.formato)
			fmt.Println("Duracion")
			fmt.Scanln(&a01.duracion)
			contenidos.multimedias = append(contenidos.multimedias, &a01)
		case 3:
			v01 := video{}
			fmt.Println("Titulo")
			fmt.Scanln(&v01.titulo)
			fmt.Println("Formato")
			fmt.Scanln(&v01.formato)
			fmt.Println("Frames")
			fmt.Scanln(&v01.frames)
			contenidos.multimedias = append(contenidos.multimedias, &v01)
		case 4:
			for _, c := range contenidos.multimedias {
				fmt.Println("---")
				c.mostrar()
			}
		case 0:
			ciclo = false
		default:
			fmt.Println("No valido")
		}
	}
}
