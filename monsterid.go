package monsterid

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
)

var (
	legs  = 5
	hair  = 5
	arms  = 5
	body  = 15
	eyes  = 15
	mouth = 10
)

var bodyParts = []string{"legs", "hair", "arms", "body", "eyes", "mouth"}

type monsterid struct {
	legs  int
	hair  int
	arms  int
	body  int
	eyes  int
	mouth int
}

func New(seed []byte) image.Image {
	buf := sha512.Sum512(seed)
	seed_int := binary.BigEndian.Uint64(buf[56:])
	rand.Seed(int64(seed_int))
	mid := &monsterid{}
	mid.legs = rand.Intn(legs) + 1
	mid.hair = rand.Intn(hair) + 1
	mid.arms = rand.Intn(arms) + 1
	mid.body = rand.Intn(body) + 1
	mid.eyes = rand.Intn(eyes) + 1
	mid.mouth = rand.Intn(mouth) + 1

	img := image.NewRGBA(image.Rect(0, 0, 120, 120))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for _, part := range bodyParts {
		fileName := fmt.Sprintf("parts/%s_%d.png", part, getPartNumber(mid, part))
		asset, err := Asset(fileName)
		if err != nil {
			log.Fatal(err)
		}
		assetFull, err := png.Decode(bytes.NewReader(asset))
		if err != nil {
			log.Fatal(err)
		}
		draw.Draw(img, img.Bounds(), assetFull, image.ZP, draw.Over)
	}
	return img
}

func getPartNumber(mid *monsterid, part string) int {
	switch part {
	case "legs":
		return mid.legs
	case "hair":
		return mid.hair
	case "arms":
		return mid.arms
	case "body":
		return mid.body
	case "eyes":
		return mid.eyes
	case "mouth":
		return mid.mouth
	}
	return 0
}
