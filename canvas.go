package eng

import (
	gl "github.com/chsc/gogl/gl33"

	"image"
	"log"
)

type Canvas struct {
	id      gl.Uint
	texture *Texture
	width   int
	height  int
}

func NewCanvas(width, height int) *Canvas {
	canvas := new(Canvas)
	canvas.width = width
	canvas.height = height

	canvas.texture = NewTexture(image.NewRGBA(image.Rect(0, 0, width, height)))
	canvas.texture.SetFilter(FilterLinear, FilterLinear)
	canvas.texture.SetWrap(WrapClampToEdge, WrapClampToEdge)

	gl.GenFramebuffers(1, &canvas.id)

	canvas.texture.Bind()
	gl.BindFramebuffer(gl.FRAMEBUFFER, canvas.id)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, canvas.texture.id, 0)

	result := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)

	canvas.texture.Unbind()
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	if result != gl.FRAMEBUFFER_COMPLETE {
		gl.DeleteFramebuffers(1, &canvas.id)
		log.Fatal("canvas couldn't be constructed")
	}

	return canvas
}

func (c *Canvas) Begin() {
	gl.Viewport(0, 0, gl.Sizei(c.texture.Width()), gl.Sizei(c.texture.Height()))
	gl.BindFramebuffer(gl.FRAMEBUFFER, c.id)
}

func (c *Canvas) End() {
	gl.Viewport(0, 0, gl.Sizei(Width()), gl.Sizei(Height()))
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (c *Canvas) Texture() *Texture {
	return c.texture
}

func (c *Canvas) Width() int {
	return c.texture.Width()
}

func (c *Canvas) Height() int {
	return c.texture.Height()
}
