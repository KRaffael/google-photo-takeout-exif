package main

import (
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/params"
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/service"
	"log"
	"os"
)

func main() {
	log.Println("Welcome to google-photo-takeout-exif-cli...")

	parameterService := handleParameters()

	log.Println("Parameters are valid, proceeding...")

	googlePhotoTakeoutExifCliService := service.New(parameterService)
	returnCode := googlePhotoTakeoutExifCliService.Run()

	log.Println("DONE")

	os.Exit(returnCode)
}

func handleParameters() params.ParameterService {
	parameterService := params.New()
	parameterService.DeclareFlags()
	parameterService.Obtain()
	err := parameterService.Validate()
	if err != nil {
		log.Println("PARAMETER ERROR: parameters are invalid. See log output above. Bailing out")
		os.Exit(1)
	}
	return parameterService
}
