package core

import "github.com/rocco-gossmann/GoWas/types"

type BitmapEntity struct {
	data *Bitmap

	// Internals
	/* clipping rect */
	clip *types.Rect

	/* Toggled by Clipping Rect size */
	visible bool

	canvasBlitOpts CanvasBlitOpts
}

// ------------------------------------------------------------------------------
// Setters
// ------------------------------------------------------------------------------
func (me *BitmapEntity) Clip(offsetX, offsetY, width, height uint16) *BitmapEntity {

	me.canvasBlitOpts.Clip.X = offsetX
	me.canvasBlitOpts.Clip.Y = offsetY

	me.canvasBlitOpts.Clip.W = min(width, me.data.width-offsetX)
	me.canvasBlitOpts.Clip.H = min(height, me.data.height-offsetY)

	me.visible = !(me.clip.W == 0 || me.clip.H == 0)

	return me
}

func (me *BitmapEntity) Alpha(alpha byte, zeroIsInvisible bool) *BitmapEntity {

	me.canvasBlitOpts.Alpha = alpha
	me.canvasBlitOpts.Alphazero = zeroIsInvisible

	return me
}

func (me *BitmapEntity) CollisionLayers(layers CanvasCollisionLayers) *BitmapEntity {
	me.canvasBlitOpts.Layers = layers

	return me
}

func (me *BitmapEntity) MoveTo(x, y int32) *BitmapEntity {

	me.canvasBlitOpts.X = x
	me.canvasBlitOpts.Y = y

	return me
}
func (me *BitmapEntity) MoveBy(x, y int32) *BitmapEntity {

	me.canvasBlitOpts.X += x
	me.canvasBlitOpts.Y += y

	return me
}

// ------------------------------------------------------------------------------
// Getter
// ------------------------------------------------------------------------------
// Internals
func (b *BitmapEntity) UnclippedWidth() uint16  { return b.data.width }
func (b *BitmapEntity) UnclippedHeight() uint16 { return b.data.height }
func (b *BitmapEntity) PPL() uint16             { return b.data.MemoryBuffer.PixelPerLine }
func (b *BitmapEntity) Pixels() int             { return len(*(b.data.MemoryBuffer).Memory) }

// Display
func (b *BitmapEntity) X() int32           { return b.canvasBlitOpts.X }
func (b *BitmapEntity) Y() int32           { return b.canvasBlitOpts.Y }
func (b *BitmapEntity) XY() (int32, int32) { return b.canvasBlitOpts.X, b.canvasBlitOpts.Y }

func (b *BitmapEntity) W() uint16            { return b.canvasBlitOpts.Clip.W }
func (b *BitmapEntity) H() uint16            { return b.canvasBlitOpts.Clip.H }
func (b *BitmapEntity) WH() (uint16, uint16) { return b.canvasBlitOpts.Clip.W, b.canvasBlitOpts.Clip.H }

// ------------------------------------------------------------------------------
// Actions
// ------------------------------------------------------------------------------
func (b *BitmapEntity) ToCanvas(ca *Canvas) CanvasCollisionLayers {
	if !b.visible {
		return CANV_CL_NONE
	}

	return ca.Blit(&b.canvasBlitOpts)
}
