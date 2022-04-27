package proto3cloneV2

import (
	"strings"

	generator "go_gen/generate"
)

const (
	fieldArray  = "array"
	fieldMap    = "map"
	fieldNormal = "normal"
)

type fileGen struct {
	gen            *generator.Generator
	sourceFileMap  map[string]*filesAllMsg   // key:文件名，如example.proto
	annotationFlag map[string][]*flagStr     // annotation struct:直接或间接引用的所有对象
	funcMapStore   map[string]*funcStoreInfo // key:funcName
	funcArrayStore map[string]*funcStoreInfo // key:funcName
}

// all files message
type filesAllMsg struct {
	packageName string                   //
	allObjects  map[string][]*fieldsDesc //structName:
}

// all field desc
type fieldsDesc struct {
	fieldName     string //
	fieldTypeDesc string // 字段类型描述，三种：array，map，normalValue
	keyType       string // 只有map使用,为map的key类型
	packageName   string
	valueType     string // value类型
	isBaseType    bool
}

type flagStr struct {
	objPackage string
	objName    string
}

type funcStoreInfo struct {
	funcName     string
	keyType      string
	valueType    string
	isBaseType   bool
	filePackage  string
	valuePackage string
	objectName   string
}

func (f *fieldsDesc) setFieldDesc(tt, defaultPackage string) {
	var (
		packageName string
		typeName    string
		isBaseType  bool
	)
	if strings.Contains(tt, ".") {
		res := strings.Split(tt, ".")
		if len(res) != 2 {
			panic("err")
		}
		packageName = res[0]
		typeName = res[1]
	} else {
		packageName = defaultPackage
		typeName = tt
		isBaseType = true
	}
	f.packageName = packageName
	f.valueType = typeName
	f.isBaseType = isBaseType
}

func NewFileGen(gen *generator.Generator) *fileGen {
	return &fileGen{
		gen:            gen,
		sourceFileMap:  make(map[string]*filesAllMsg),
		annotationFlag: make(map[string][]*flagStr),
		funcMapStore:   make(map[string]*funcStoreInfo),
		funcArrayStore: make(map[string]*funcStoreInfo),
	}
}

func (g *fileGen) ParseCloneFile(files *filesAllMsg) {
	for structName, allFields := range files.allObjects {
		// filter no annotation object
		if _, ok := g.annotationFlag[structName]; !ok {
			continue
		}
		g.parseObject(structName, allFields)
	}
}

func (g *fileGen) parseObject(structName string, allFields []*fieldsDesc) {
	// 解析出这个结构体直接引用和间接引用的所有对象
	for _, v := range allFields {
		// 过滤非对象字段
		if v.isBaseType {
			continue
		}
		// 过滤一下未注释的对象
		mm, ok := g.annotationFlag[structName]
		if !ok {
			return
		}
		// 过滤一下已经存在的对象名
		if !containString(mm, v.valueType) {
			g.annotationFlag[structName] = append(g.annotationFlag[structName], &flagStr{
				objPackage: v.packageName,
				objName:    v.valueType,
			})
			// 查找子对象
			rs := g.findObjAllFields(v.packageName, v.valueType)
			g.parseObject(structName, rs)
		}
	}
}

func containString(arr []*flagStr, f string) bool {
	for _, v := range arr {
		if v.objName == f {
			return true
		}
	}
	return false
}

func (g *fileGen) findObjAllFields(packageName, valueType string) []*fieldsDesc {
	fileName := packageName + ".proto"
	objs, ok := g.sourceFileMap[fileName]
	if !ok {
		panic("findObjAllFields error get file")
	}
	rs, ok2 := objs.allObjects[strings.Title(valueType)]
	if !ok2 {
		panic("findObjAllFields error get objects")
	}
	return rs
}

func (g *fileGen) GenClone(filePackage string) {
	for _, fileMsg := range g.sourceFileMap {
		for obName, v := range fileMsg.allObjects {
			if _, ok := g.annotationFlag[obName]; !ok {
				continue
			}
			g.genParentCloneWay(filePackage, obName, v)
		}
	}
	g.genMapFuncWay()
	g.genArrayFuncWay()
	g.genSonCloneWay(filePackage)
}

// generate export way
func (g *fileGen) genParentCloneWay(filePackage string, obName string, allFields []*fieldsDesc) {
	g.gen.P()
	g.gen.P("func Clone", obName, "(in *", obName, ") *", obName, " {")
	g.gen.P("if in == nil {")
	g.gen.P("return nil")
	g.gen.P("}")
	g.gen.P("out := &", obName, "{}")
	g.genObjectClone(filePackage, obName, allFields)
	g.gen.P("return out")
	g.gen.P("}")
}

func (g *fileGen) genObjectClone(filePackage string, obName string, allFields []*fieldsDesc) {
	for _, v := range allFields {
		fieldName := strings.Title(v.fieldName)
		var (
			funcMapName   string
			funcArrayName string
		)
		if v.isBaseType {
			if v.fieldTypeDesc == fieldNormal {
				g.gen.P("out.", fieldName, " = in.", fieldName)
			} else if v.fieldTypeDesc == fieldArray {
				g.gen.P("out.", fieldName, " = ", "append(in.", fieldName, "[:0:0],", "in.", fieldName, "...)")
			} else if v.fieldTypeDesc == fieldMap {
				funcMapName = "clone_" + obName + "_Map_" + v.keyType + "_" + v.valueType
				g.gen.P("out.", fieldName, " = ", funcMapName, "(in.", fieldName, ")")
			}
		} else {
			if v.fieldTypeDesc == fieldNormal {
				g.gen.P("out.", fieldName, " = clone_", obName, "_", v.packageName, "_", v.valueType, "(in.", fieldName, ")")
			} else if v.fieldTypeDesc == fieldArray {
				funcArrayName = "clone_" + obName + "_Array_" + v.packageName + "_" + v.valueType
				g.gen.P("out.", fieldName, " = ", funcArrayName, "(in.", fieldName, ")")
			} else if v.fieldTypeDesc == fieldMap {
				funcMapName = "clone_" + obName + "_Map_" + v.keyType + "_" + v.packageName + "_" + v.valueType
				g.gen.P("out.", fieldName, " = ", funcMapName, "(in.", fieldName, ")")
			}
		}
		if funcMapName != "" {
			if _, ok := g.funcMapStore[funcMapName]; !ok {
				g.funcMapStore[funcMapName] = &funcStoreInfo{
					funcName:     funcMapName,
					keyType:      v.keyType,
					valueType:    v.valueType,
					isBaseType:   v.isBaseType,
					filePackage:  filePackage,
					valuePackage: v.packageName,
					objectName:   obName,
				}
			}
		}
		if funcArrayName != "" {
			if _, ok := g.funcArrayStore[funcArrayName]; !ok {
				g.funcArrayStore[funcArrayName] = &funcStoreInfo{
					funcName:     funcArrayName,
					valueType:    v.valueType,
					filePackage:  filePackage,
					valuePackage: v.packageName,
					objectName:   obName,
				}
			}
		}

	}
}

func (g *fileGen) genMapFuncWay() {
	for _, v := range g.funcMapStore {
		value := v.valueType
		// 判断一下是不是同包
		if !v.isBaseType {
			if v.filePackage != v.valuePackage {
				value = "*" + v.valuePackage + "." + strings.Title(value)
			} else {
				value = "*" + strings.Title(value)
			}
		}
		g.gen.P()
		g.gen.P("func ", v.funcName, "(in map[", v.keyType, "]", value, ") map[", v.keyType, "]", value, " {")
		g.gen.P("if in == nil {")
		g.gen.P("return nil")
		g.gen.P("}")
		g.gen.P("a := make(map[", v.keyType, "]", value, ", len(in))")
		g.gen.P("for k,v := range in {")
		if v.isBaseType {
			g.gen.P("a[k] = v")
		} else {
			fn := "clone_" + v.objectName + "_" + v.valuePackage + "_" + v.valueType
			g.gen.P("a[k] = ", fn+"(v)")
		}
		g.gen.P("}")
		g.gen.P("return a")
		g.gen.P("}")
	}
}

func (g *fileGen) genArrayFuncWay() {
	for _, v := range g.funcArrayStore {
		value := v.valueType
		if v.filePackage != v.valuePackage {
			value = v.valuePackage + "." + strings.Title(value)
		} else {
			value = strings.Title(value)
		}
		g.gen.P()
		g.gen.P("func ", v.funcName, "(in []*", value, ") []*", value, " {")
		g.gen.P("if in == nil {")
		g.gen.P("return nil")
		g.gen.P("}")
		g.gen.P("a := make([]*", value, ",len(in))")
		g.gen.P("for k,v := range in {")
		g.gen.P("a[k] = ", "clone_", v.objectName, "_", v.valuePackage, "_", v.valueType, "(v)")
		g.gen.P("}")
		g.gen.P("return a")
		g.gen.P("}")
	}
}

// generate no export ways
func (g *fileGen) genSonCloneWay(filePackage string) {
	for topObjName, v := range g.annotationFlag {
		for _, flags := range v {
			g.sonObjClone(filePackage, topObjName, flags.objName, flags.objPackage)
		}
	}
}

func (g *fileGen) sonObjClone(filePackage, topObjName, objName, objSelfPackName string) {
	for _, v := range g.sourceFileMap {
		obn := strings.Title(objName)
		mm, ok := v.allObjects[obn]
		if !ok {
			continue
		}

		funcName := "clone_" + topObjName + "_" + objSelfPackName + "_" + objName
		var value string
		if filePackage == objSelfPackName {
			value = obn
		} else {
			value = objSelfPackName + "." + obn
		}
		g.gen.P()
		g.gen.P("func ", funcName, "(in *", value, ") *", value, " {")
		g.gen.P("if in == nil {")
		g.gen.P("return nil")
		g.gen.P("}")
		g.gen.P("out := &", value, "{}")
		g.genObjectClone(filePackage, topObjName, mm)
		g.gen.P("return out")
		g.gen.P("}")
	}
}
