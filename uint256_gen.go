package uint256

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Int) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != uint32(4) {
		err = msgp.ArrayError{Wanted: uint32(4), Got: zb0001}
		return
	}
	for za0001 := range z {
		z[za0001], err = dc.ReadUint64()
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Int) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(4))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for za0001 := range z {
		err = en.WriteUint64(z[za0001])
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Int) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(4))
	for za0001 := range z {
		o = msgp.AppendUint64(o, z[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Int) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != uint32(4) {
		err = msgp.ArrayError{Wanted: uint32(4), Got: zb0001}
		return
	}
	for za0001 := range z {
		z[za0001], bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Int) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (4 * (msgp.Uint64Size))
	return
}
