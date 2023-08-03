package utils

type SPADirection uint32

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

type Rectangle struct {
	Width  uint32
	Height uint32
}

// InitRectangle implements the SPA_RECTANGLE macro.
func InitRectangle(width, height uint32) Rectangle {
	return Rectangle{width, height}
}

type Point struct {
	X int32
	Y int32
}

// InitPoint implement the SPA_POINT macro.
func InitPoint(x, y int32) Point {
	return Point{x, y}
}

type Region struct {
	Position Point
	Size     Rectangle
}

// InitRegion implement the SPA_REGION macro
func InitRegion(x, y int32, width, height uint32) Region {
	return Region{InitPoint(x, y), InitRectangle(width, height)}
}

type Fraction struct {
	Num   uint32
	Denom uint32
}

// InitFraction implement the SPA_FRACTION macro.
func InitFraction(num, denom uint32) Fraction {
	return Fraction{num, denom}
}

const (
	TimeInvalid int64  = -1 << 63
	IDXInvalid         = -1 // #define SPA_IDX_INVALID  ((unsigned int)-1)
	IDInvalid   uint32 = 0xffffffff
)

const (
	NSecPerSec  int64 = 1000000000
	NSecPerMsec int64 = 1000000
	NSecPerUsec int64 = 1000
	USecPerSec  int64 = 1000000
	USecPerMsec int64 = 1000
	MSecPerSec  int64 = 1000
)

func TimespecToNSec(tvSec, tvNSec int64) int64 {
	return tvSec*NSecPerSec + tvNSec
}

func TimeSpecToUSec(tvSec, tvNSec int64) int64 {
	return tvSec*USecPerSec + tvNSec/NSecPerUsec
}

func TimeValToNSec(tvSec, tvUsec int64) int64 {
	return tvSec*NSecPerSec + tvUsec*NSecPerUsec
}

func TimeValToUSec(tvSec, tvUsec int64) int64 {
	return tvSec*USecPerSec + tvUsec
}

// Nop does exactly nothing
func Nop() {
	for {
	}
}
