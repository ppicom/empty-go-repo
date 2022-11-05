package app

type Cell struct {
	isAlive bool
}

func NewCell() Cell {
	return Cell{}
}

func (r *Cell) IsAlive() bool {

	return true

}
