package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
	width int
	height int
}

// 사용할 컬러 모델 지정
func (img *Image) ColorModel() color.Model {
    return color.RGBAModel
}

// 이미지의 크기 지정
func (img *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

// 이미지의 각 x, y 좌표마다 색 지정
func (img *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), 255, 255, uint8(y)}
}

func main() {
    m := &Image{256, 256}
    pic.ShowImage(m)
}
