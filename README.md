# QR Code Generator and Decoder

This is a simple command-line tool written in Go for generating and decoding QR codes using the skip2qrcode and tuotooqrcode packages.

## Usage

### Generate QR Code
To generate a QR code, use the following command:
```bash
go run main.go -mode=generate -input="YOUR_TEXT_HERE" -output="output.png"
```

Replace "YOUR_TEXT_HERE" with the text for which you want to generate the QR code. The output will be saved in the file specified by the -output flag (default: qrcode.png).

### Decode QR Code
To decode a QR code, use the following command:


```bash
go run main.go -mode=decode -input="input.png"
```

Replace "input.png" with the filename of the QR code image you want to decode.

### Building from Source
Make sure you have Go installed on your machine.

#### 1. Clone the repository:

```bash
git clone https://github.com/dhawal-pandya/qrcode-generator-decoder.git
```
#### 2. Navigate to the project directory:

```bash
cd qrcode-generator-decoder
```

#### 3. Build and run the project:

```bash
go build
./qrcode-generator-decoder -mode=generate -input="YOUR_TEXT_HERE" -output=output.png
```


### Dependencies

[skip2qrcode](https://github.com/skip2/go-qrcode) for generating QR codes \
[tuotooqrcode](https://github.com/tuotoo/qrcode) for decoding QR codes

Feel free to contribute or report issues!




