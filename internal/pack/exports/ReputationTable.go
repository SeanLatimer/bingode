// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReputationTable struct {
	_tab flatbuffers.Table
}

func GetRootAsReputationTable(buf []byte, offset flatbuffers.UOffsetT) *ReputationTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ReputationTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ReputationTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ReputationTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ReputationTable) Reputations(obj *Reputation, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ReputationTable) ReputationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ReputationTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ReputationTableAddReputations(builder *flatbuffers.Builder, Reputations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Reputations), 0)
}
func ReputationTableStartReputationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ReputationTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Reputation struct {
	_tab flatbuffers.Table
}

func GetRootAsReputation(buf []byte, offset flatbuffers.UOffsetT) *Reputation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Reputation{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Reputation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Reputation) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Reputation) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Reputation) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Reputation) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Reputation) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Reputation) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Reputation) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ReputationStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ReputationAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func ReputationAddNameEn(builder *flatbuffers.Builder, NameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NameEn), 0)
}
func ReputationAddNameFr(builder *flatbuffers.Builder, NameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameFr), 0)
}
func ReputationAddNameDe(builder *flatbuffers.Builder, NameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameDe), 0)
}
func ReputationAddNameJa(builder *flatbuffers.Builder, NameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameJa), 0)
}
func ReputationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
