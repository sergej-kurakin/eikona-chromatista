package processor

import "image/color"

type ColorProcessor func(color.Color) (color.Color, error)
