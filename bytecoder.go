// bytecoder provides functions for converting standard arrays into byte arrays.
// Based on the Bytes function from golang.org\x\mobile\exp\f32\f32.go
package bytecoder

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"image/color"
)

// getByteOrder returns false for LittleEndian and true for BigEndian. It panics if the provided parameter isn't one of the two.
func getByteOrder(byteOrder binary.ByteOrder) bool {
	le := false
	switch byteOrder {
	case binary.BigEndian:
	case binary.LittleEndian:
		le = true
	default:
		panic(fmt.Sprintf("invalid byte order %v", byteOrder))
	}
	return le
}

// Float32 returns the byte representation of float32 values in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func Float32(byteOrder binary.ByteOrder, values ...float32) []byte {
	le := getByteOrder(byteOrder)
	width := 4
	arr := make([]byte, width*len(values))
	for i, v := range values {
		u := math.Float32bits(v)
		if le {
			arr[width*i+0] = byte(u >> 0)
			arr[width*i+1] = byte(u >> 8)
			arr[width*i+2] = byte(u >> 16)
			arr[width*i+3] = byte(u >> 24)
		} else {
			arr[width*i+0] = byte(u >> 24)
			arr[width*i+1] = byte(u >> 16)
			arr[width*i+2] = byte(u >> 8)
			arr[width*i+3] = byte(u >> 0)
		}
	}
	return arr
}

func Vec2(byteOrder binary.ByteOrder, values ...mgl32.Vec2) []byte {
	le := getByteOrder(byteOrder)
	width := 8
	arr := make([]byte, width*len(values))
	for i, v := range values {
		x := math.Float32bits(v[0])
		y := math.Float32bits(v[1])
		if le {
			arr[width*i+0] = byte(x >> 0)
			arr[width*i+1] = byte(x >> 8)
			arr[width*i+2] = byte(x >> 16)
			arr[width*i+3] = byte(x >> 24)
			arr[width*i+4] = byte(y >> 0)
			arr[width*i+5] = byte(y >> 8)
			arr[width*i+6] = byte(y >> 16)
			arr[width*i+7] = byte(y >> 24)
		} else {
			arr[width*i+0] = byte(x >> 24)
			arr[width*i+1] = byte(x >> 16)
			arr[width*i+2] = byte(x >> 8)
			arr[width*i+3] = byte(x >> 0)
			arr[width*i+4] = byte(y >> 24)
			arr[width*i+5] = byte(y >> 16)
			arr[width*i+6] = byte(y >> 8)
			arr[width*i+7] = byte(y >> 0)
		}
	}
	return arr
}

func Vec3(byteOrder binary.ByteOrder, values ...mgl32.Vec3) []byte {
	le := getByteOrder(byteOrder)
	width := 12
	arr := make([]byte, width*len(values))
	for i, v := range values {
		x := math.Float32bits(v[0])
		y := math.Float32bits(v[1])
		z := math.Float32bits(v[2])
		if le {
			arr[width*i+0] = byte(x >> 0)
			arr[width*i+1] = byte(x >> 8)
			arr[width*i+2] = byte(x >> 16)
			arr[width*i+3] = byte(x >> 24)
			arr[width*i+4] = byte(y >> 0)
			arr[width*i+5] = byte(y >> 8)
			arr[width*i+6] = byte(y >> 16)
			arr[width*i+7] = byte(y >> 24)
			arr[width*i+8] = byte(z >> 0)
			arr[width*i+9] = byte(z >> 8)
			arr[width*i+10] = byte(z >> 16)
			arr[width*i+11] = byte(z >> 24)
		} else {
			arr[width*i+0] = byte(x >> 24)
			arr[width*i+1] = byte(x >> 16)
			arr[width*i+2] = byte(x >> 8)
			arr[width*i+3] = byte(x >> 0)
			arr[width*i+4] = byte(y >> 24)
			arr[width*i+5] = byte(y >> 16)
			arr[width*i+6] = byte(y >> 8)
			arr[width*i+7] = byte(y >> 0)
			arr[width*i+8] = byte(z >> 24)
			arr[width*i+9] = byte(z >> 16)
			arr[width*i+10] = byte(z >> 8)
			arr[width*i+11] = byte(z >> 0)
		}
	}
	return arr
}

func Vec4(byteOrder binary.ByteOrder, values ...mgl32.Vec4) []byte {
	le := getByteOrder(byteOrder)
	width := 16
	arr := make([]byte, width*len(values))
	for i, v := range values {
		x := math.Float32bits(v.X())
		y := math.Float32bits(v.Y())
		z := math.Float32bits(v.Z())
		w := math.Float32bits(v.W())
		if le {
			arr[width*i+0] = byte(x >> 0)
			arr[width*i+1] = byte(x >> 8)
			arr[width*i+2] = byte(x >> 16)
			arr[width*i+3] = byte(x >> 24)
			arr[width*i+4] = byte(y >> 0)
			arr[width*i+5] = byte(y >> 8)
			arr[width*i+6] = byte(y >> 16)
			arr[width*i+7] = byte(y >> 24)
			arr[width*i+8] = byte(z >> 0)
			arr[width*i+9] = byte(z >> 8)
			arr[width*i+10] = byte(z >> 16)
			arr[width*i+11] = byte(z >> 24)
			arr[width*i+12] = byte(w >> 0)
			arr[width*i+13] = byte(w >> 8)
			arr[width*i+14] = byte(w >> 16)
			arr[width*i+15] = byte(w >> 24)
		} else {
			arr[width*i+0] = byte(x >> 24)
			arr[width*i+1] = byte(x >> 16)
			arr[width*i+2] = byte(x >> 8)
			arr[width*i+3] = byte(x >> 0)
			arr[width*i+4] = byte(y >> 24)
			arr[width*i+5] = byte(y >> 16)
			arr[width*i+6] = byte(y >> 8)
			arr[width*i+7] = byte(y >> 0)
			arr[width*i+8] = byte(z >> 24)
			arr[width*i+9] = byte(z >> 16)
			arr[width*i+10] = byte(z >> 8)
			arr[width*i+11] = byte(z >> 0)
			arr[width*i+12] = byte(w >> 24)
			arr[width*i+13] = byte(w >> 16)
			arr[width*i+14] = byte(w >> 8)
			arr[width*i+15] = byte(w >> 0)
		}
	}
	return arr
}

// NRGBA returns the byte representation of color.NRGBA values in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
// Note that while NRGBA consists of four uint8 values, this function returns
// four float32 values, since floating point colors are standard for shaders.
func NRGBA(byteOrder binary.ByteOrder, values ...color.NRGBA) []byte {
	le := getByteOrder(byteOrder)
	width := 16
	arr := make([]byte, width*len(values))
	for i, c := range values {
		r := math.Float32bits(float32(c.R) / 255.0)
		g := math.Float32bits(float32(c.G) / 255.0)
		b := math.Float32bits(float32(c.B) / 255.0)
		a := math.Float32bits(float32(c.A) / 255.0)
		if le {
			arr[width*i+0] = byte(r >> 0)
			arr[width*i+1] = byte(r >> 8)
			arr[width*i+2] = byte(r >> 16)
			arr[width*i+3] = byte(r >> 24)
			arr[width*i+4] = byte(g >> 0)
			arr[width*i+5] = byte(g >> 8)
			arr[width*i+6] = byte(g >> 16)
			arr[width*i+7] = byte(g >> 24)
			arr[width*i+8] = byte(b >> 0)
			arr[width*i+9] = byte(b >> 8)
			arr[width*i+10] = byte(b >> 16)
			arr[width*i+11] = byte(b >> 24)
			arr[width*i+12] = byte(a >> 0)
			arr[width*i+13] = byte(a >> 8)
			arr[width*i+14] = byte(a >> 16)
			arr[width*i+15] = byte(a >> 24)
		} else {
			arr[width*i+0] = byte(r >> 24)
			arr[width*i+1] = byte(r >> 16)
			arr[width*i+2] = byte(r >> 8)
			arr[width*i+3] = byte(r >> 0)
			arr[width*i+4] = byte(g >> 24)
			arr[width*i+5] = byte(g >> 16)
			arr[width*i+6] = byte(g >> 8)
			arr[width*i+7] = byte(g >> 0)
			arr[width*i+8] = byte(b >> 24)
			arr[width*i+9] = byte(b >> 16)
			arr[width*i+10] = byte(b >> 8)
			arr[width*i+11] = byte(b >> 0)
			arr[width*i+12] = byte(a >> 24)
			arr[width*i+13] = byte(a >> 16)
			arr[width*i+14] = byte(a >> 8)
			arr[width*i+15] = byte(a >> 0)
		}
	}
	return arr
}

// Uint32 returns the byte representation a uint32 array in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func Uint32(byteOrder binary.ByteOrder, values ...uint32) []byte {
	le := getByteOrder(byteOrder)
	width := 4
	b := make([]byte, 4*len(values))
	for i, v := range values {
		if le {
			b[width*i+0] = byte(v >> 0)
			b[width*i+1] = byte(v >> 8)
			b[width*i+2] = byte(v >> 16)
			b[width*i+3] = byte(v >> 24)
		} else {
			b[width*i+0] = byte(v >> 24)
			b[width*i+1] = byte(v >> 16)
			b[width*i+2] = byte(v >> 8)
			b[width*i+3] = byte(v >> 0)
		}
	}
	return b
}

// Uint16 returns the byte representation a uint16 array in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func Uint16(byteOrder binary.ByteOrder, values ...uint16) []byte {
	le := getByteOrder(byteOrder)
	width := 2
	b := make([]byte, 2*len(values))
	for i, v := range values {
		if le {
			b[width*i+0] = byte(v >> 0)
			b[width*i+1] = byte(v >> 8)
		} else {
			b[width*i+0] = byte(v >> 8)
			b[width*i+1] = byte(v >> 0)
		}
	}
	return b
}
