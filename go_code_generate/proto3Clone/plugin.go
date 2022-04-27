package proto3clone

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	generator "go_gen/generate"
)

// Plugin is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for Plugin support.
type Plugin struct {
	gen   *generator.Generator
	clone *GenClone
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
	g.clone = NewGenClone(gen)
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
	for _, m := range file.FileDescriptorProto.MessageType {
		mName := generator.CamelCase(m.GetName())
		if _, ok := g.clone.aggregateObject[mName]; ok {
			continue
		}

		for _, l := range m.NestedType {
			// not found map key and value
			if len(l.GetField()) != 2 {
				panic("get NestedType not found key and value")
			}
			keyType := getValueType(l.GetField()[0])
			valueType := getValueType(l.GetField()[1])
			q, _ := spiltStar(valueType)
			funcName := "CloneMap" + generator.CamelCase(keyType) + generator.CamelCase(q)
			g.clone.mapFunc[mName] = append(g.clone.mapFunc[mName], &mapInfo{
				funcName:  funcName,
				mapName:   l.GetName(), // TemEntry
				keyType:   keyType,
				valueType: valueType,
			})
		}
		for _, f := range m.Field {
			bName := generator.CamelCase(f.GetName())
			valueType := getValueType(f) // string,*AppleServerIn,int32
			switch f.GetLabel() {
			case descriptor.FieldDescriptorProto_LABEL_REPEATED:
				if getValueType(f) == "map" { //map 类型
					op, ok := g.clone.mapFunc[mName]
					if !ok {
						panic("not found map key")
					}
					var funcName string
					for _, info := range op {
						if info.mapName == generator.CamelCase(f.GetName())+"Entry" {
							funcName = info.funcName
							break
						}
					}
					g.clone.aggregateObject[mName] = append(g.clone.aggregateObject[mName], &maps{
						field:    bName,
						funcName: funcName,
					})
				} else { // repeated
					b, _ := spiltStar(valueType)
					funcName := "CloneArray" + b
					g.clone.aggregateObject[mName] = append(g.clone.aggregateObject[mName], &repeated{
						field:    bName,
						funcName: funcName,
					})
					if _, ok := g.clone.arrayFunc[funcName]; ok {
						continue
					}
					g.clone.arrayFunc[funcName] = &arrInfo{
						valueType: valueType,
						cb:        g.clone.genArrayTypeFunc,
					}
				}
			case descriptor.FieldDescriptorProto_LABEL_OPTIONAL:
				if _, ok := spiltStar(valueType); ok {
					g.clone.aggregateObject[mName] = append(g.clone.aggregateObject[mName], &optionalObject{
						field: bName,
					})
				} else {
					g.clone.aggregateObject[mName] = append(g.clone.aggregateObject[mName], &optionalNormal{
						field: bName,
					})
				}
			default:
				panic("label_type_err")
			}
		}
	}
	g.clone.generateClone()
}

// GenerateImports generates the import declaration for this file.
func (g *Plugin) GenerateImports(file *generator.FileDescriptor) {
	// g.P("var _ = context.Backgroud()")
}

func getValueType(fType *descriptor.FieldDescriptorProto) string {
	switch fType.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "byte"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		res := strings.Split(fType.GetTypeName(), ".")
		if len(res) == 0 {
			panic("message_type not found typeName")
		}
		msgType := generator.CamelCase(res[len(res)-1])
		// judge map or repeated struct
		if generator.CamelCase(fType.GetName())+"Entry" == msgType { // map
			return "map"
		} else { // repeated struct
			return "*" + msgType
		}
	default:
		panic("type_err")
	}
}

func spiltStar(str string) (string, bool) {
	if len(str) == 0 {
		return "", false
	}
	if string(str[0]) == "*" {
		return str[1:], true
	}
	return str, false
}
