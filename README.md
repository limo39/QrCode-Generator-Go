# QrCode-Generator-Go

A simple command-line tool written in Go that generates QR codes from URLs and saves them with filenames based on the input URL.

## Features

- Converts any valid URL into a QR code
- Automatically generates clean, filesystem-safe filenames from URLs
- Supports all standard URL formats (http, https, with/without www)
- Customizable QR code size and error correction level
- Lightweight and fast

## Installation

1. Ensure you have [Go installed](https://golang.org/doc/install) (version 1.16+ recommended)
2. Clone this repository or download the source code
3. Install the dependencies:

```bash
go get github.com/skip2/go-qrcode
````
## Usage

```bash
go run qrGen.go
```
When prompted, enter the URL you want to encode. The QR code will be saved as a PNG file with a name derived from your URL.

## Example
```bash
$ go run qrgen.go
QR Code Generator
Enter the URL you want to encode: https://github.com/user/repo

Generating QR code for: https://github.com/user/repo
Saving as: github.com_user_repo.png
QR code generated successfully as 'github.com_user_repo.png'
```

## Filename Generation Logic

The tool automatically creates filenames by:
Removing http:// and https:// prefixes
Removing www. prefix
Removing query parameters
Replacing special characters with underscores
Limiting to 50 characters

### Examples:

https://www.google.com → google.com.png
http://example.com/path/to/page → example.com_path_to_page.png
https://github.com/user/repo?branch=main → github.com_user_repo.png

## Customization
You can modify the QR code generation by editing these parameters in the code:

```go
// In generateQRCode function:
qrcode.Medium // Error correction level (options: Low, Medium, High, Highest)
256           // Image size in pixels
```
## Requirements

Go 1.16 or higher

GitHub.com/skip2/go-qrcode package

## License

This project is open source and available under the GNU License.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements.


