package main

import "net/http"

var (
	defaultSniffers = map[string]func([]byte) bool{
		"image/jp2":                 imageJP2,
		"image/tiff":                imageTIFF,
		"image/vnd.adobe.photoshop": imageVNDAdobePhotoshop,
		"image/x-canon-cr2":         imageXCanonCR2,
	}
)

func Sniff(b []byte) string {
	if len(b) == 0 {
		return "application/octet-stream"
	}

	for mt, s := range defaultSniffers {
		if s(b) {
			return mt
		}
	}

	return http.DetectContentType(b)
}

// imageJP2 reports whether the b's MIME type is "image/jp2".
func imageJP2(b []byte) bool {
	return len(b) > 12 &&
		b[0] == 0x0 &&
		b[1] == 0x0 &&
		b[2] == 0x0 &&
		b[3] == 0xc &&
		b[4] == 0x6a &&
		b[5] == 0x50 &&
		b[6] == 0x20 &&
		b[7] == 0x20 &&
		b[8] == 0xd &&
		b[9] == 0xa &&
		b[10] == 0x87 &&
		b[11] == 0xa &&
		b[12] == 0x0
}

// imageTIFF reports whether the b's MIME type is "image/tiff".
func imageTIFF(b []byte) bool {
	return len(b) > 3 &&
		(b[0] == 0x49 &&
			b[1] == 0x49 &&
			b[2] == 0x2a &&
			b[3] == 0x0 ||
			b[0] == 0x4d &&
				b[1] == 0x4d &&
				b[2] == 0x0 &&
				b[3] == 0x2a)
}

// imageVNDAdobePhotoshop reports whether the b's MIME type is
// "image/vnd.adobe.photoshop".
func imageVNDAdobePhotoshop(b []byte) bool {
	return len(b) > 3 &&
		b[0] == 0x38 &&
		b[1] == 0x42 &&
		b[2] == 0x50 &&
		b[3] == 0x53
}

// imageXCanonCR2 reports whether the b's MIME type is "image/x-canon-cr2".
func imageXCanonCR2(b []byte) bool {
	return len(b) > 9 &&
		(b[0] == 0x49 &&
			b[1] == 0x49 &&
			b[2] == 0x2a &&
			b[3] == 0x0 ||
			b[0] == 0x4d &&
				b[1] == 0x4d &&
				b[2] == 0x0 &&
				b[3] == 0x2a) &&
		b[8] == 0x43 && b[9] == 0x52
}
