package main

import (
    "fmt"
    "sudoku/globals"
    "sudoku/solucionador"
    "sudoku/interfaz"
)

func main() {
    fmt.Println("Iniciando Sudoku...")
    
    miTablero := &globals.Board{}
    solucionador.InicializarTablero(miTablero)
    
    tableroSolucion := solucionador.EncontrarSolucion(miTablero)
    if tableroSolucion == nil {
        fmt.Println(" Error: No se pudo generar el Sudoku")
        return
    }
    
    fmt.Print("Presiona Enter para comenzar...")
    fmt.Scanln()
    
    interfaz.JugarTerminal(miTablero, tableroSolucion)
}