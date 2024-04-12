package base64util

import (
	"bytes"
	"image"
	"image/png"
)

// ImgToBase64 将图像转换为base64字符串
// img: 要转换的图像
// return: base64字符串
func ImgToBase64(img image.Image) (string, error) {
	// Create a buffer to store the encoded image
	buffer := new(bytes.Buffer)
	// Encode the image to the buffer
	err := png.Encode(buffer, img)
	if err != nil {
		return "", err
	}

	// Convert the buffer to a base64 string
	return Base64Encode(buffer.String()), nil
}

// Base64ToImg 将base64字符串转换为图像
// base64Str: 要转换的base64字符串
// return: 图像
func Base64ToImg(base64Str string) (image.Image, error) {
	// Decode the base64 string
	decodedStr := Base64Decode(base64Str)
	// Create a buffer to store the decoded image
	buffer := bytes.NewBufferString(decodedStr)
	// Decode the image from the buffer
	img, _, err := image.Decode(buffer)
	if err != nil {
		return nil, err
	}

	return img, nil
}
