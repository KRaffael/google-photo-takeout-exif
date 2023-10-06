# go-google-photo-takeout-exif-cli
[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-exif-update.svg)](https://pkg.go.dev/github.com/sfomuseum/go-exif-update)

Fix missing exif metadata (photo taken daten) in Takeouts from Google Photo with JSON metadata.

Using [go-exif-update](https://github.com/sfomuseum/go-exif-update).

## how to use

```
$> ./bin/takeout-cli -h

2023/10/06 07:45:03 Welcome to google-photo-takeout-exif-cli...
Usage of C:\develop\github\google-photo-takeout-exif\main.exe:
      --source-directory string   path to the extracted directory of Google takeout container
      --target-directory string   path to the target directory where the edited photos has to be saved to
pflag: help requested
```

```
%> ./bin/takeout-cli --source-directory=<C:\foo\bar> --target-directory=<C:\foobar>
```

## development

### initial setup

Clone this outside your GOPATH

### build

`go build -o takeout-cli`

### tests

`go test -coverpkg=./internal/... -v ./...`

## license

see license file [here](LICENSE)
