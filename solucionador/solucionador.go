package solucionador

import (
    "math/rand"
    "time"
    "sudoku/globals"
    "sudoku/tablero"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func NumeroAleatorio(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func InicializarTablero(tab *globals.Board) {
    cantidadNumeros := NumeroAleatorio(15,22)
    colocados := 0
    intentos := 0
    maxIntentos := 10000
    
    for colocados < cantidadNumeros  && intentos < maxIntentos{
        fila := NumeroAleatorio(0, 8)
        columna := NumeroAleatorio(0, 8)
        valor := NumeroAleatorio(1, 9)
        
        if tab.Celdas[fila][columna].Valor == 0 && tablero.EsValido(tab, fila, columna, valor) {
            tablero.InsertarValor(tab, fila, columna, valor)
            tab.Celdas[fila][columna].Lockeado=true
            colocados++
        }
        intentos++
    }
}

func CopiarTablero(tableroOriginal *globals.Board) *globals.Board {
    tableroCopia := &globals.Board{}
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            tableroCopia.Celdas[i][j] = tableroOriginal.Celdas[i][j]
        }
    }
    return tableroCopia
}

func EncontrarSolucion(tab *globals.Board) *globals.Board {
    solucion := CopiarTablero(tab)
    
    if Resolver(solucion) {
        return solucion
    }
    return nil
}

func Resolver(tab *globals.Board) bool {
    fila, columna, encontrada := tablero.EncontrarVacia(tab)
    
    if !encontrada {
        return true
    }

    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    rand.Shuffle(len(numeros), func(i, j int) {
        numeros[i], numeros[j] = numeros[j], numeros[i]
    })
    
    for _, num := range numeros {
        if tablero.EsValido(tab, fila, columna, num) {
            tab.Celdas[fila][columna].Valor = num
            
            if Resolver(tab) {
                return true
            }
            
            tab.Celdas[fila][columna].Valor = 0
        }
    }
    
    return false
}