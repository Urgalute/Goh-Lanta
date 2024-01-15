package main

import "fmt"

func stringToBytes(input string) []byte {
    // Utilisez []byte() pour convertir la chaîne en tableau d'octets
    result := []byte(input)
    return result
}

func bytesToString(input []byte) string {
    // Utilisez string() pour convertir le tableau d'octets en chaîne de caractères
    result := string(input)
    return result
}
func main() {
    // Exemple d'utilisation de la fonction
    inputString := "Hello, World!"
    byteSlice := stringToBytes(inputString)

    fmt.Printf("Input String: %s\n", inputString)
    fmt.Printf("Byte Slice: %v\n", byteSlice)
}

