package resp

import "strconv"

// Write RESP
func (w *Writer) Write(v Value) error {
	var bytes = v.Marshal()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (v Value) Marshal() []byte {
	switch v.Typ {
		case "array":
			return v.marshalArray()
		case "string":
			return v.marshalString()
		case "bulk":
			return v.marshalBulk()
		case "null":
			return v.marshallNull()
		case "error":
			return v.marshalError()
		default:
			return []byte{}
	}
}

func (v Value) marshalString() []byte {
	var  bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalArray() []byte {
	var  bytes []byte
	len := len(v.Array)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}
	return bytes
}

func (v Value) marshalBulk() []byte {
	var  bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")

}

func (v Value) marshalError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}


