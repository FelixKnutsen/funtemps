package conv

import (
    "fmt"
)

func celsiusToFahrenheit(value float64) float64 {
    return ((celsius * (9/5)) + 32)
}

func fahrenheitToKelvin(value float64) float64 {
    return ((fahrenheit - 32) * 5/9 + 273.15)
}

func kelvinToCelsius(value float64) float64 {
    return (kelvin - 273.15)
}

func fahrenheitToCelsius(value float64) float64 {
    return ((fahrenheit - 32) * 5/9)
}

func kelvinToFahrenheit(value float64) float64 {
    return ((kelvin - 273.15) * 9/5 + 32) 
}

func celsiusToKelvin(value float64) float64 {
    return (celsius + 273.15)
}