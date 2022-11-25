package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

// 传输包体
type NormalPacker struct {
	// 数据存入，开头大端/小端
	Order binary.ByteOrder
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	normalPacker := &NormalPacker{
		Order: order,
	}
	return normalPacker
}

// 打包 |data 长度|id|data|
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(buffer)))
	p.Order.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

// 解包
func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second))

	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)

	if err != nil {
		return nil, err
	}

	totalLen := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:])

	dataLen := totalLen - 16

	dataBuffer := make([]byte, dataLen)

	_, err = io.ReadFull(reader, dataBuffer)

	if err != nil {
		return nil, err
	}

	msg := &Message{
		Id:   id,
		Data: dataBuffer,
	}

	return msg, nil
}
