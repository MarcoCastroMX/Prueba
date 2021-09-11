package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	var slice []string
	var n int
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresa numero de strings: ")
	fmt.Scan(&n)

	for i:=0; i<n; i++{
		var newString string
		
		fmt.Printf("#%d: ", i+1)
		for s.Scan(){
			newString = s.Text()
			if newString != ""{
				break
			}
		}
		slice = append(slice, newString)
	}

	var opc int
	fmt.Print("1.-Ascendente \n2.-Descendente\nElige una opcion: ")
	fmt.Scan(&opc)

	switch opc{
		case 1:{
			sort.StringSlice.Sort(slice)
			ascendente :=strings.Join(slice,"\n")
			file, err := os.Create("ascendente.txt")
			if err != nil{
				fmt.Println(err)
				return
			}
			defer file.Close()

			file.WriteString(ascendente)
		}
		case 2:{
			sort.StringSlice.Sort(slice)
			for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
				slice[i], slice[j] = slice[j], slice[i]
			}
			descendente :=strings.Join(slice,"\n")
			file, err := os.Create("descendente.txt")
			if err != nil{
				fmt.Println(err)
				return
			}
			defer file.Close()

			file.WriteString(descendente)
		}
	}

	
}