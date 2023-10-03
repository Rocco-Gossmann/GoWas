# PNG 2 GoWas-Bitmap

this script is meant to convert png images into GoWas Bitmaps ( who could have guessed).

Any RGBA16 Png should be fine.
Any pixel that is not 100% opaque will be ignored.

The colors themselfs will be recalculated to fit a 24 Bit RGB Color-Space. 

## Usage:
```bash
go run ./tools/png2gowasbmp.go   /path/source.png   /path/output.go   gopackagename  GoBitmapName
```

| Param              | Description                                                                                                                                                                                                                 |
|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/path/source.png` | The PNG-File that should be converted                                                                                                                                                                                       |
| `/path/output.go`  | the Go-File that is genrated and contains the Bitmap-Definition<br>For now Bitmaps are part of the main-executable (similar to how it would be in a rom <br> Loading ressources from the browser has yet to be implemented) |
| `gopackagename`    | The name of the package, tha the Generated Go file belongs to                                                                                                                                                               |
| `GoBitmapName`     | The Name of the Variable, that makes the Bitmap accessabel from within Go                                                                                                                                                   |


## Compression
Each Pixel/PixelGroup will be represented in the following way.

```
Binary:  0000000 a rrrrrrrr gggggggg bbbbbbbb
            |    |  |        |        |jjjjjjjjj
            |    |  |        |        ---Blue Color 0 - 255
            |    |  |        ------------Green Color 0 - 255
            |    |   --------------------Red Color 0-255
            |    ------------------------Alpha Value 1 = visible 0 = invisible
            -----------------------------Repeat 0-127 (0 = draw only once; 127 = draw 128 times)
```
In a compressed BMP definition, each Pixel starts with 7 bits, that tell how often the pixel is repeated.
In a sesse this you are defining lines of pixels, rathern than pixels themself.

so `0x08000000` would be equivalent to   
`0x00000000 0x00000000 0x00000000 0x00000000`   

`0x09ffffff` would become   
`0x01ffffff 0x01ffffff 0x01ffffff 0x01ffffff`   







