package processing

import (
	"encoding/binary"
	"errors"
	"fmt"
	"unsafe"
)

type MP4File struct {
	data         []byte
	nativeEndian binary.ByteOrder
}

//Method used to prepare the MP4File for later usage
func initializeMP4File(data []byte) MP4File {

	return MP4File{data, checkEndianness()}
}

//Method used to check the system endianness
//Since the box size is represented as uint32, we need to decode it later on
//and decoding integers from bytes depends on endianness of the system which needs
//to be found beforehand.
func checkEndianness() binary.ByteOrder {
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

	return nativeEndian
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
func (file MP4File) getMP4FileData(currentPos, size int) ([]byte, error) {
	startPos := currentPos + 8
	finalPos := currentPos + size

	if currentPos+size > len(file.data) {
		return make([]byte, 0), errors.New("Corrupt file. Size of the last box overflows file size!")
	}

	return file.data[startPos:finalPos], nil
}

//Function used to extract box from file
func (file MP4File) extractBoxFromByteArray(currentPos int) (Box, error) {
	boxSize := file.getBoxSize(currentPos)
	boxType := file.getBoxType(currentPos)
	boxData, err := file.getMP4FileData(currentPos, int(boxSize))

	if err != nil {
		return Box{}, err
	}
	return Box{BoxSize: boxSize, BoxType: boxType, BoxData: boxData}, nil
}

//Function used to get all boxes from the file
func GetFileBoxes(bytes []byte) ([]Box, error) {
	file := initializeMP4File(bytes)

	boxesInFile := make([]Box, 0)

	fileSize := len(file.data)
	currentPos := 0

	for fileSize > currentPos {
		box, err := file.extractBoxFromByteArray(currentPos)
		if err != nil {
			errMessage := fmt.Sprintf("Error when extracting the box in position %d.\r\nNested error: %s", currentPos, err.Error())
			return make([]Box, 0), errors.New(errMessage)
		}

		currentPos += int(box.BoxSize)
		boxesInFile = append(boxesInFile, box)
	}

	return boxesInFile, nil
}

func FindBox(boxes []Box, boxType BoxType) (Box, error) {
	for _, box := range boxes {
		if box.BoxType == boxType {
			return box, nil
		}
	}

	return Box{}, errors.New(fmt.Sprintf("Box with type %s Not found\r\n", boxType.Name()))
}

func ConvertBoxesToByteArray(initSegment []Box) []byte {
	byteData := make([]byte, 0)

	for _, box := range initSegment {

		//append size
		size := make([]byte, 4)
		checkEndianness().PutUint32(size, box.BoxSize)
		byteData = append(byteData, size...)

		//append type
		byteData = append(byteData, []byte(box.BoxType.Name())...)

		//append box data
		byteData = append(byteData, box.BoxData...)
	}

	return byteData
}

//function used to unwrap all the boxes,
//get the ftyp and moov boxes
//generate init segment
//and to transform the init segment to byte array
func GetInitSegmentAsBytes(bytes []byte) ([]byte, error) {
	boxes, err := GetFileBoxes(bytes)

	if err != nil {
		return make([]byte, 0), err
	}

	//get ftyp box
	ftypBox, err := FindBox(boxes, FTYP)
	if err != nil {
		return make([]byte, 0), err
	}

	//get moov box
	moovBox, err := FindBox(boxes, MOOV)
	if err != nil {
		return make([]byte, 0), err
	}

	return ConvertBoxesToByteArray([]Box{ftypBox, moovBox}), nil
}
