// Code generated by protoc-gen-base. DO NOT EDIT.
// source: test.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	"go_gen/proto3cloneV2/example/sometype"
	_ "go_gen/proto3cloneV2/example/sometype"
	"go_gen/proto3cloneV2/example2"
	_ "go_gen/proto3cloneV2/example2"
	"go_gen/proto3cloneV2/example3"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func CloneRequestServiceDebugCmd(in *RequestServiceDebugCmd) *RequestServiceDebugCmd {
	if in == nil {
		return nil
	}
	out := &RequestServiceDebugCmd{}
	out.Cmd = in.Cmd
	out.Commanders = append(in.Commanders[:0:0], in.Commanders...)
	out.Binary = clone_RequestServiceDebugCmd_Array_sometype_applesServerin(in.Binary)
	out.SandBox = clone_RequestServiceDebugCmd_Array_example_localWay(in.SandBox)
	out.Apple = clone_RequestServiceDebugCmd_sometype_applesServerin(in.Apple)
	out.Ll = clone_RequestServiceDebugCmd_example_localWay(in.Ll)
	out.Tem = clone_RequestServiceDebugCmd_Map_string_int32(in.Tem)
	out.Tem2 = clone_RequestServiceDebugCmd_Map_string_example_localWay(in.Tem2)
	out.AddrEntry = clone_RequestServiceDebugCmd_Map_string_sometype_applesServerin(in.AddrEntry)
	out.Hero = clone_RequestServiceDebugCmd_example2_Hero(in.Hero)
	return out
}

func CloneRequestServiceDebugCmd2(in *RequestServiceDebugCmd2) *RequestServiceDebugCmd2 {
	if in == nil {
		return nil
	}
	out := &RequestServiceDebugCmd2{}
	out.Cmd = in.Cmd
	out.Commanders = append(in.Commanders[:0:0], in.Commanders...)
	out.Apple = clone_RequestServiceDebugCmd2_sometype_applesServerin(in.Apple)
	out.Tem = clone_RequestServiceDebugCmd2_Map_string_int32(in.Tem)
	out.AddrEntry = clone_RequestServiceDebugCmd2_Map_string_sometype_applesServerin(in.AddrEntry)
	out.Hero = clone_RequestServiceDebugCmd2_example2_Hero(in.Hero)
	return out
}

func clone_RequestServiceDebugCmd_Map_string_example_localWay(in map[string]*LocalWay) map[string]*LocalWay {
	if in == nil {
		return nil
	}
	a := make(map[string]*LocalWay, len(in))
	for k, v := range in {
		a[k] = clone_RequestServiceDebugCmd_example_localWay(v)
	}
	return a
}

func clone_RequestServiceDebugCmd_Map_string_sometype_applesServerin(in map[string]*sometype.ApplesServerin) map[string]*sometype.ApplesServerin {
	if in == nil {
		return nil
	}
	a := make(map[string]*sometype.ApplesServerin, len(in))
	for k, v := range in {
		a[k] = clone_RequestServiceDebugCmd_sometype_applesServerin(v)
	}
	return a
}

func clone_RequestServiceDebugCmd2_Map_string_int32(in map[string]int32) map[string]int32 {
	if in == nil {
		return nil
	}
	a := make(map[string]int32, len(in))
	for k, v := range in {
		a[k] = v
	}
	return a
}

func clone_RequestServiceDebugCmd2_Map_string_sometype_applesServerin(in map[string]*sometype.ApplesServerin) map[string]*sometype.ApplesServerin {
	if in == nil {
		return nil
	}
	a := make(map[string]*sometype.ApplesServerin, len(in))
	for k, v := range in {
		a[k] = clone_RequestServiceDebugCmd2_sometype_applesServerin(v)
	}
	return a
}

func clone_RequestServiceDebugCmd_Map_string_int32(in map[string]int32) map[string]int32 {
	if in == nil {
		return nil
	}
	a := make(map[string]int32, len(in))
	for k, v := range in {
		a[k] = v
	}
	return a
}

func clone_RequestServiceDebugCmd_Array_sometype_applesServerin(in []*sometype.ApplesServerin) []*sometype.ApplesServerin {
	if in == nil {
		return nil
	}
	a := make([]*sometype.ApplesServerin, len(in))
	for k, v := range in {
		a[k] = clone_RequestServiceDebugCmd_sometype_applesServerin(v)
	}
	return a
}

func clone_RequestServiceDebugCmd_Array_example_localWay(in []*LocalWay) []*LocalWay {
	if in == nil {
		return nil
	}
	a := make([]*LocalWay, len(in))
	for k, v := range in {
		a[k] = clone_RequestServiceDebugCmd_example_localWay(v)
	}
	return a
}

func clone_RequestServiceDebugCmd_sometype_applesServerin(in *sometype.ApplesServerin) *sometype.ApplesServerin {
	if in == nil {
		return nil
	}
	out := &sometype.ApplesServerin{}
	out.Id = in.Id
	out.Name = in.Name
	out.Student = append(in.Student[:0:0], in.Student...)
	out.Teacher = append(in.Teacher[:0:0], in.Teacher...)
	out.Bb = clone_RequestServiceDebugCmd_sometype_Binary(in.Bb)
	return out
}

func clone_RequestServiceDebugCmd_sometype_Binary(in *sometype.Binary) *sometype.Binary {
	if in == nil {
		return nil
	}
	out := &sometype.Binary{}
	out.Id = append(in.Id[:0:0], in.Id...)
	out.Hero = clone_RequestServiceDebugCmd_example2_Hero(in.Hero)
	out.Tt = clone_RequestServiceDebugCmd_example3_Teacher(in.Tt)
	return out
}

func clone_RequestServiceDebugCmd_example2_Hero(in *example2.Hero) *example2.Hero {
	if in == nil {
		return nil
	}
	out := &example2.Hero{}
	out.Id = in.Id
	return out
}

func clone_RequestServiceDebugCmd_example3_Teacher(in *example3.Teacher) *example3.Teacher {
	if in == nil {
		return nil
	}
	out := &example3.Teacher{}
	out.Id = in.Id
	return out
}

func clone_RequestServiceDebugCmd_example_localWay(in *LocalWay) *LocalWay {
	if in == nil {
		return nil
	}
	out := &LocalWay{}
	out.Id = in.Id
	return out
}

func clone_RequestServiceDebugCmd2_sometype_applesServerin(in *sometype.ApplesServerin) *sometype.ApplesServerin {
	if in == nil {
		return nil
	}
	out := &sometype.ApplesServerin{}
	out.Id = in.Id
	out.Name = in.Name
	out.Student = append(in.Student[:0:0], in.Student...)
	out.Teacher = append(in.Teacher[:0:0], in.Teacher...)
	out.Bb = clone_RequestServiceDebugCmd2_sometype_Binary(in.Bb)
	return out
}

func clone_RequestServiceDebugCmd2_sometype_Binary(in *sometype.Binary) *sometype.Binary {
	if in == nil {
		return nil
	}
	out := &sometype.Binary{}
	out.Id = append(in.Id[:0:0], in.Id...)
	out.Hero = clone_RequestServiceDebugCmd2_example2_Hero(in.Hero)
	out.Tt = clone_RequestServiceDebugCmd2_example3_Teacher(in.Tt)
	return out
}

func clone_RequestServiceDebugCmd2_example2_Hero(in *example2.Hero) *example2.Hero {
	if in == nil {
		return nil
	}
	out := &example2.Hero{}
	out.Id = in.Id
	return out
}

func clone_RequestServiceDebugCmd2_example3_Teacher(in *example3.Teacher) *example3.Teacher {
	if in == nil {
		return nil
	}
	out := &example3.Teacher{}
	out.Id = in.Id
	return out
}