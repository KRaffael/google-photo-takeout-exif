package service

import "github.com/KRaffael/google-photo-takeout-exif-cli/internal/params"

type GooglePhotoTakeoutExifCliService interface {
	Run() int
}

func New(parameters params.ParameterService) GooglePhotoTakeoutExifCliService {
	return &GooglePhotoTakeoutExifCli{Parameters: parameters}
}
