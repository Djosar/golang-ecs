package util

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation represents a sequence of frames for an animation. It contains the
// frames, marker, counter, animation speed, and the current frame index.
type Animation struct {
	Frames         []*ebiten.Image
	Marker         int
	Counter        int
	AnimationSpeed int
	FrameIndex     int
}

// NewAnimation creates and returns a new Animation instance.
//
// Parameters:
//
//	frames ([]*ebiten.Image): The frames of the animation.
//	marker (int): The marker indicating the end of the animation sequence.
//	speed (int): The speed of the animation.
//
// Returns:
//
//	*Animation: A pointer to the newly created Animation instance.
func NewAnimation(frames []*ebiten.Image, marker int, speed int) *Animation {
	return &Animation{
		Frames:         frames,
		Marker:         marker,
		AnimationSpeed: speed,
		Counter:        0,
		FrameIndex:     0,
	}
}

// GetCurrentFrame returns the current frame of the animation.
//
// Returns:
//
//	*ebiten.Image: The current frame of the animation.
func (a *Animation) GetCurrentFrame() *ebiten.Image {
	return a.Frames[a.FrameIndex]
}

// GenerateFrames generates a sequence of frames from a sprite sheet based on the
// provided tile coordinates.
//
// Parameters:
//
//	sprite (*ebiten.Image): The sprite sheet image.
//	tileWidth (int): The width of each tile/frame.
//	tileHeight (int): The height of each tile/frame.
//	tileCoordinates ([]*Coordinate[int]): The coordinates of the tiles to be extracted.
//
// Returns:
//
//	[]*ebiten.Image: A slice of ebiten.Image pointers representing the frames.
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
