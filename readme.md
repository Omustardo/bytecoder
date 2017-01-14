bytecoder provides functions for converting various slices into byte slices.
Based on the Bytes function from https://github.com/golang/mobile/blob/master/exp/f32/f32.go

I have found this to be of use when working with OpenGL buffers:
```
// Store basic rectangle vertices in a buffer.
vertices := bytecoder.Float32(binary.LittleEndian,
    0, 0,
    1, 0,
    1, 1,
    0, 1,
)
buf := gl.CreateBuffer()
gl.BindBuffer(gl.ARRAY_BUFFER, buf)
gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
```
