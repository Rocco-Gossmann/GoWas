# Using PNG Files

For now It is only possible to load PNG-Files.\
At the moment, these are converted into `*.go` files and become part of the WASM
binary.

The Conversion process is handled by the `Makefile` as long, as you follow the
following steps.

> [!warning] 
> To follow this process, you need to have finished the [Setup](./Setup.md) first.

1. Put any PNG, you wish to become part of the Project into the `./assets` -
   Folder in your Project

2. run the `make` command

`make` should now pick up any new or changed PNG-File and create a
`./bmps/bmp.[PNGFileName].go` File.

3. You can then access the PNG via the following way

```go
import (
    "github.com/rocco-gossmann/GoWas/core"
    "GoWasProject/bmps"
)

var myPNG *core.Bitmap = bmps.BMPPngFileName
```

The Bitmap is available via a `*core.Bitmap` pointer.

### Filename-Conversion

The PNGs Filename is converted into a VariableName, but due to the nature of
variable names, not all characters of a filename can be put into a
variable name, theirefore all characters, that don't match the following regexp
are stripped from the file name `[A-Za-z0-9_]`

Here are a few examples of file- to variable name conversions

| Asset-File                       | Go-File                       | Go-Variable          |
| -------------------------------- | ----------------------------- | -------------------- |
| `./assets/cursor.png`            | `./bmps/bmp.cursor.go`        | `BMPcursor`          |
| `./assets/map 1.png`             | `./bmps/bmp.map 1.go`         | `BMPmap1`            |
| `./assets/1-SpriteSheet.new.png` | `./bmps/1-SpriteSheet.new.go` | `BMP1SpriteSheetnew` |


> [!todo] 
> Implement Sound
