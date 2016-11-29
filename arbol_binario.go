package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// arbol binario con valores enteros.
type Arbol struct {
	Izquierda *Arbol
	Valor     string
	Derecha   *Arbol
}

//Funcion que recorre y opera el arbol
func RecorrerInorden(t *Arbol, s string) string {
	if t == nil {
		return s
	} else {
		var a string = t.Valor
		s = RecorrerInorden(t.Izquierda, s) + a + RecorrerInorden(t.Derecha, s)
		return s
	}
}

//Función recorrer in orden original
func RecorrerInorden1(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInorden1(t.Izquierda)
	fmt.Print(t.Valor)
	fmt.Print(" - ")
	RecorrerInorden1(t.Derecha)
}
func Separar(s string) {
	//fmt.Println(len(s))
	//fmt.Println(s[:6])
	a := 0

	var cont int = 0
	var res string = ""
	for a < len(s) {

		switch s[a] {
		case 43:
			fmt.Println("hola")
			res = res + s[:(cont+1)]
			s = s[:(cont + 1)]
			fmt.Println(res)
		case 45:
			fmt.Println("bien")
			res = res + s[:(cont+1)]
			s = s[:(cont + 1)]
			fmt.Println(res)
		default:
			res = ""
		}
		cont++
		a++
	}

}
func Operar(t *Arbol) int {
	if t.Izquierda == nil && t.Derecha == nil {
		num, _ := strconv.Atoi(t.Valor)
		return num
	} else {
		switch t.Valor {
		case "+":
			return Operar(t.Izquierda) + Operar(t.Derecha)
		case "-":
			return Operar(t.Izquierda) - Operar(t.Derecha)
		case "*":
			return Operar(t.Izquierda) * Operar(t.Derecha)
		case "/":
			return Operar(t.Izquierda) / Operar(t.Derecha)
		default:
			return 9999999999
		}
	}

}

//LA siguiente función detecta errores de sintaxis
func detecError(t *Arbol, er1 string) string {
	if t.Izquierda == nil && t.Derecha == nil {
		_, error := strconv.Atoi(t.Valor)
		if error != nil {
			er1 = "no hay un valor valido en la hoja, cambie " + t.Valor + " por un valor valido"
			return er1
		} else {
			er1 = ""
		}
		return er1
	} else {
		switch t.Valor {
		case "+", "-", "*", "/":
			er1 = detecError(t.Izquierda, "") + "\n" + detecError(t.Derecha, "")
			return er1
		default:
			er1 = "No hay un operador valido en el nodo cambie el valor de " + t.Valor + "\n" + detecError(t.Izquierda, "") + "\n" + detecError(t.Derecha, "")
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

//La siguiente función crea un arbol a partir de una entrada en postfijo
func hacerArbol(ent string) *Arbol {
	cadena := strings.Split(ent, " ")
	arbol := &Arbol{nil, "", nil}

	contador := 0
	s := NewStack()
	for i := 0; i < len(cadena); i++ {
		_, error := strconv.Atoi(cadena[i])
		if i != 0 && error != nil && contador >= 2 && (cadena[i] == "+" || cadena[i] == "-" || cadena[i] == "*" || cadena[i] == "/") {
			td := s.Pop()
			contador = contador - 1
			ti := s.Pop()
			contador = contador - 1
			arbol = &Arbol{ti, cadena[i], td}
			s.Push(arbol)
			contador = contador + 1

		} else {
			arbol = &Arbol{nil, cadena[i], nil}
			contador = contador + 1
			s.Push(arbol)
		}

	}
	arbol = s.Pop()
	return arbol
	/*	for i:=0; i< len(cadena); i++{
		tree:=s.Pop()
		fmt.Println(tree.Valor)
	} */

}

//Entran una serie de operaciones asignadas a distintas variables
//y se cierra cuando se da enter después de un espacio en blanco
func variables() {
	reader := bufio.NewReader(os.Stdin)
	//txt, _ := reader.ReadString('\n')
	cadena := ""
	for {
		entrada, _ := reader.ReadString('\n')
		cadena = cadena + "" + entrada
		if entrada == ""+"\n" {
			break
		}

	}
	entrada := strings.Split(cadena, "\n")
	entrada = entrada[:(len(entrada) - 2)]

	a := make(map[int]string)
	b := make(map[string]int)
	for i := 0; i < len(entrada); i++ {
		ver, valor := expresionesRegulares(entrada[i])
		if !ver {
			fmt.Println("los valores: " + valor + " no son valido")
			break
		}
		validar := strings.Split(entrada[i], " ")
		separa := strings.Split(entrada[i], ":")
		if validar[len(validar)-1] == ":=" {
			if separa[1] == "=" {
				//Se quita un espacio que queda al
				separa[0] = separa[0][:(len(separa[0]) - 1)]
				//aqui se deben evaluar los token
				variable := separa[0][(len(separa[0]) - 1):]
				operacion := separa[0][:(len(separa[0]) - 2)]
				comprobar := strings.Split(operacion, " ")
				for j := 0; j < len(comprobar); j++ {
					_, error := strconv.Atoi(comprobar[j])
					if error != nil && comprobar[j] != "+" && comprobar[j] != "-" && comprobar[j] != "*" && comprobar[j] != "/" {
						if len(a) > 0 {
							ver := false
							for k := 0; k < len(a); k++ {
								if comprobar[j] == a[k] {
									comprobar[j] = strconv.Itoa(b[a[k]])
									ver = true
								}
							}
							if ver == false {
								fmt.Println("La variable: " + comprobar[j] + " no está almacenada")
								break
							}
						} else {
							fmt.Println("La variable: " + comprobar[j] + " no está almacenada")
							break
						}
					}
				}
				operacion = ""
				for j := 0; j < len(comprobar); j++ {
					operacion = operacion + " " + comprobar[j]
				}
				operacion = operacion[1:]
				a[len(a)] = variable
				b[variable] = Operar(hacerArbol(operacion))
			}
		} else {
			fmt.Println("Ningun asignador encontrado")
		}

	}
	/*for i := 0; i < len(a); i++ {
		fmt.Println(b[a[i]])
	}*/
}

func expresionesRegulares(exp string) (bool, string) {
	comp := true
	valor := ""
	elementos := strings.Split(exp, " ")
	for i := 0; i < len(elementos); i++ {
		matchOp, _ := regexp.MatchString("^(\\+){1}$|^(\\-){1}$|^(\\*){1}$|^(\\/){1}$", elementos[i])
		matchVar, _ := regexp.MatchString("^([A-Za-z_])([a-zA-Z0-9_]*)$", elementos[i])
		matchNum, _ := regexp.MatchString("^([0-9]*)$", elementos[i])
		matchAsig, _ := regexp.MatchString("^(:=){1}$", elementos[i])
		if matchOp {
			fmt.Println("Operador -> " + elementos[i])
		} else if matchVar {
			fmt.Println("Variable -> " + elementos[i])
		} else if matchNum {
			fmt.Println("Numero -> " + elementos[i])
		} else if matchAsig {
			fmt.Println("Asignador -> " + elementos[i])
		} else {
			comp = false
			valor = valor + ", " + elementos[i]
		}

	}
	fmt.Println("----------------------------------------------")
	return comp, valor
}

//Haciendo tokens para validar las expresiones de las variables
func armarToken(exp string) string {
	elementos := strings.Split(exp, " ")
	fmt.Println(len(elementos))
	for i := 0; i < len(elementos); i++ {
		_, error := strconv.Atoi(elementos[i])
		if error == nil {
			fmt.Println("Valor -> " + elementos[i])
		} else if elementos[i] == "+" || elementos[i] == "-" || elementos[i] == "*" || elementos[i] == "/" {
			fmt.Println("Operador -> " + elementos[i])
		} else if elementos[i] == ":=" {
			fmt.Println("Asignador -> " + elementos[i])
		} else {
			fmt.Println("Variable ->" + elementos[i])
		}
	}

	return exp
}
func main() {
	/*
	   var s string =""
	    t2:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "+", &Arbol{nil,"3",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{nil,"3",nil} } }
	    t3:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "+", &Arbol{nil,"3",nil} }, "+", &Arbol{&Arbol{nil,"4",nil}, "*",&Arbol{nil,"6",nil} } }
	    var t1 *Arbol
	   t3:=&Arbol{nil,"5",nil}
	   t4:=&Arbol{&Arbol{nil,"-",nil},"2",&Arbol{nil,"*",nil}}
	   fmt.Println(Operar(t4))
	    fmt.Println(detecError(t2,""))
	   ent:= "5 6 * 5 3 * /"
	   t4:=hacerArbol(ent)
	   fmt.Println("Arbol t4: "+RecorrerInorden(t4,""))
	   fmt.Println(Operar(t4))
	*/
	/*
		//Prueba de construcción de arbol a partir de una entrada en postfijo
		ent := "- * + 10 3 * +"
		t4 := hacerArbol(ent)
		//prueba de detección de errores
		fmt.Println(detecError(t4, ""))
		//prueba de la funcion que opera un arbol
		fmt.Println(Operar(t4))
		//prueba de la función que recibe y asgina valores a ciertas variables y luego opera con estas
	*/
	//ent := "- + + 10 3 * +"
	//t4 := hacerArbol(ent)
	//fmt.Println(detecError(t4, ""))
	variables()
}
