# Simple Serial Console
_replace `ssc` with the name of your binary, e.g. `ssc-windows-amd64` or `ssc-linux-arm64`_

## Usage
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
