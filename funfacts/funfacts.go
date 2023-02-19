package main

import "fmt"

func main() {
    var word string
    fmt.Print("Vil ha funfact om Sun, Luna eller Terra:")
    fmt.Scanln(&word)

    var response string
    switch word {
    case "sun":
        response = "Du har skrvet Sun!"
    case "luna":
        response = "Du har skrvet Luna!"
    case "terra":
        response = "Du har skrvet Terra!"
    default:
        response = "Du har ikke valgt et av alternativene."
    }

    fmt.Println(response)
}
