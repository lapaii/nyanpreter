package serialiser

import (
	"bytes"
	"encoding/binary"
	"shared"
)

func Serialise(input []shared.Instruction) (*bytes.Buffer, error) {
	// this is whats going to output a buffer with the binary format

	buf := new(bytes.Buffer)

	_, err := buf.WriteString("nya:3")

	if err != nil {
		return buf, err
	}

	for _, instruction := range input {
		header := uint16(0) // 2 bytes, set

		header = uint16(instruction.SourceType) | header // insert source type

		targetType := instruction.DestType << 2 // is now like 0000 xx 00

		header = uint16(targetType) | header // insert target type

		opcodeShifted := uint16(instruction.Opcode) << 4

		header = uint16(opcodeShifted) | header

		// write header bytes to buf
		err := binary.Write(buf, binary.BigEndian, header)

		if err != nil {
			return buf, err
		}

		// dest 1 byte or 4 bytes next
		// register or register pointer = 1 byte for target operand
		if instruction.DestType == shared.Register || instruction.DestType == shared.RegisterPointer {
			// write register num to buffer

			err := binary.Write(buf, binary.BigEndian, uint8(instruction.Dest))

			if err != nil {
				return buf, err
			}
		} else {
			// the destination type must now be immediate or immediate pointer so i need to write 4 bytes

			err := binary.Write(buf, binary.BigEndian, uint32(instruction.Dest))

			if err != nil {
				return buf, err
			}
		}

		// source 1 byte or 4 bytes
		if instruction.SourceType == shared.Register || instruction.SourceType == shared.RegisterPointer {
			err := binary.Write(buf, binary.BigEndian, uint8(instruction.Source))

			if err != nil {
				return buf, err
			}
		} else {
			// the destination type must now be immediate or immediate pointer so i need to write 4 bytes

			err := binary.Write(buf, binary.BigEndian, uint32(instruction.Source))

			if err != nil {
				return buf, err
			}
		}
	}

	return buf, nil
}
