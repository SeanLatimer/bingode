// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MinionTable struct {
	_tab flatbuffers.Table
}

func GetRootAsMinionTable(buf []byte, offset flatbuffers.UOffsetT) *MinionTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MinionTable{}
	x.Init(buf, n+offset)
	return x
}

func FinishMinionTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMinionTable(buf []byte, offset flatbuffers.UOffsetT) *MinionTable {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MinionTable{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMinionTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MinionTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MinionTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MinionTable) Minions(obj *Minion, j int) bool {
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

func (rcv *MinionTable) MinionsByKey(obj *Minion, key uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *MinionTable) MinionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MinionTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MinionTableAddMinions(builder *flatbuffers.Builder, minions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(minions), 0)
}
func MinionTableStartMinionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MinionTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Minion struct {
	_tab flatbuffers.Table
}

func GetRootAsMinion(buf []byte, offset flatbuffers.UOffsetT) *Minion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Minion{}
	x.Init(buf, n+offset)
	return x
}

func FinishMinionBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMinion(buf []byte, offset flatbuffers.UOffsetT) *Minion {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Minion{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMinionBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Minion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Minion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Minion) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Minion) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func MinionKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Minion{}
	obj2 := &Minion{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf))-o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf))-o2)
	return obj1.Id() < obj2.Id()
}

func (rcv *Minion) LookupByKey(key uint32, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation-4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+4*(start+middle))
		obj := &Minion{}
		obj.Init(buf, tableOffset)
		val := obj.Id()
		comp := 0
		if val > key {
			comp = 1
		} else if val < key {
			comp = -1
		}
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *Minion) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Minion) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MinionStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MinionAddId(builder *flatbuffers.Builder, id uint32) {
	builder.PrependUint32Slot(0, id, 0)
}
func MinionAddNameEn(builder *flatbuffers.Builder, nameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nameEn), 0)
}
func MinionAddNameFr(builder *flatbuffers.Builder, nameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nameFr), 0)
}
func MinionAddNameDe(builder *flatbuffers.Builder, nameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(nameDe), 0)
}
func MinionAddNameJa(builder *flatbuffers.Builder, nameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(nameJa), 0)
}
func MinionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
