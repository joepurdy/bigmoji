# Bigmoji

Bigmoji is a command-line tool that splits a PNG image into 16 equal parts, making it perfect for creating large emoji mosaics or splitting images into a 4x4 grid. The tool automatically handles image padding to ensure square output pieces.

## Features

- Splits PNG images into 16 equal parts (4x4 grid)
- Automatically pads images to make them square
- Preserves transparency
- Simple command-line interface

## Installation

### Using Homebrew

```bash
brew install joepurdy/tap/bigmoji
```

### From Source

```bash
git clone https://github.com/joepurdy/bigmoji.git
cd bigmoji
go build
```

### Using Go

```bash
go install github.com/joepurdy/bigmoji@latest
```

## Usage

```bash
bigmoji <input.png>
```

The tool will:
1. Create an `out` directory if it doesn't exist
2. Split the input image into 16 parts
3. Save the parts as numbered PNG files in the `out` directory

### Example

```bash
bigmoji myimage.png
```

This will create 16 files in the `out` directory:
- `out/bigmyimage_1.png`
- `out/bigmyimage_2.png`
- ...
- `out/bigmyimage_16.png`

## Requirements

- Go 1.19 or later
- PNG input files

## License

This project is open source and available under the MIT License.
