package main
import "fmt"
func main() {
	fmt.Println("Olá! Qual o seu nome?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Olá %s. Esse foi meu segundo programa em Go!\n", name)
    }
