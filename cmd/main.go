package main

import (
    "fmt"
    "os"
    "github.com/boysimon10/go-simple-calc/pkg/handler"
)

func main() {
    fmt.Println("Calculatrice Go !")
    fmt.Println("[nombre 1] [opérateur] [nombre 2]")
    fmt.Println("[opérateur]: +, -, *, /")

    num1, operator, num2, err := handler.GetInput()
    if err != nil {
        fmt.Println("Erreur dans l'entrée:", err)
        os.Exit(1)
    }

    result, err := handler.Calculate(num1, operator, num2)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Le résultat de %.2f %s %.2f est %.2f\n", num1, operator, num2, result)
}