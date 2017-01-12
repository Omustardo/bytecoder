bytecoder provides functions for converting standard arrays into byte arrays.
Based on the Bytes function from https://github.com/golang/mobile/blob/master/exp/f32/f32.go

I have found this to be of use when working with OpenGL buffers:
```
// Store basic rectangle vertices in a buffer.
lower, upper := float32(-0.5), float32(0.5)
vertices := bytecoder.Float32(binary.LittleEndian,
    lower, lower,
    upper, lower,
    upper, upper,
    lower, upper,
)
buf := gl.CreateBuffer()
gl.BindBuffer(gl.ARRAY_BUFFER, buf)
gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
```
