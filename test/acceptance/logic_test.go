package acceptance

import (
	"fmt"
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/params"
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/service"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func Test_parameter_validation(t *testing.T) {
	fmt.Println("    Given empty parameters")
	parameterService := &params.Parameters{
		SourceDirectory: "",
		TargetDirectory: "",
	}

	fmt.Println("    When parameter validation executed")
	err := parameterService.Validate()
	if err == nil {
		require.Fail(t, "test failed")
		return
	}

	fmt.Println("    Then an validation error occurred")
	require.Equal(t, "there were parameter errors", err.Error())
}

func Test_run(t *testing.T) {
	fmt.Println("    Given valid parameters")
	targetPath := filepath.Join("..", "resources", "output")

	parameterService := &params.Parameters{
		SourceDirectory: filepath.Join("..", "resources", "takeout"),
		TargetDirectory: targetPath,
	}
	err := parameterService.Validate()
	if err != nil {
		require.Fail(t, "test failed")
		return
	}

	fmt.Println("    When exif logic executed")
	googlePhotoTakeoutExifCliService := service.New(parameterService)
	returnCode := googlePhotoTakeoutExifCliService.Run()

	fmt.Println("    Then a valid result exists")
	require.Equal(t, 0, returnCode)
	checkFileExistsWithExifTags(t, targetPath, "album-one")
	checkFileExistsWithExifTags(t, targetPath, "")
	checkFileNotExists(t, targetPath, "album-one", "IMG_1078-no-exif.JPG.json")
	checkFileNotExists(t, targetPath, "album-two", "IMG_1078-no-exif.JPG")
	checkFileNotExists(t, targetPath, "album-two", "IMG_1078-no-exif.JPG.json")
	checkFileNotExists(t, targetPath, "", "IMG_1078-no-exif.JPG.json")

	cleanUp(t, targetPath)
}

// --- checks

func checkFileNotExists(t *testing.T, path string, album string, fileName string) {
	_, err := os.Open(filepath.Join(path, album, fileName))
	if err == nil {
		require.Fail(t, "test failed")
		return
	}
}

func checkFileExistsWithExifTags(t *testing.T, path string, album string) {
	file, _ := os.Open(filepath.Join(path, album, "IMG_1078-no-exif.JPG"))
	defer file.Close()
	x, _ := exif.Decode(file)
	dateTimeOriginal, _ := x.Get("DateTimeOriginal")
	require.Equal(t, "\"2014:01:03 14:20:53\"", dateTimeOriginal.String())
	dateTimeDigitized, _ := x.Get("DateTimeDigitized")
	require.Equal(t, "\"2014:01:03 14:20:53\"", dateTimeDigitized.String())
}

// --- helper

func cleanUp(t *testing.T, path string) {
	err := os.RemoveAll(path)
	if err != nil {
		require.Fail(t, "test failed")
	}
}
