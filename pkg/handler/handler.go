package handler

import (
    "fmt"
    "github.com/boysimon10/go-simple-calc/pkg/calculator"
    "github.com/boysimon10/go-simple-calc/pkg/utils"
)

func GetInput() (float64, string, float64, error) {
    var num1, num2 float64
    var operator string

    _, err := fmt.Scanf("%f %s %f", &num1, &operator, &num2)
    if err != nil {
        return 0, "", 0, err
    }

    if !utils.IsValidOperator(operator) {
        return 0, "", 0, fmt.Errorf("invalid operator")
    }

    return num1, operator, num2, nil
}

func Calculate(num1 float64, operator string, num2 float64) (float64, error) {
    switch operator {
    case "+":
        return calculator.Add(num1, num2), nil
    case "-":
        return calculator.Subtract(num1, num2), nil
    case "*":
        return calculator.Multiply(num1, num2), nil
    case "/":
        if num2 == 0 {
            return 0, fmt.Errorf("division by zero")
        }
        return calculator.Divide(num1, num2), nil
    default:
        return 0, fmt.Errorf("invalid operator")
    }
}
