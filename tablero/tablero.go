package tablero
import "sudoku/globals"


func InsertarValor(tablero *globals.Board, fila int , columna int , nuevoValor int) bool {
	if tablero.Celdas[fila][columna].Lockeado {
		return false
	}
	// tablero.Celdas[fila][columna].Lockeado = true
	tablero.Celdas[fila][columna].Valor = nuevoValor
	return true
}
func EliminarValor(tablero *globals.Board, fila int , columna int) bool {
	if tablero.Celdas[fila][columna].Lockeado {
		return false
	}
	tablero.Celdas[fila][columna].Valor = 0
	return true
}


func EsValido(tablero *globals.Board, fila, columna, valor int) bool {
	return ValidarColumna(tablero,fila,columna,valor) && ValidarFila(tablero,fila,columna,valor) && Validar3x3(tablero,fila,columna,valor)
}

func ValidarFila(tablero *globals.Board, fila, columna, valor int) bool{
	for col := 0; col < 9; col++ {
        if col != columna && tablero.Celdas[fila][col].Valor == valor {
            return false
        }
    }
	return true
}

func ValidarColumna(tablero *globals.Board, fila, columna, valor int) bool{
    for row := 0; row < 9; row++ {
        if row != fila && tablero.Celdas[row][columna].Valor == valor {
            return false
        }
    }
	return true
}

func Validar3x3(tablero *globals.Board, fila, columna, valor int) bool{
	inicioFila := (fila / 3) * 3
    inicioColumna := (columna / 3) * 3
    
    for i := inicioFila; i < inicioFila + 3; i++ {
        for j := inicioColumna; j < inicioColumna + 3; j++ {
            if (i != fila || j != columna) && tablero.Celdas[i][j].Valor == valor {
                return false
            }
        }
    }
	return true
}

func EncontrarVacia(tablero *globals.Board) (int, int, bool) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if tablero.Celdas[i][j].Valor == 0 {
                return i, j, true
            }
        }
    }
    return -1, -1, false
}

func VerificarMovimiento(tablero *globals.Board, solucion *globals.Board, fila, columna int) bool {
    return tablero.Celdas[fila][columna].Valor == solucion.Celdas[fila][columna].Valor
}

func TableroCompleto(tablero *globals.Board, solucion *globals.Board) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if tablero.Celdas[i][j].Valor != solucion.Celdas[i][j].Valor {
                return false
            }
        }
    }
    return true
}