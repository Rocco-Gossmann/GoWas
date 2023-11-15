# Alpha-Blending

Various graphics elements of GoWas can be rendered with Transparency.

<!-- TOC -->

- [Alpha-Blending](#alpha-blending)
- [Methods](#methods)
    - [AlphaSet](#alphaset)
    - [AlphaReset](#alphareset)

<!-- /TOC -->


This Transparency is set via the Methods:

```go
func (*GraphicsElement) AlphaSet(a core.CanvasAlpha)
//and 
func (*GraphicsElement) AlphaReset()
```

Valid values for `core.CanvasAlpha` are:

| Constant          | Number-Value |
|-------------------|--------------|
| `CANV_ALPHA_NONE` | 0x00         |
| `CANV_ALPHA_0   ` | 0x00         |
| `CANV_ALPHA_1   ` | 0x06         |
| `CANV_ALPHA_2   ` | 0x05         |
| `CANV_ALPHA_3   ` | 0x04         |
| `CANV_ALPHA_4   ` | 0x03         |
| `CANV_ALPHA_5   ` | 0x02         |
| `CANV_ALPHA_6   ` | 0x01         |
| `CANV_ALPHA_7   ` | 0x07         |
| `CANV_ALPHA_FULL` | 0x07         |

# Methods
## AlphaSet
```go
func (*GraphicsElement) AlphaSet(a core.CanvasAlpha) *GraphicsElement
```
### todo


## AlphaReset
```go
func (*GraphicsElement) AlphaReset() *GraphicsElement
```
### todo

