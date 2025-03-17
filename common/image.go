// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/ygpkg/gooxml/measurement"

	// Add image format support
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// Image is a container for image information. It's used as we need format and
// and size information to use images.
// It contains either the filesystem path to the image, or the image itself.
type Image struct {
	Size   image.Point
	Format string
	Path   string
	Data   *[]byte
}

// ImageRef is a reference to an image within a document.
type ImageRef struct {
	d     *DocBase
	rels  Relationships
	img   Image
	relID string
}

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef(img Image, d *DocBase, rels Relationships) ImageRef {
	return ImageRef{img: img, d: d, rels: rels}
}

func (i *ImageRef) SetRelID(id string) {
	i.relID = id
}

// RelID returns the relationship ID.
func (i ImageRef) RelID() string {
	return i.relID
}

// Format returns the format of the underlying image
func (i ImageRef) Format() string {
	return i.img.Format
}

// Path returns the path to an image file, if any.
func (i ImageRef) Path() string {
	return i.img.Path
}

// Data returns the data of an image file, if any.
func (i ImageRef) Data() *[]byte {
	return i.img.Data
}

// Size returns the size of an image
func (i ImageRef) Size() image.Point {
	return i.img.Size
}

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (i ImageRef) RelativeHeight(w measurement.Distance) measurement.Distance {
	hScale := float64(i.Size().Y) / float64(i.Size().X)
	return w * measurement.Distance(hScale)
}

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (i ImageRef) RelativeWidth(h measurement.Distance) measurement.Distance {
	wScale := float64(i.Size().X) / float64(i.Size().Y)
	return h * measurement.Distance(wScale)
}

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size.  You can also
// construct an Image directly if the file and size are known.
func ImageFromFile(path string) (Image, error) {
	f, err := os.Open(path)
	r := Image{}
	if err != nil {
		return r, fmt.Errorf("error reading image: %s", err)
	}
	defer f.Close()
	imgDec, ifmt, err := image.Decode(f)
	if err != nil {
		return r, fmt.Errorf("unable to parse image: %s", err)
	}

	r.Path = path
	r.Format = ifmt
	r.Size = imgDec.Bounds().Size()
	return r, nil
}

// ImageFromBytes returns an Image struct for an in-memory image. You can also
// construct an Image directly if the file and size are known.
func ImageFromBytes(data []byte) (Image, error) {
	r := Image{}
	imgDec, ifmt, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return r, fmt.Errorf("unable to parse image: %s", err)
	}

	r.Data = &data
	r.Format = ifmt
	r.Size = imgDec.Bounds().Size()
	return r, nil
}

// ImageFromURL reads an image from a URL. It downloads the image, determines
// its format and size, and returns an Image struct.
func ImageFromURL(url string) (Image, error) {
	r := Image{}

	// Download the image from the URL
	resp, err := http.Get(url)
	if err != nil {
		return r, fmt.Errorf("error downloading image: %s", err)
	}
	defer resp.Body.Close()

	// Read the image data into a byte slice
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, fmt.Errorf("error reading image data: %s", err)
	}

	// Decode the image to determine its format and size
	imgDec, ifmt, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return r, fmt.Errorf("unable to parse image: %s", err)
	}
	// Encode the image back to a byte slice with optimized compression
	var buf bytes.Buffer
	switch ifmt {
	case "jpeg":
		// Use the standard JPEG encoder with specified quality
		err = jpeg.Encode(&buf, imgDec, &jpeg.Options{Quality: 50})
	case "png":
		// Use the standard PNG encoder with default compression
		err = png.Encode(&buf, imgDec)
	default:
		return r, fmt.Errorf("unsupported image format: %s", ifmt)
	}
	if err != nil {
		return r, fmt.Errorf("error encoding image: %s", err)
	}
	// Update the image data
	compressedData := buf.Bytes()
	r.Data = &compressedData
	r.Format = ifmt
	r.Size = imgDec.Bounds().Size()
	return r, nil
}
