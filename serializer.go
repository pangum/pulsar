package pulsar

import (
	`encoding/json`
	`encoding/xml`

	`github.com/vmihailenco/msgpack/v5`
	`google.golang.org/protobuf/proto`
)

const (
	serializerUnknown serializer = ``
	serializerJson    serializer = `json`
	serializerProto   serializer = `proto`
	serializerMsgpack serializer = `msgpack`
	serializerXml     serializer = `xml`
	serializerString  serializer = `string`
	serializerBytes   serializer = `bytes`
)

type serializer string

func (s serializer) Marshal(from interface{}) (to []byte, err error) {
	switch s {
	case serializerProto:
		to, err = proto.Marshal(from.(proto.Message))
	case serializerJson:
		to, err = json.Marshal(from)
	case serializerXml:
		to, err = xml.Marshal(from)
	case serializerMsgpack:
		to, err = msgpack.Marshal(from)
	case serializerBytes:
		to = from.([]byte)
	case serializerString:
		to = []byte(from.(string))
	}

	return
}
