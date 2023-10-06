package service

import (
	"encoding/json"
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/model"
	"github.com/KRaffael/google-photo-takeout-exif-cli/internal/params"
	exifupdate "github.com/sfomuseum/go-exif-update"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var jpgFileRegex = regexp.MustCompile("\\.(jpg|JPG|jpeg|JPEG|png|PNG)$")

type GooglePhotoTakeoutExifCli struct {
	Parameters params.ParameterService
}

func (n *GooglePhotoTakeoutExifCli) Run() int {
	return n.workingOnDirectory(n.Parameters.GetSourceDirectory(), n.Parameters.GetTargetDirectory())
}

func (n *GooglePhotoTakeoutExifCli) workingOnDirectory(sourceDirectory string, targetDirectory string) int {
	sourceDirEntries := n.readDirectory(sourceDirectory)

	sourcePath := n.getAbs(sourceDirectory)
	targetPath := n.getAbs(targetDirectory)

	n.createFolderIfNotExists(targetPath)

	n.workingOnDirEntries(sourceDirEntries, sourcePath, targetPath)

	return 0
}

func (n *GooglePhotoTakeoutExifCli) readDirectory(sourceDirectory string) []os.DirEntry {
	sourceDirEntries, err := os.ReadDir(sourceDirectory)
	if err != nil {
		log.Printf("ERROR: folder can not be read/found %s\n", err.Error())
		os.Exit(1)
	}
	return sourceDirEntries
}

func (n *GooglePhotoTakeoutExifCli) getAbs(sourceDirectory string) string {
	sourcePath, err := filepath.Abs(sourceDirectory)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	return sourcePath
}

func (n *GooglePhotoTakeoutExifCli) createFolderIfNotExists(targetPath string) {
	_, err := os.Stat(targetPath)
	if err != nil {
		err = os.Mkdir(targetPath, os.ModePerm)
		if err != nil {
			log.Printf("ERROR: directory can not be created %s", targetPath)
			os.Exit(1)
		}
	}
}

func (n *GooglePhotoTakeoutExifCli) workingOnDirEntries(sourceDirEntries []os.DirEntry, sourcePath string, targetPath string) {
	for _, sourceDirEntry := range sourceDirEntries {
		if sourceDirEntry.IsDir() {
			n.workingOnDirectory(filepath.Join(sourcePath, sourceDirEntry.Name()), filepath.Join(targetPath, sourceDirEntry.Name()))
		} else {
			sourceDirEntry.Type()
			if jpgFileRegex.MatchString(sourceDirEntry.Name()) {
				sourceFilePath := filepath.Join(sourcePath, sourceDirEntry.Name())
				targetFilePath := filepath.Join(targetPath, sourceDirEntry.Name())

				sourceJsonFile, err := os.ReadFile(sourceFilePath + ".json")
				if err == nil {

					log.Printf("processing file %s target %s\n", sourceFilePath, targetFilePath)

					var exifDto model.Exif
					timestamp := n.getPhotoTakenTimestamp(sourceJsonFile, exifDto)
					_ = os.Chtimes(sourceFilePath, timestamp, timestamp)

					exifProps := map[string]interface{}{
						"DateTimeOriginal":  timestamp.UTC(),
						"DateTimeDigitized": timestamp.UTC(),
						//"GPSLatitude":       37.61799,
						//"GPSLongitude":      -122.384864,
					}

					fh, _ := os.Open(sourceFilePath)
					defer fh.Close()

					newFh, _ := os.Create(targetFilePath)
					defer newFh.Close()

					err = exifupdate.UpdateExif(fh, newFh, exifProps)
					if err != nil {
						log.Printf("SKIP: update exif data failed %s\n", err.Error())
					}

					log.Printf("done\n")
				} else {
					log.Printf("SKIP: corresponding json file for %s can not be read/found\n", sourceFilePath)
				}
			} else {
				log.Printf("SKIP: %s is not a image (jpg|jpeg|png)", sourceDirEntry.Name())
			}
		}
	}
}

func (n *GooglePhotoTakeoutExifCli) getPhotoTakenTimestamp(jsonFile []byte, exifDto model.Exif) time.Time {
	_ = json.Unmarshal(jsonFile, &exifDto)
	t, _ := strconv.ParseInt(exifDto.PhotoTakenTime.Timestamp, 10, 64)
	return time.Unix(t, 0)
}
