package main

import (
	"fmt"
	"strings"
	"strconv"
//	"math/rand"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}



func RecorrerInorden(t *Arbol,s string)string{
	if t == nil {
		return s
	}else{
		var a string =t.Valor
		s=RecorrerInorden(t.Izquierda,s)+a+RecorrerInorden(t.Derecha,s)
		return s
	}
}
func Separar(s string){
	//fmt.Println(len(s))
	//fmt.Println(s[:6])
	a:=0

	var cont int=0
	var res string =""
	for a<len(s){

			switch s[a] {
				case 43:
					fmt.Println("hola")
					res=res+s[:(cont+1)]
					s=s[:(cont+1)]
					fmt.Println(res)
				case 45:
					fmt.Println("bien")
					res=res+s[:(cont+1)]
					s=s[:(cont+1)]
					fmt.Println(res)
				default:
					res=""
			}
		cont++
		a++
	}

}
func Operar(t *Arbol)int{
		if(t.Izquierda==nil && t.Derecha==nil){
			num,_ := strconv.Atoi(t.Valor)
			return num
		}else{
			switch t.Valor {
			case "+":
				return Operar(t.Izquierda)+Operar(t.Derecha)
			case "-":
				return Operar(t.Izquierda)-Operar(t.Derecha)
			case "*":
				return Operar(t.Izquierda)*Operar(t.Derecha)
			case "/":
				return Operar(t.Izquierda)/Operar(t.Derecha)
			default:
				return 9999999999
			}
		}

}
func detecError(t *Arbol, er1 string)string{
	if(t.Izquierda==nil && t.Derecha==nil){
		_,error := strconv.Atoi(t.Valor)
		if(error!=nil){
			er1="no hay un valor valido en la hoja, cambie "+t.Valor+ " por un valor valido"
			return er1
		}else{
			er1=""
		}
		return er1
	}else{
		switch t.Valor {
		case "+","-","*","/":
			er1= detecError(t.Izquierda,"")+"\n"+detecError(t.Derecha,"")
			return er1
		default:
			er1 = "No hay un operador valido en el nodo cambie el valor de "+t.Valor+"\n"+detecError(t.Izquierda,"")+"\n"+detecError(t.Derecha,"")
			return er1
		}
	}
}

type Stack struct {
  nodes []*Arbol
  count int
}
// NewStack returns a new pila.
func NewStack() *Stack {
  return &Stack{}
}
// Push adds a node to the pila.
func (s *Stack) Push(n *Arbol) {
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}
func (s *Stack) Pop() *Arbol {
  if s.count == 0 {
    return nil
  }
  s.count--
  return s.nodes[s.count]
}
func hacerArbol(ent string)*Arbol{
	cadena:=strings.Split(ent, " ")
	arbol:=&Arbol{nil,"",nil}
	contador:= 0
	s:= NewStack()
	for i:=0; i< len(cadena); i++{
			_,error := strconv.Atoi(cadena[i])
		if(i!=0 && error!=nil && contador>=2 && (cadena[i]=="+" || cadena[i]=="-"|| cadena[i]=="*" || cadena[i]=="/")){
				td:=s.Pop()
				contador=contador-1
				ti:=s.Pop()
				contador=contador-1
				arbol:=&Arbol{ti,cadena[i],td}
				s.Push(arbol)
				contador=contador+1


		}else if(error==nil){
			arbol:=&Arbol{nil,cadena[i],nil}
			contador=contador+1
			s.Push(arbol)
		}

	}
	arbol=s.Pop()
	return arbol
/*	for i:=0; i< len(cadena); i++{
		tree:=s.Pop()
		fmt.Println(tree.Valor)
	} */

}

func main() {
//	var s string =""
	//t2:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "+", &Arbol{nil,"3",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{nil,"3",nil} } }
	//t3:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "+", &Arbol{nil,"3",nil} }, "+", &Arbol{&Arbol{nil,"4",nil}, "*",&Arbol{nil,"6",nil} } }
	//var t1 *Arbol
//	t3:=&Arbol{nil,"5",nil}
//t4:=&Arbol{&Arbol{nil,"-",nil},"2",&Arbol{nil,"*",nil}}
//	fmt.Println(Operar(t4))
	//fmt.Println(detecError(t2,""))
	ent:= "5 6 * 5 3 * /"
	t4:=hacerArbol(ent)
	fmt.Println("Arbol t4: "+RecorrerInorden(t4,""))
	fmt.Println(Operar(t4))
}
