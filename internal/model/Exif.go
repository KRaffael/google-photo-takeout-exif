package model

type Exif struct {
	PhotoTakenTime PhotoTakenTime `json:"photoTakenTime"`
	GeoDataExif    GeoDataExif    `json:"geoDataExif"`
}

type PhotoTakenTime struct {
	Timestamp string `json:"timestamp"`
}

type GeoDataExif struct {
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	Altitude      float32 `json:"altitude"`
	LatitudeSpan  float32 `json:"latitudeSpan"`
	LongitudeSpan float32 `json:"longitudeSpan"`
}
