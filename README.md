# Simple Serial Console (SSC)
SSC is exactly what you think it is: A serial console without bloat.
For simple communication you don't need anything else! You can use it to control most 3D-printers, an Arduino and loads more.

You only need to know two things:
1. The baud rate the device expects, which you can usually define manually or find somewhere in the info screen
2. The line endings the device expects. It's not always documented, but usually starting with the default (`LF`) works.


## Usage
_replace `ssc` with the name of your binary, e.g. `ssc-windows-amd64` or `ssc-linux-arm64`, or rename your binary to `ssc`_
```
ssc <port> [<baud rate>] [<newline>]
```
### `<port>`
#### Windows
A COM Port `COMX` where X is a number

#### Linux
A USB TTY, usually something like `/dev/ttyUSBX` where X is a number

### `[<baud rate>]` (optional)
The baud rate to use for communication. Default is 115200.

### `[<newline>]` (optional)
The newline character(s) to use while communicating. Default is LF (\n).
Allowed values are:
- `CR` (carriage return `\r`)
- `LF` (line feed `\n`)
- `CRLF` (carriage return + line feed `\r\n`)
- `LFCR` (line feed + carriage return `\n\r`)

Usually, this is one of `LF` and `CRLF`.
