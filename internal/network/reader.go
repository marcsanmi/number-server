package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Read(conn net.Conn) (string, error) {
	// Make a buffer to hold length of data
	lenBuffer := make([]byte, 4)
	_, err := conn.Read(lenBuffer)
	if err != nil {
		return "", err
	}
	lenData, err := fromBytes(lenBuffer)
	if err != nil {
		return "", err
	}
	// Buffer to hold incoming data
	dataBuffer := make([]byte, lenData)
	length := 0
	// Read all the data
	for length < int(lenData) {
		requestLength, err := conn.Read(dataBuffer[length:])
		length += requestLength
		if err == io.EOF {
			return "", fmt.Errorf("> EOF before processing all the data")
		}
		if err != nil {
			return "", fmt.Errorf("> Error reading: %s", err.Error())
		}
	}
	return string(dataBuffer), nil
}

func fromBytes(b []byte) (int32, error) {
	buffer := bytes.NewReader(b)
	var result int32
	err := binary.Read(buffer, binary.BigEndian, &result)
	return result, err
}
