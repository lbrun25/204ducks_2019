package main

import (
    "ducks"
    "os"
    "fmt"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./204ducks a")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\ta\t\tconstant")
}

func main() {
    if ducks.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !ducks.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    ducks.Ducks();
}