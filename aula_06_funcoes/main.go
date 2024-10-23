package main
import "fmt"



func main() {
 fmt.Println("Funções em golang")

sub:=subtrair(20,10)
fmt.Println(sub)	

 fmt.Println(somar(2,2))

 printName("styven")

 fmt.Println(printNameAndAge("styven", 20))


}

func somar(a int, b int) int {
 return a + b
}

func subtrair(a int, b int) int {
 return a - b
}

func printName(name string) {
 fmt.Println("Hello", name)
}

func printNameAndAge(name string, age int)(string, int) {
	return name, age
}