package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	var ArrayProceso []Proceso
	var n int
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresa numero de procesos: ")
	fmt.Scan(&n)

	for i:=0; i<n; i++{
		var id uint64
		var Prioridad int64
		var Tiempo uint64
		var Estatus string
		
		fmt.Printf("#%d: ", i+1)
		fmt.Print("Ingresa un id: ")
		fmt.Scan(&id)
		fmt.Print("Ingresa la prioridad: ")
		fmt.Scan(&Prioridad)
		fmt.Print("Ingresa el tiempo: ")
		fmt.Scan(&Tiempo)
		fmt.Println("Ingresa el estatus: ")
		for s.Scan(){
			Estatus = s.Text()
			if Estatus != ""{
				break
			}
		}
		P := Proceso{id,Prioridad,Tiempo,Estatus}
		ArrayProceso = append(ArrayProceso, P)
	}

	var opc int
	fmt.Print("1.-Ascendente \n2.-Descendente\nElige una opcion: ")
	fmt.Scan(&opc)
	switch opc {
		case 1:{
			fmt.Println(ArrayProceso)
			sort.Sort(ByPrioridadAscendente(ArrayProceso))
			fmt.Println(ArrayProceso)
		}
		case 2:{
			fmt.Println(ArrayProceso)
			sort.Sort(ByPrioridadDescendente(ArrayProceso))
			fmt.Println(ArrayProceso)
		}
	}
}

type Proceso struct{
	id uint64
	Prioridad int64
	Tiempo uint64
	Estatus string
}

type ByPrioridadAscendente []Proceso

func (a ByPrioridadAscendente) Len() int           { return len(a) }
func (a ByPrioridadAscendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrioridadAscendente) Less(i, j int) bool { return a[i].Prioridad < a[j].Prioridad }

type ByPrioridadDescendente []Proceso

func (a ByPrioridadDescendente) Len() int           { return len(a) }
func (a ByPrioridadDescendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrioridadDescendente) Less(i, j int) bool { return a[i].Prioridad > a[j].Prioridad }