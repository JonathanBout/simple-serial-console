# Simple Serial Console
_replace `ssc` with the name of your binary, e.g. `ssc-windows-amd64`_

## Usage
```
ssc <port> [<baud rate>]
```
### `<port>`
#### Windows
A COM Port `COMX` where X is a number

#### Linux
A USB TTY, usually something like `/dev/ttyUSBX` where X is a number

### `[<baud rate>]` (optional)
The baud rate to use for communication. Default is 115200.