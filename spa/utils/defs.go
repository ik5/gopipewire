package utils

type SPADirection int

const (
	SPADirectionInput SPADirection = iota
	SPADirectionOutput
)

func DirectionReverse(d SPADirection) SPADirection {
	return d ^ 1
}

func (d SPADirection) DirectionnReverse() SPADirection {
	return d ^ 1
}
