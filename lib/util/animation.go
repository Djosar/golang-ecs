package util

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames         []*ebiten.Image
	Marker         int
	Counter        int
	AnimationSpeed int
	FrameIndex     int
}

func NewAnimation(frames []*ebiten.Image, marker int, speed int) *Animation {
	return &Animation{
		Frames:         frames,
		Marker:         marker,
		AnimationSpeed: speed,
		Counter:        0,
		FrameIndex:     0,
	}
}

func (a *Animation) GetCurrentFrame() *ebiten.Image {

	return a.Frames[a.FrameIndex]
}

func GenerateFrames(
	sprite *ebiten.Image,
	tileWidth, tileHeight int,
	tileCoordinates []*Coordinate[int],
) (frames []*ebiten.Image) {
	dimensions := sprite.Bounds().Size()

	for _, coord := range tileCoordinates {
		frameStartX := coord.X * tileWidth
		frameStartY := coord.Y * tileHeight

		if frameStartX <= dimensions.X && frameStartY <= dimensions.Y {
			frameEndX := frameStartX + tileWidth
			frameEndY := frameStartY + tileWidth

			if frameEndX > dimensions.X {
				frameEndX = dimensions.X
			}
			if frameEndY > dimensions.Y {
				frameEndY = dimensions.Y
			}

			frame := sprite.SubImage(image.Rect(
				frameStartX,
				frameStartY,
				frameEndX,
				frameEndY,
			))

			frames = append(frames, ebiten.NewImageFromImage(frame))
		}
	}

	return frames
}
