package aimage

import (
	"errors"
	"github.com/asktop/gotools/arand"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path"
)

/*
* 缩略图生成
* 入参:
* 规则: 如果width 或 hight其中有一个为0，则大小不变 如果精度quality为0则精度保持不变；tmpThumbDir缩略图缓存文件夹
* 矩形坐标系起点是左上
* 返回:error
 */
func ThumbFileHeader(fileHeader *multipart.FileHeader, width, height, quality int, tmpThumbDir string) (tmpThumbPath string, err error) {
	in, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer in.Close()

	origin, typ, err := image.Decode(in)
	if err != nil {
		return
	}
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	tmpThumbName := arand.RandUUID() + "-" + fileHeader.Filename
	tmpThumbPath = path.Join(tmpThumbDir, tmpThumbName)

	out, _ := os.Create(tmpThumbPath)
	defer out.Close()

	switch typ {
	case "jpeg":
		err = jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		err = png.Encode(out, canvas)
	case "gif":
		err = gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		err = bmp.Encode(out, canvas)
	default:
		err = errors.New("Thumb IMG ERROR FORMAT")
	}
	return
}
