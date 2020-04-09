package utils

import (
        "strconv"
)

func StringToFloat64(str string) float64 {

        var result float64

        result, _ = strconv.ParseFloat(str, 64)

        return result
}

func BoolToFloat64(b bool) float64 {

        var result float64

        if b == true {
                result = 1
        } else {
                result = 0
        }

        return result
}

