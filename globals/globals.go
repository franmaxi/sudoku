package globals

type Board struct{
	Celdas [9][9] Celda 
}
type Celda struct{
	Lockeado 	bool
	Valor 		int
}