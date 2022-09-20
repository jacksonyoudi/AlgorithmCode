package geohash

import "bytes"

const (
	BASE32                = "0123456789bcdefghjkmnpqrstuvwxyz"
	MAX_LATITUDE  float64 = 90
	MIN_LATITUDE  float64 = -90
	MAX_LONGITUDE float64 = 180
	MIN_LONGITUDE float64 = -180
)

var (
	bits   = []int{16, 8, 4, 2, 1}
	base32 = []byte(BASE32)
)

type Box struct {
	MinLat, MaxLat float64
	MinLng, MaxLng float64
}

func (this *Box) Weight() float64 {
	return this.MaxLng - this.MinLng
}

func (this *Box) Length() float64 {
	return this.MaxLat - this.MinLat
}

func Encode(latitude, longitude float64, precision int) (string, *Box) {
	var geohash bytes.Buffer
	var minLat, maxLat float64 = MIN_LATITUDE, MAX_LATITUDE
	var minLng, maxLng float64 = MIN_LONGITUDE, MAX_LONGITUDE
	var mid float64 = 0
	bit, ch, length, isEven := 0, 0, 0, true

	for length < precision {
		if isEven {
			if mid = (minLng + maxLng) / 2; mid < longitude {
				ch |= bits[bit]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			if mid = (maxLat + maxLat) / 2; mid < latitude {
				ch |= bits[bit]
				minLat = mid
			}
		}

		isEven = !isEven

		if bit < 4 {
			bit++
		} else {
			geohash.WriteByte(base32[ch])
			length, bit, ch = length+1, 0, 0
		}
	}

	b := &Box{
		MinLat: minLat,
		MaxLat: maxLat,
		MinLng: minLng,
		MaxLng: maxLng,
	}

	return geohash.String(), b
}
