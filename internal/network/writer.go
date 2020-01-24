package network

import (
	"bytes"
	"encoding/binary"
	"net"
)

func toBytes(i int32) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, i)
	return buffer.Bytes(), err
}

func Write(conn net.Conn, message string) error {
	// Write the size of the message to be sent
	data, err := toBytes(int32(len([]byte(message))))
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	// Send the message
	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}
