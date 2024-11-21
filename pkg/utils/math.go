package utils

func IsValidOperator(operator string) bool {
    switch operator {
    case "+", "-", "*", "/":
        return true
    default:
        return false
    }
}