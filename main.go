package main

import (
	"bufio"
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
			fmt.Println(err)
			criticalError("Invalid baud rate")
		}

		if i < 0 {
			criticalError("Invalid baud rate")
		}

		inputBaudRate = i
	}

	begin(inputPort, inputBaudRate)
}

// Opens a new connection on the specified port at the specified baud rate
func begin(portName string, baudRate int) {
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
	go userInput(port)

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
func userInput(port *serial.Port) {
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
		// after appending a newline character to it
		_, err = port.Write([]byte(input + "\n"))

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
