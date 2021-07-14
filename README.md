[![Release](https://img.shields.io/github/v/release/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go/releases/latest)
[![Release Date](https://img.shields.io/github/release-date/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/jolsfd/imagenamer-go)](https://goreportcard.com/report/github.com/jolsfd/imagenamer-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go)
[![License](https://img.shields.io/github/license/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go/blob/main/LICENSE)
[![Issues](https://img.shields.io/github/issues/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go/issues/)
[![Pull-requests](https://img.shields.io/github/issues-pr/jolsfd/imagenamer-go.svg)](https://github.com/jolsfd/imagenamer-go/pulls)
[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/jolsfd/imagenamer-go)

# ImageNamer-Go

## Installation

ImageNamer-Go is written in Go, so you can install it through `go install`:

```bash
go install github.com/jolsfd/imagenamer-go
```

Pre-compiled binaries for Linux and Windows on the [releases page](https://github.com/jolsfd/imagenamer-go/releases/latest). After you download the file archive for your operating system and architecture extract the file archive:

**For Linux:**

```bash
# make binary executable
$ chmod +x imagenamer-go

# move binary to path
$ sudo mv imagenamer-go /usr/local/bin
```

**For Windows:**

```powershell
# move binary to a directory
$ move imagenamer-go.exe C:\Users\<user>\.imagenamer-go
```

## Disclaimer

ImageNamer-Go command line interface is currently a work in progress. Please keep this in mind when using.
Make sure you use ImagerNamer-Go CLI with caution and at your own risk. **We do not raise any liability or guarantee.**

## Documentation

Visit the documentation site [here](https://github.com/jolsfd/imagenamer-go/blob/main/docs/README.md). Please note that the documentation is still a work in progress.

## Credits

ImageNamer-Go relies on other open source software listed below:
* olekukonko/tablewriter
* rwcarlsen/goexif
* fatih/color
* spf13/cobra
* spf13/viper
* ayoisaiah/f2

## License

Released under the terms of the [MIT License](https://github.com/jolsfd/imagenamer-go/blob/main/LICENSE).
