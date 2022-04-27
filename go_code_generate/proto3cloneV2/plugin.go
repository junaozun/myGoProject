package proto3cloneV2

import (
	"strings"

	generator "go_gen/generate"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Plugin is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for Plugin support.
type Plugin struct {
	gen          *generator.Generator
	fileGenClone *fileGen
}

// Name returns the name of this plugin, "Plugin".
func (g *Plugin) Name() string {
	return "clone"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.

// Init initializes the plugin.
func (g *Plugin) Init(gen *generator.Generator) {
	g.gen = gen
	g.fileGenClone = NewFileGen(gen)
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *Plugin) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *Plugin) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *Plugin) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *Plugin) Generate(file *generator.FileDescriptor) {
	fsGen := &filesAllMsg{
		packageName: file.GetPackage(),
		allObjects:  make(map[string][]*fieldsDesc),
	}
	g.fileGenClone.sourceFileMap[file.GetName()] = fsGen

	for _, m := range file.FileDescriptorProto.MessageType {
		mName := generator.CamelCase(m.GetName())

		tempMap := make([]*fieldsDesc, 0, len(m.NestedType))
		for _, l := range m.NestedType {
			// not found map key and value
			if len(l.GetField()) != 2 {
				panic("get NestedType not found key and value")
			}
			_, keyType := getValueType(l.GetField()[0])
			_, valueType := getValueType(l.GetField()[1])
			fd := &fieldsDesc{
				fieldName:     l.GetName()[:len(l.GetName())-5], // map类型尾部都拼接了Entry，所以这里要删掉
				fieldTypeDesc: fieldMap,
				keyType:       keyType,
			}
			fd.setFieldDesc(valueType, file.GetPackage())
			tempMap = append(tempMap, fd)
		}

		for _, f := range m.Field {
			td, tt := getValueType(f)
			// filter map object
			if td == fieldMap {
				mapFd := tempMap[0]
				fsGen.allObjects[mName] = append(fsGen.allObjects[mName], mapFd)
				tempMap = tempMap[1:]
				continue
			}
			fd := &fieldsDesc{
				fieldName:     f.GetName(),
				fieldTypeDesc: td,
			}
			fd.setFieldDesc(tt, file.GetPackage())
			fsGen.allObjects[mName] = append(fsGen.allObjects[mName], fd)
		}
	}
	// get annotation object
	for _, lc := range file.GetSourceCodeInfo().GetLocation() {
		if lc.GetLeadingComments() == "" {
			continue
		}
		var strName string
		for i, v := range lc.GetLeadingComments() {
			if string(v) == "|" {
				strName = lc.GetLeadingComments()[i+1:]
				break
			}
		}
		if strName == "" {
			panic("get annotation error")
		}
		strName = strings.Replace(strName, " ", "", -1)
		// 去除换行符
		strName = strings.Replace(strName, "\n", "", -1)
		g.fileGenClone.annotationFlag[strName] = nil
	}

	if len(g.fileGenClone.sourceFileMap) != len(g.gen.Request.ProtoFile) {
		return
	}

	for fileName, fileInfo := range g.fileGenClone.sourceFileMap {
		if !g.isSourceFile(fileName) {
			continue
		}
		g.fileGenClone.ParseCloneFile(fileInfo)
	}
	g.fileGenClone.GenClone(file.GetPackage())
}

// GenerateImports generates the import declaration for this file.
func (g *Plugin) GenerateImports(file *generator.FileDescriptor) {
	// g.P("var _ = context.Backgroud()")
}

func getLabelType(label *descriptor.FieldDescriptorProto) string {
	switch label.GetLabel() {
	case descriptor.FieldDescriptorProto_LABEL_REPEATED:
		return fieldArray
	case descriptor.FieldDescriptorProto_LABEL_OPTIONAL:
		return fieldNormal
	default:
		panic("getLabelType label type error")
	}
}

// return fieldTypeDesc,fieldType
func getValueType(fType *descriptor.FieldDescriptorProto) (string, string) {
	switch fType.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return getLabelType(fType), "string"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return getLabelType(fType), "int32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return getLabelType(fType), "int64"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return getLabelType(fType), "byte"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return getLabelType(fType), "uint32"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return getLabelType(fType), "uint64"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		str := fType.GetTypeName()
		if len(str) == 0 {
			panic("message_type not found typeName")
		}
		// delete .
		newStr := str[1:]
		res := strings.Split(newStr, ".")
		if len(res) == 0 {
			panic("message_type split error")
		}

		if len(res) == 3 { // map
			return fieldMap, ""
		} else if len(res) == 2 { // repeated struct (example.localWay)(packageName.structName)
			return getLabelType(fType), newStr
		} else {
			panic("typeName format error")
		}
	default:
		panic("type_err")
	}
}

func (g *Plugin) isSourceFile(file string) bool {
	for _, sourceFile := range g.gen.Request.GetFileToGenerate() {
		if file == sourceFile {
			return true
		}
	}
	return false
}
