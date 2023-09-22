package GoWas

import (
	"github.com/rocco-gossmann/GoWas/canvas"
	"github.com/rocco-gossmann/go_wasmcanvas"
	canvasfragments "github.com/rocco-gossmann/go_wasmcanvas/canvas_fragments"
)

type EngineCanvas struct {
	wasmcanvas go_wasmcanvas.Canvas
	engine     *Engine
	buffer     canvas.Buffer
	Mouse      MouseState
}

// Constructors
// ==============================================================================
func CreateCanvas(e *Engine, width, height uint16) *EngineCanvas {

	if width == 0 || height == 0 {
		panic("GoWas.Init(setup.WindowWidth and setup.WindowHeight) must be at least 1px")
	}

	ec := EngineCanvas{
		engine:     e,
		wasmcanvas: go_wasmcanvas.Create(width, height),
		buffer:     canvas.Buffer{PixelPerLine: width},
	}

	return &ec
}

// Methods
// ==============================================================================
func (ec *EngineCanvas) Run() {

	if ec == nil {
		panic("PANIC !!!!:  EngineCanvas is nil *runs in circles*")
	}

	ec.wasmcanvas.Run((*ec).canvasTick)
}

// Drawing Functions
// ==============================================================================
func (ec *EngineCanvas) FillRGBA(r, g, b, a byte) {
	if a > 0 {
		fill.Color = go_wasmcanvas.CombineRGB(r, g, b)
		fill.Alpha = a
		ec.wasmcanvas.Draw(&fill)
	}
}
func (ec *EngineCanvas) FillColorA(color uint32, a byte) {
	if a > 0 {
		fill.Color = go_wasmcanvas.Color(color)
		fill.Alpha = a
		ec.wasmcanvas.Draw(&fill)
	}
}
func (ec *EngineCanvas) BlitBitmap(bmp *canvas.Bitmap, x, y int32) {

	cw, ch := int32(ec.wasmcanvas.Width()), int32(ec.wasmcanvas.Height())
	bw, bh, bppl, pl := int32(bmp.Width()), int32(bmp.Height()), int32(bmp.PPL()), int(bmp.Pixels())

	//fmt.Println("drawing: ", x, y, cw, ch, bw, bh, bppl, pl)

	var bmpOffsetX, bmpOffsetY, bmpOverflowX, caOverflowX int32
	renderPPL := bppl

	// Trim BMP Lines from the Left
	if x < 0 {
		renderPPL += x
		bmpOffsetX -= x
		x = 0
	}

	// Trim BMP Lines from the Top
	if y < 0 {
		bmpOffsetY -= y
		y = 0
	}

	//fmt.Println("bmpOffset: ", bmpOffsetX, bmpOffsetY)

	// Check if any pixel is still on the canavs
	if (bmpOffsetX >= bw) || (bmpOffsetY >= bh) || (x >= cw) || (y >= ch) {
		//fmt.Println("bmp not on screen: ", (bmpOffsetX >= bw), (bmpOffsetY >= bh), (x >= cw), (y <= ch))
		return
	}
	//fmt.Println("bmp on screen")

	// Trim BMP from the Right
	if x > cw-bw {
		//fmt.Println("bmp trim right", cw, bw, x, cw-bw, x > cw-bw)
		bmpOverflowX = (x + bw - cw)
		renderPPL -= bmpOverflowX
	}

	// Trim BMP from the Bottom
	if y > ch-bh {
		//fmt.Println("bmp trim height", ch, bh, y, ch-bh, y > ch-bh)
		pl -= int((y + bh - ch) * bw)
	}

	caOverflowX = cw - renderPPL
	bmpOverflowX += bmpOffsetX

	// Prepare Memory pointers
	var (
		bmpPtr = (bmpOffsetY * bppl) + bmpOffsetX
		caPtr  = y*cw + x
	)

	// Start to walk
	for int(bmpPtr) < pl { //<- until the pointer reached the area the end of the last BMP line to draw

		//fmt.Println("Draw Line", bmpPtr, caPtr, renderPPL, caOverflowX, bmpOverflowX)
		for i := int32(0); i < renderPPL; i++ { //<- Render all Pixels to draw for the line
			(*((*ec).buffer.Memory))[caPtr] = (*((*bmp).MemoryBuffer.Memory))[bmpPtr]

			bmpPtr++ //<- move both pointers formward by one
			caPtr++
			//fmt.Println("drawn pixel", bmpPtr, caPtr)
		}

		caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
		bmpPtr += bmpOverflowX

	}
}

// Private Helpers
// ==============================================================================
var fill = canvasfragments.Fill{}

func (ec *EngineCanvas) canvasDraw(c uint32, w, h uint16, px *[]uint32) {
	(*ec).buffer.Memory = px
	(*(*(*ec).engine).Draw).Draw((*ec).engine, ec)
}

func (ec *EngineCanvas) canvasTick(c *go_wasmcanvas.Canvas, deltaTime float64) go_wasmcanvas.CanvasTickFunction {

	ec.Mouse = *UpdateMouse()

	engine := &(*(*ec).engine)
	if (*(*engine).Tick).Tick(engine, deltaTime) {
		ec.wasmcanvas.Apply(ec.canvasDraw)

	} else {
		engine.switchScene((*(*engine).Unload).Unload(engine))

	}

	return (*ec).canvasTick
}
