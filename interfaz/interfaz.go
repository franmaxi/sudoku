package interfaz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sudoku/globals"
	"sudoku/tablero"
)

func limpiarPantalla() {
	fmt.Print("\033[H\033[2J")
}

func copiarTableroLocal(origen *globals.Board) *globals.Board {
	destino := &globals.Board{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			destino.Celdas[i][j] = origen.Celdas[i][j]
		}
	}
	return destino
}

// Recibe 'tab(actual), 'sol' (solución) e 'inicial' (pistas originales)
func MostrarTablero(tab *globals.Board, sol *globals.Board, inicial *globals.Board) {
	fmt.Println("\n    0 1 2   3 4 5   6 7 8")
	fmt.Println("  ┌───────┬───────┬───────┐")

	for i := 0; i < 9; i++ {
		if i == 3 || i == 6 {
			fmt.Println("  ├───────┼───────┼───────┤")
		}

		fmt.Printf("%d │ ", i)

		for j := 0; j < 9; j++ {
			valor := tab.Celdas[i][j].Valor
			valorCorrecto := sol.Celdas[i][j].Valor
			esPistaInicial := inicial.Celdas[i][j].Lockeado 

			if valor == 0 {
				fmt.Print(". ")
			} else if esPistaInicial {
				// 1. Si era una pista inicial: CYAN
				fmt.Printf("\033[1;36m%d\033[0m ", valor)
			} else if valor == valorCorrecto {
				// 2. Si el usuario acertó (coincide con solución): VERDE
				fmt.Printf("\033[1;32m%d\033[0m ", valor)
			} else {
				// 3. Si el usuario erró: ROJO
				fmt.Printf("\033[1;31m%d\033[0m ", valor)
			}

			if (j+1)%3 == 0 && j < 8 {
				fmt.Print("│ ")
			}
		}
		fmt.Println("│")
	}

	fmt.Println("  └───────┴───────┴───────┘")
}

func JugarTerminal(tab *globals.Board, sol *globals.Board) {
	reader := bufio.NewReader(os.Stdin)

	// IMPORTANTE: Guardamos una "foto" del tablero inicial para saber qué celdas eran pistas.
	tableroInicial := copiarTableroLocal(tab)

	for {
		limpiarPantalla()
		// Pasamos los 3 estados a la vista
		MostrarTablero(tab, sol, tableroInicial)

		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║ 1. Insertar número             ║")
		fmt.Println("║ 2. Borrar número               ║")
		fmt.Println("║ 3. Ver solución                ║")
		fmt.Println("║ 4. Salir                       ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("\nOpción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			fmt.Print("Fila (0-8): ")
			filaStr, _ := reader.ReadString('\n')
			fila, _ := strconv.Atoi(strings.TrimSpace(filaStr))

			fmt.Print("Columna (0-8): ")
			colStr, _ := reader.ReadString('\n')
			col, _ := strconv.Atoi(strings.TrimSpace(colStr))

			fmt.Print("Valor (1-9): ")
			valStr, _ := reader.ReadString('\n')
			val, _ := strconv.Atoi(strings.TrimSpace(valStr))

			if fila < 0 || fila > 8 || col < 0 || col > 8 || val < 1 || val > 9 {
				fmt.Println("\nValores inválidos")
				reader.ReadString('\n')
				continue
			}

			// Intentamos insertar. Si la celda está bloqueada (Cyan o Verde ya colocada), devolverá false.
			if !tablero.InsertarValor(tab, fila, col, val) {
				fmt.Println("\n⛔ No se puede modificar esta celda (Bloqueada)")
				fmt.Print("Presiona Enter...")
				reader.ReadString('\n')
				continue
			}

			// LÓGICA DE BLOQUEO DINÁMICO
			if tablero.VerificarMovimiento(tab, sol, fila, col) {
				tab.Celdas[fila][col].Lockeado = true
				
				// Verificamos victoria
				if tablero.TableroCompleto(tab, sol) {
					limpiarPantalla()
					MostrarTablero(tab, sol, tableroInicial)
					fmt.Println("\n ¡FELICIDADES! ¡GANASTE!")
					return
				}
			} 
			// Si es incorrecto, NO hacemos nada extra. 
			// No lo bloqueamos. MostrarTablero lo pintará de ROJO.

		case "2":
			fmt.Print("Fila (0-8): ")
			filaStr, _ := reader.ReadString('\n')
			fila, _ := strconv.Atoi(strings.TrimSpace(filaStr))

			fmt.Print("Columna (0-8): ")
			colStr, _ := reader.ReadString('\n')
			col, _ := strconv.Atoi(strings.TrimSpace(colStr))

			// Intentamos eliminar. Si está lockeado (Cyan inicial o Verde acertado), no dejará.
			if tablero.EliminarValor(tab, fila, col) {
				fmt.Println("\n✅ Eliminado")
			} else {
				fmt.Println("\n❌ No se puede eliminar (Celda bloqueada)")
			}
			fmt.Print("Presiona Enter...")
			reader.ReadString('\n')

		case "3":
			limpiarPantalla()
			fmt.Println("\n SOLUCIÓN:")
			// Para mostrar la solución, usamos 'sol' como tablero actual y como solución
			// Y pasamos 'sol' como inicial también para que se vea todo uniforme o cyan/verde según prefieras
			MostrarTablero(sol, sol, tableroInicial)
			fmt.Print("\n Presiona Enter para continuar...")
			reader.ReadString('\n')

		case "4":
			fmt.Println("\n ¡Gracias por jugar!")
			return

		default:
			fmt.Println("\n Opción inválida")
			reader.ReadString('\n')
		}
	}
}