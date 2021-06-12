package processing

import (
	"encoding/binary"
	"unsafe"
)

type MP4File struct {
	data         []byte
	nativeEndian binary.ByteOrder
}

//Method used to prepare the MP4File for later usage
//Since the box size is represented as uint32, we need to decode it later on
//and decoding integers from bytes depends on endianness of the system which needs
//to be found beforehand.
func initializeMP4File(data []byte) MP4File {
	//create 2 byte array
	buf := [2]byte{}
	//load it with uint16
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	var nativeEndian binary.ByteOrder

	//check significance of the bytes
	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.BigEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.LittleEndian
	default:
		panic("Could not determine native endianness.")
	}

	return MP4File{data, nativeEndian}
}

//method used to calculate the box size from the file
func (file MP4File) getBoxSize(currentPos int) uint32 {
	currentPositionSlice := file.data[currentPos : currentPos+4]

	size := file.nativeEndian.Uint32(currentPositionSlice)
	return size
}

//method used to get the box type from the file
func (file MP4File) getBoxType(currentPos int) BoxType {
	//since the box size represents the first 4 bytes, we need to adjust the offset
	//to properly capture the box type
	typeOffset := currentPos + 4

	byteType := file.data[typeOffset : typeOffset+4]

	//convert the type from byte array into string
	stringType := string(byteType)

	//compare the conversion result and find the appropriate box type
	if stringType == FTYP.Name() {
		return FTYP
	} else if stringType == MOOV.Name() {
		return MOOV
	} else {
		return OTHER
	}
}

//method used to get the data from the box
func (file MP4File) getMP4FileData(currentPos, size int) []byte {
	startPos := currentPos + 8
	finalPos := currentPos + size

	if currentPos + size > len(file.data) {
		panic("Corrupt file. Size of the last box overflows file size!")
	}

	return file.data[startPos:finalPos]
}

//Function used to extract box from file
func (file MP4File) extractBoxFromByteArray(currentPos int) Box {
	boxSize := file.getBoxSize(currentPos)
	boxType := file.getBoxType(currentPos)
	boxData := file.getMP4FileData(currentPos, int(boxSize))

	return Box{BoxSize: boxSize, BoxType: boxType, BoxData: boxData}
}

//Function used to get all boxes from the file
func GetFileBoxes(bytes []byte) []Box {
	file := initializeMP4File(bytes)

	boxesInFile := make([]Box, 0)

	fileSize := len(file.data)
	currentPos := 0

	for fileSize > currentPos {
		box := file.extractBoxFromByteArray(currentPos)
		currentPos += int(box.BoxSize)
		boxesInFile = append(boxesInFile, box)
	}

	return boxesInFile
}
