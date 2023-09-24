package GoWas

import (
	"github.com/rocco-gossmann/GoWas/canvas"
	"github.com/rocco-gossmann/go_wasmcanvas"
)

type CanvasFlag uint32
type CanvasCollisionLayers uint32

const (
	// Collisiion Layers These layers are processed when a sprite is drawn
	// The Blit function will return a byte that contains all collision layers, that
	// already contained pixels, during drawing
	// [!NOTICE] regalrdless of BMP_TRANSPARENCY or BMP_FRONT this bits are always processed
	CANV_CL_1 CanvasCollisionLayers = 0x01000000
	CANV_CL_2 CanvasCollisionLayers = 0x02000000
	CANV_CL_3 CanvasCollisionLayers = 0x04000000
	CANV_CL_4 CanvasCollisionLayers = 0x08000000
	CANV_CL_5 CanvasCollisionLayers = 0x10000000
	CANV_CL_6 CanvasCollisionLayers = 0x20000000

	//Bits 0x40000000 and 0x80000000 are reserved for now

	CANV_CL_NONE CanvasCollisionLayers = 0
	CANV_CL_ALL  CanvasCollisionLayers = CANV_CL_1 | CANV_CL_2 | CANV_CL_3 | CANV_CL_4 | CANV_CL_5 | CANV_CL_6
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
func (ec *EngineCanvas) FillRGBA(r, g, b, alpha byte, layerReset CanvasCollisionLayers) {
	if alpha > 0 {
		fillJob.Color = uint32(go_wasmcanvas.CombineRGB(r, g, b))
		fillJob.Alpha = alpha
		fillJob.layers = layerReset
		ec.wasmcanvas.Draw(&fillJob)
	}
}
func (ec *EngineCanvas) FillColorA(color uint32, alpha byte, layerReset CanvasCollisionLayers) {
	if alpha > 0 {
		fillJob.Color = color
		fillJob.Alpha = alpha
		fillJob.layers = layerReset
		ec.wasmcanvas.Draw(&fillJob)
	}
}

type BlitSettings struct {
	Bmp       *canvas.Bitmap        // What to blit
	X, Y      int32                 // Where to blit it on the screen
	Alpha     byte                  // how strong transparency is
	Alphazero bool                  // if true, an alpha value of 0 mean "draw nothing", otherwise 0 would mean ignore alpha
	Layers    CanvasCollisionLayers // What collision layers the drawn object occupies
}

func (ec *EngineCanvas) Blit(opts BlitSettings) CanvasCollisionLayers {
	if opts.Bmp == nil {
		panic("nothing to blit")
	}

	if opts.Alpha == 0 && !opts.Alphazero {
		opts.Alpha = 0xff
	}

	return ec.BlitBitmap(opts.Bmp, opts.X, opts.Y, opts.Alpha, opts.Layers)
}

func (ec *EngineCanvas) BlitBitmap(bmp *canvas.Bitmap, x, y int32, alpha byte, layers CanvasCollisionLayers) CanvasCollisionLayers {

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
		return 0
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

	var outbyte CanvasCollisionLayers = CANV_CL_NONE

	if alpha == 0x00 {
		// Only process meta data
		//================================================================================
		for int(bmpPtr) < pl { //<- until the pointer reached the area the end of the last BMP line to draw

			//fmt.Println("Draw Line", bmpPtr, caPtr, renderPPL, caOverflowX, bmpOverflowX)
			for i := int32(0); i < renderPPL; i++ { //<- Render all Pixels to draw for the line

				// Only modify pixels Meta data
				var cpx = (*((*ec).buffer.Memory))[caPtr]
				outbyte |= CanvasCollisionLayers(cpx & uint32(CANV_CL_ALL))
				cpx |= uint32(layers) * (cpx & uint32(CANV_CL_ALL) >> 24)
				(*((*ec).buffer.Memory))[caPtr] = cpx

				bmpPtr++ //<- move both pointers formward by one
				caPtr++
				//fmt.Println("drawn pixel", bmpPtr, caPtr)
			}

			caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
			bmpPtr += bmpOverflowX
		}
	} else if alpha == 0xff {
		// Draw Full Pixel + Meta Data
		//================================================================================
		for int(bmpPtr) < pl { //<- until the pointer reached the area the end of the last BMP line to draw

			//fmt.Println("Draw Line", bmpPtr, caPtr, renderPPL, caOverflowX, bmpOverflowX)
			for i := int32(0); i < renderPPL; i++ { //<- Render all Pixels to draw for the line

				var cpx = (*((*ec).buffer.Memory))[caPtr]
				outbyte |= CanvasCollisionLayers(cpx & uint32(CANV_CL_ALL))

				var transparencybit = ((*((*bmp).MemoryBuffer.Memory))[bmpPtr] & uint32(canvas.BMP_OPAQUE)) >> 24
				cpx |= uint32(layers) * transparencybit
				var transparencyinvers = (transparencybit ^ 1)

				var px = cpx*transparencyinvers +
					(*((*bmp).MemoryBuffer.Memory))[bmpPtr]*transparencybit

				(*((*ec).buffer.Memory))[caPtr] = px

				bmpPtr++ //<- move both pointers formward by one
				caPtr++
				//fmt.Println("drawn pixel", bmpPtr, caPtr)
			}

			caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
			bmpPtr += bmpOverflowX
		}
	} else {
		// Blend Pixel + Meta Data
		//================================================================================
		factor := float64(alpha) / 255.0

		for int(bmpPtr) < pl { //<- until the pointer reached the area the end of the last BMP line to draw

			//fmt.Println("Draw Line", bmpPtr, caPtr, renderPPL, caOverflowX, bmpOverflowX)
			for i := int32(0); i < renderPPL; i++ { //<- Render all Pixels to draw for the line

				var cpx = (*((*ec).buffer.Memory))[caPtr]

				opaque := ((*((*bmp).MemoryBuffer.Memory))[bmpPtr] & canvas.BMP_OPAQUE) >> 24

				// Mixel Meta
				outbyte |= CanvasCollisionLayers(cpx & uint32(CANV_CL_ALL))

				go_wasmcanvas.BlendPixel(
					&cpx,
					(*((*bmp).MemoryBuffer.Memory))[bmpPtr],
					float64(opaque)*factor,
				)

				cpx |= uint32(layers) * opaque

				(*((*ec).buffer.Memory))[caPtr] = cpx

				bmpPtr++ //<- move both pointers formward by one
				caPtr++
				//fmt.Println("drawn pixel", bmpPtr, caPtr)
			}

			caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
			bmpPtr += bmpOverflowX
		}

	}

	return outbyte
}

// Private Helpers
// ==============================================================================

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

// Fill-Jobs
// ==============================================================================
type sFillJob struct {
	layers CanvasCollisionLayers
	Color  uint32
	Alpha  byte
}

var fillJob = sFillJob{}

func (f *sFillJob) Draw(pc uint32, _ uint16, _ uint16, pxs *[]uint32) {

	if f.Alpha == 0 {
		return
	}
	var resetMask = ^(uint32(f.layers) << 24)
	var resetcolor = (f.Color & 0x00ffffff)
	if f.Alpha == 0xff {
		for i := uint32(0); i < pc; i++ {
			(*pxs)[i] = (((*pxs)[i] & 0xff000000) | resetcolor) & resetMask
		}

	} else {
		factor := float64(f.Alpha) / 255.0
		for i := uint32(0); i < pc; i++ {
			resetPixel := (*pxs)[i]
			go_wasmcanvas.BlendPixel(&resetPixel, resetcolor, factor)
			(*pxs)[i] = (((*pxs)[i] & 0xff000000) | resetPixel) & resetMask
		}
	}

}
