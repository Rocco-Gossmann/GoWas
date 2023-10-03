package core

import (
	"github.com/rocco-gossmann/GoWas/types"
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

	//Bits 0x10000000 and 0x80000000 are reserved for tile-collision layers

	CANV_CL_NONE CanvasCollisionLayers = 0
	CANV_CL_ALL  CanvasCollisionLayers = CANV_CL_1 | CANV_CL_2 | CANV_CL_3 | CANV_CL_4
)

type BlitSettings struct {
	Bmp       *Bitmap               // What to blit
	X, Y      int32                 // Where to blit it on the screen
	Alpha     byte                  // how strong transparency is
	Alphazero bool                  // if true, an alpha value of 0 mean "draw nothing", otherwise 0 would mean ignore alpha
	Layers    CanvasCollisionLayers // What collision layers the drawn object occupies
	Clip      *types.Rect           // Clipping Rectangle to only draw a certain area of the bitmap
}
type Canvas struct {
	wasmcanvas go_wasmcanvas.Canvas
	engine     *Engine
	buffer     Buffer
	Mouse      MouseState
}

func (ca *Canvas) GetWidth() uint16  { return ca.wasmcanvas.Width() }
func (ca *Canvas) GetHeight() uint16 { return ca.wasmcanvas.Height() }

// Constructors
// ==============================================================================
func CreateCanvas(e *Engine, width, height uint16) *Canvas {

	if width == 0 || height == 0 {
		panic("GoWas.Init(setup.WindowWidth and setup.WindowHeight) must be at least 1px")
	}

	ec := Canvas{
		engine:     e,
		wasmcanvas: go_wasmcanvas.Create(width, height),
		buffer:     Buffer{PixelPerLine: width},
	}

	return &ec
}

// Methods
// ==============================================================================
func (ec *Canvas) Run() {

	if ec == nil {
		panic("PANIC !!!!:  EngineCanvas is nil *runs in circles*")
	}

	ec.wasmcanvas.Run((*ec).canvasTick)
}

// Drawing Functions
// ==============================================================================
func (ec *Canvas) FillRGBA(r, g, b, alpha byte, layerReset CanvasCollisionLayers) {
	if alpha > 0 {
		fillJob.Color = uint32(go_wasmcanvas.CombineRGB(r, g, b))
		fillJob.Alpha = alpha
		fillJob.layers = layerReset
		ec.wasmcanvas.Draw(&fillJob)
	}
}
func (ec *Canvas) FillColorA(color uint32, alpha byte, layerReset CanvasCollisionLayers) {
	if alpha > 0 {
		fillJob.Color = color
		fillJob.Alpha = alpha
		fillJob.layers = layerReset
		ec.wasmcanvas.Draw(&fillJob)
	}
}

func (ec *Canvas) Blit(opts *BlitSettings) CanvasCollisionLayers {

	if opts.Bmp == nil {
		panic("nothing to blit")
	}

	// Handle Alpha
	if opts.Alpha == 0 && !opts.Alphazero {
		opts.Alpha = 0xff
	}

	// Set Clipping Bounderys
	bw, bh := opts.Bmp.Width(), opts.Bmp.Height()

	if opts.Clip == nil {
		opts.Clip = &types.Rect{0, 0, bw, bh}

	} else {
		// If clip starts outside of BMP === no render
		if (*opts.Clip).X >= opts.Bmp.Width() || (*opts.Clip).Y >= opts.Bmp.Height() {
			return CANV_CL_NONE
		}

		// No W or H == take W and H from Bitmap
		if (*opts.Clip).W == 0 {
			(*opts.Clip).W = bw
		}

		if (*opts.Clip).H == 0 {
			(*opts.Clip).H = bh
		}

	}

	// Check right and bottom Clip for overflows
	br, bb := (*opts.Clip).X+opts.Clip.W, (*opts.Clip).Y+opts.Clip.H

	// If Clipping zone overflows right => cut overflow off
	if br >= bw {
		(*opts.Clip).W -= (br - bw)
	}

	// If Clipping zone overflows bottom => cut overflow off
	if bb >= bh {
		(*opts.Clip).H -= (bb - bh)
	}

	return ec.blitBitmapClipped(opts.Bmp, opts.X, opts.Y, opts.Alpha, opts.Layers, opts.Clip)
}

func (ec *Canvas) BlitBitmap(bmp *Bitmap, x, y int32, alpha byte, layers CanvasCollisionLayers) CanvasCollisionLayers {
	return ec.blitBitmapClipped(bmp, x, y, alpha, layers, &types.Rect{0, 0, bmp.Width(), bmp.Height()})
}

// Implement go_wasm_canvas
// ==============================================================================
func (ec *Canvas) canvasDraw(c uint32, w, h uint16, px *[]uint32) {
	(*ec).buffer.Memory = px
	(*(*(*ec).engine).Draw).Draw((*ec).engine, ec)
}

func (ec *Canvas) canvasTick(c *go_wasmcanvas.Canvas, deltaTime float64) go_wasmcanvas.CanvasTickFunction {

	ec.Mouse = *UpdateMouse()

	engine := &(*(*ec).engine)
	if (*(*engine).Tick).Tick(engine, deltaTime) {
		ec.wasmcanvas.Apply(ec.canvasDraw)

	} else {
		engine.SwitchScene((*(*engine).Unload).Unload(engine))

	}

	return (*ec).canvasTick
}

// Private Helpers
// ==============================================================================
func (ec *Canvas) blitBitmapClipped(bmp *Bitmap, x, y int32, alpha byte, layers CanvasCollisionLayers, clip *types.Rect) CanvasCollisionLayers {

	//what to draw
	var bitmapByteOffset uint32 = uint32((*clip).X)
	var bitmapIndexStart uint32 = uint32((*clip).Y)*uint32(bmp.PPL()) + uint32(bitmapByteOffset)
	var bitmapRenderLinePixels uint32 = uint32((*clip).W)
	bitmapByteOffset += uint32(uint32(bmp.Width()) - bitmapByteOffset - bitmapRenderLinePixels)

	var bitmapRenderLines int32 = int32((*clip).H)

	cw, ch := int32(ec.wasmcanvas.Width()), int32(ec.wasmcanvas.Height())
	bw, bh := int32(clip.W), int32(clip.H)

	// if the bmp-clip is to far off of one of the four canvas sides
	// => Render nothing
	if y+bh <= 0 || x+bw <= 0 || x >= cw || y >= ch {
		return CANV_CL_NONE
	}

	// Trim BMP Lines from the Left
	if x < 0 {
		bitmapIndexStart = uint32(int32(bitmapIndexStart) - x)
		bitmapByteOffset = uint32(int32(bitmapByteOffset) - x)
		bitmapRenderLinePixels = uint32(int32(bitmapRenderLinePixels) + x)
		x = 0
	}

	// Trim BMP Lines from the Top
	if y < 0 {
		bitmapIndexStart = uint32(int32(bitmapIndexStart) - (y * int32(bmp.Width())))
		bitmapRenderLines += y
		y = 0
	}

	// Trim BMP Lines from the Right
	if x > cw-bw {
		bmpOverflowX := (x + bw - cw)
		bitmapByteOffset = uint32(int32(bitmapByteOffset) + bmpOverflowX)
		bitmapRenderLinePixels -= uint32(bmpOverflowX)
	}

	// Trim BMP from the Bottom
	if y > ch-bh {
		bitmapRenderLines -= (y + bh - ch)
	}

	// Prepare Memory pointers
	var (
		bmpPtr      = bitmapIndexStart
		caPtr       = y*cw + x
		caOverflowX = cw - int32(bitmapRenderLinePixels)
	)

	var outbyte CanvasCollisionLayers = CANV_CL_NONE

	//	fmt.Println(bitmapByteOffset, bitmapIndexStart, bitmapRenderLinePixels, bitmapRenderLines, "\n", cw, ch, bw, bh, "\n", bmpPtr, caPtr, caOverflowX, "\n", alpha)
	if alpha == 0x00 {
		// Only process meta data
		//================================================================================
		for line := int32(0); line < bitmapRenderLines; line++ {
			for i := uint32(0); i < bitmapRenderLinePixels; i++ { //<- Render all Pixels to draw for the line

				// Only modify pixels Meta data
				var cpx = (*(ec.buffer.Memory))[caPtr]
				outbyte |= CanvasCollisionLayers(cpx & uint32(CANV_CL_ALL))
				cpx |= uint32(layers) * (cpx & uint32(CANV_CL_ALL) >> 24)
				(*(ec.buffer.Memory))[caPtr] = cpx

				bmpPtr++ //<- move both pointers formward by one
				caPtr++
				//fmt.Println("drawn pixel", bmpPtr, caPtr)
			}

			caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
			bmpPtr += uint32(bitmapByteOffset)
		}
	} else if alpha == 0xff {
		// Draw Full Pixel + Meta Data
		//================================================================================
		for line := int32(0); line < bitmapRenderLines; line++ {
			for i := uint32(0); i < bitmapRenderLinePixels; i++ { //<- Render all Pixels to draw for the line

				var cpx = (*(ec.buffer.Memory))[caPtr]
				outbyte |= CanvasCollisionLayers(cpx & uint32(CANV_CL_ALL))

				var transparencybit = ((*(bmp.MemoryBuffer.Memory))[bmpPtr] & uint32(BMP_OPAQUE)) >> 24
				cpx |= uint32(layers) * transparencybit
				var transparencyinvers = (transparencybit ^ 1)

				var px = cpx*transparencyinvers +
					(*(bmp.MemoryBuffer.Memory))[bmpPtr]*transparencybit

				(*(ec.buffer.Memory))[caPtr] = px

				bmpPtr++ //<- move both pointers formward by one
				caPtr++
				//fmt.Println("drawn pixel", bmpPtr, caPtr)
			}

			caPtr += caOverflowX //<- reset canvas Pointer to next lines X Coord
			bmpPtr += uint32(bitmapByteOffset)
		}
	} else {
		// Blend Pixel + Meta Data
		//================================================================================
		factor := float64(alpha) / 255.0

		for line := int32(0); line < bitmapRenderLines; line++ {
			for i := uint32(0); i < bitmapRenderLinePixels; i++ { //<- Render all Pixels to draw for the line

				var cpx = (*((*ec).buffer.Memory))[caPtr]

				opaque := ((*((*bmp).MemoryBuffer.Memory))[bmpPtr] & uint32(BMP_OPAQUE)) >> 24

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
			bmpPtr += uint32(bitmapByteOffset)
		}
	}

	return outbyte
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
