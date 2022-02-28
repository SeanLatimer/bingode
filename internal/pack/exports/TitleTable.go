// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TitleTable struct {
	_tab flatbuffers.Table
}

func GetRootAsTitleTable(buf []byte, offset flatbuffers.UOffsetT) *TitleTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TitleTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TitleTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TitleTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TitleTable) Titles(obj *Title, j int) bool {
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

func (rcv *TitleTable) TitlesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TitleTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TitleTableAddTitles(builder *flatbuffers.Builder, Titles flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Titles), 0)
}
func TitleTableStartTitlesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TitleTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Title struct {
	_tab flatbuffers.Table
}

func GetRootAsTitle(buf []byte, offset flatbuffers.UOffsetT) *Title {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Title{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Title) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Title) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Title) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Title) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Title) IsPrefix() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Title) MutateIsPrefix(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Title) NameMasculineEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameMasculineFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameMasculineDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameMasculineJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameFeminineEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameFeminineFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameFeminineDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Title) NameFeminineJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TitleStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func TitleAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func TitleAddIsPrefix(builder *flatbuffers.Builder, IsPrefix bool) {
	builder.PrependBoolSlot(1, IsPrefix, false)
}
func TitleAddNameMasculineEn(builder *flatbuffers.Builder, NameMasculineEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameMasculineEn), 0)
}
func TitleAddNameMasculineFr(builder *flatbuffers.Builder, NameMasculineFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameMasculineFr), 0)
}
func TitleAddNameMasculineDe(builder *flatbuffers.Builder, NameMasculineDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameMasculineDe), 0)
}
func TitleAddNameMasculineJa(builder *flatbuffers.Builder, NameMasculineJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(NameMasculineJa), 0)
}
func TitleAddNameFeminineEn(builder *flatbuffers.Builder, NameFeminineEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(NameFeminineEn), 0)
}
func TitleAddNameFeminineFr(builder *flatbuffers.Builder, NameFeminineFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(NameFeminineFr), 0)
}
func TitleAddNameFeminineDe(builder *flatbuffers.Builder, NameFeminineDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(NameFeminineDe), 0)
}
func TitleAddNameFeminineJa(builder *flatbuffers.Builder, NameFeminineJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(NameFeminineJa), 0)
}
func TitleEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
