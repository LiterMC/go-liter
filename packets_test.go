package liter_test

import (
	"bytes"
	"testing"

	. "github.com/kmcsr/go-liter"
)

func TestPacketForward(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	var (
		pb   *PacketBuilder = NewPacket(V1_19_2, 987)
		pr   *PacketReader
		buf  = bytes.NewBuffer(nil)
		buf2 = bytes.NewBuffer(nil)
		err  error
	)
	pb.ByteArray(data)
	pb.WriteTo(buf)
	if pr, err = ReadPacket(V1_19_2, bytes.NewReader(buf.Bytes()), false); err != nil {
		t.Fatalf("Error when read packet: %v", err)
	}
	if pr.Protocol() != pb.Protocol() {
		t.Fatalf("Protocol not match")
	}
	if pr.Id() != pb.Id() {
		t.Fatalf("Id not match")
	}
	if !bytes.Equal(pr.Bytes(), data) {
		t.Fatalf("Data not match, %v != %v", pr.Bytes(), data)
	}
	pb.Reset(pr.Protocol(), pr.Id()).ByteArray(pr.Bytes())
	pb.WriteTo(buf2)
	if !bytes.Equal(buf.Bytes(), buf2.Bytes()) {
		t.Fatalf("Second bytes not match")
	}
}

func TestPacketForwardNoCompressed(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	var (
		pb   *PacketBuilder = NewPacket(V1_19_2, 987)
		pr   *PacketReader
		buf  = bytes.NewBuffer(nil)
		buf2 = bytes.NewBuffer(nil)
		err  error
	)
	pb.ByteArray(data)
	pb.WriteCompressedTo(buf, 10000)
	if pr, err = ReadPacket(V1_19_2, bytes.NewReader(buf.Bytes()), true); err != nil {
		t.Fatalf("Error when read packet: %v", err)
	}
	if pr.Protocol() != pb.Protocol() {
		t.Fatalf("Protocol not match")
	}
	if pr.Id() != pb.Id() {
		t.Fatalf("Id not match")
	}
	if !bytes.Equal(pr.Bytes(), data) {
		t.Fatalf("Data not match, %v != %v", pr.Bytes(), data)
	}
	pb.Reset(pr.Protocol(), pr.Id()).ByteArray(pr.Bytes())
	pb.WriteCompressedTo(buf2, 10000)
	if !bytes.Equal(buf.Bytes(), buf2.Bytes()) {
		t.Fatalf("Second bytes not match")
	}
}

func TestPacketForwardCompressed(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	var (
		pb   *PacketBuilder = NewPacket(V1_19_2, 987)
		pr   *PacketReader
		buf  = bytes.NewBuffer(nil)
		buf2 = bytes.NewBuffer(nil)
		err  error
	)
	pb.ByteArray(data)
	pb.WriteCompressedTo(buf, 0)
	if pr, err = ReadPacket(V1_19_2, bytes.NewReader(buf.Bytes()), true); err != nil {
		t.Fatalf("Error when read packet: %v", err)
	}
	if pr.Protocol() != pb.Protocol() {
		t.Fatalf("Protocol not match")
	}
	if pr.Id() != pb.Id() {
		t.Fatalf("Id not match")
	}
	if !bytes.Equal(pr.Bytes(), data) {
		t.Fatalf("Data not match, %v != %v", pr.Bytes(), data)
	}
	pb.Reset(pr.Protocol(), pr.Id()).ByteArray(pr.Bytes())
	pb.WriteCompressedTo(buf2, 0)
	if !bytes.Equal(buf.Bytes(), buf2.Bytes()) {
		t.Fatalf("Second bytes not match")
	}
}
