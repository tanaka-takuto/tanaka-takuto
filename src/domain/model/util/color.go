package util

import (
	"strings"

	"github.com/hisamafahri/coco"
)

// カラー
type Color struct {
	// フォントカラー
	FontColor string

	// イメージカラー
	ImageColor string
}

// JSONからアンマーシャルする
func (c *Color) UnmarshalJSON(bytes []byte) error {
	c.ImageColor = strings.ReplaceAll(string(bytes), "\"", "")

	rgb := coco.Hex2Rgb(c.ImageColor)
	gray := coco.Rgb2Gray(float64(rgb[0]), float64(rgb[1]), float64(rgb[2]))
	if gray[0] < 0x88 {
		c.FontColor = coco.Rgb2Hex(float64(0xff), float64(0xff), float64(0xff))
	} else {
		c.FontColor = coco.Rgb2Hex(float64(0x00), float64(0x00), float64(0x00))
	}

	return nil
}
