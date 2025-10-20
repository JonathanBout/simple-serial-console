package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

func main() {
	// args[0] is the exe path, which we don't need
	args := os.Args[1:]

	inputPort := ""
	inputBaudRate := 115200
	inputNewline := "\n"

	// the first argument has to be the USB/Serial Device
	if len(args) > 0 {
		// find args[0] in serialPorts
		inputPort = args[0]
	} else {
		criticalError("Provide a port name")
	}

	// Optionally, the second argument can be the baud rate.
	// Default rate is 115200
	if len(args) > 1 {
		i, err := strconv.Atoi(args[1])

		if err != nil {
			// if the 2nd parameter is not a valid baudrate,
			// it could be the newline as baudrate is optional.
			newline, newline_err := parseNewline(args[1])

			// if not a valid newline, show the baudrate error
			if newline_err != nil {
				criticalError("Invalid baud rate")
			}

			inputNewline = newline
		} else if i < 0 {
			criticalError("Invalid baud rate")
		} else {
			inputBaudRate = i

			// if it was a valid baudrate and there is a third argument,
			// try to parse it as a newline.
			if len(args) > 2 {
				newline, err := parseNewline(args[2])

				if err != nil {
					fmt.Println(err)
					criticalError("Invalid newline character")
				}

				inputNewline = newline
			}
		}

	}

	begin(inputPort, inputBaudRate, inputNewline)
}

func parseNewline(arg string) (string, error) {
	switch arg {
	case "CR": // carriage-return
		return "\r", nil
	case "LF": // line-feed
		return "\n", nil
	case "CRLF": // carriage-return line-feed
		return "\r\n", nil
	case "LFCR": // line-feed carriage-return
		return "\n\r", nil
	default:
		return "", errors.New("invalid newline")
	}
}

// Opens a new connection on the specified port at the specified baud rate
func begin(portName string, baudRate int, newline string) {
	config := &serial.Config{
		Baud: baudRate,
		Name: portName,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		criticalError("Error opening serial port: " + err.Error())
	}

	// close the port when this method exits
	defer func(port *serial.Port) {
		err := port.Close()
		if err != nil {
			criticalError("Error closing serial port: " + err.Error() + ". The connection might stay open.")
		}
	}(port)

	// asynchronously read input
	go userInput(port, newline)

	// while also receiving data from the serial device
	buf := make([]byte, 1024)

	for {
		n, err := port.Read(buf)

		if err != nil {
			criticalError("Error reading from serial port: " + err.Error())
		}

		fmt.Printf("%s", string(buf[:n]))
	}
}

// reads the standard input until the next EOF,
// or until the input line equals "exit"
func userInput(port *serial.Port, newline string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}

			criticalError("Error reading input: " + err.Error())
		}

		if len(input) == 0 {
			continue
		}

		input = strings.TrimSpace(input)

		// write the input data to the port,
		// after appending the defined newline character to it
		_, err = port.Write([]byte(input + newline))

		if err != nil {
			criticalError("Error writing to serial port: " + err.Error())
		}
	}
}

// helper method to print a critical error and exit with code 1
func criticalError(message string) {
	fmt.Println(message)
	os.Exit(1)
}
