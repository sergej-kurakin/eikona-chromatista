package processor

import "image/color"

func calculateLuminocity(originalColor color.RGBA) uint8 {
	// Y = 0.2126 R + 0.7152 G + 0.0722 B
	return uint8(0.2126*float32(originalColor.R) + 0.7152*float32(originalColor.G) + 0.0722*float32(originalColor.B))
}
