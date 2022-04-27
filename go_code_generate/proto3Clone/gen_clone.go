package proto3clone

import (
	generator "go_gen/generate"
)

type GenClone struct {
	gen             *generator.Generator
	arrayFunc       map[string]*arrInfo        // funName:funInfo
	mapFunc         map[string][]*mapInfo      // structName：mapInfo
	mapFuncName     map[string]struct{}        // funcName:nil
	aggregateObject map[string][]waysInterface //key:structName : allFields
}

// new clone
func NewGenClone(gen *generator.Generator) *GenClone {
	return &GenClone{
		gen:             gen,
		arrayFunc:       make(map[string]*arrInfo),
		mapFunc:         make(map[string][]*mapInfo),
		mapFuncName:     make(map[string]struct{}),
		aggregateObject: make(map[string][]waysInterface),
	}
}

// funcInfo func callBack
type arrInfo struct {
	valueType string
	cb        func(string, string)
}

type mapInfo struct {
	funcName  string
	mapName   string //addrEntry
	keyType   string // string
	valueType string // string,*AppleServerIn
}

type waysInterface interface {
	fieldClone(*generator.Generator)
}

type repeated struct {
	field    string // 字段名
	funcName string // 函数名
}

func (r repeated) fieldClone(gen *generator.Generator) {
	gen.P("out.", r.field, "=", r.funcName, "(in.", r.field, ")")
}

type maps struct {
	field    string // 字段名
	funcName string // 函数名
}

func (m maps) fieldClone(gen *generator.Generator) {
	gen.P("out.", m.field, "=", m.funcName, "(in.", m.field, ")")
}

type optionalNormal struct {
	field string // 字段名
}

func (o optionalNormal) fieldClone(gen *generator.Generator) {
	gen.P("out.", o.field, "=", "in.", o.field)
}

type optionalObject struct {
	field string // 字段名
}

func (o optionalObject) fieldClone(gen *generator.Generator) {
	gen.P("out.", o.field, "=", "in.Get", o.field, "().Clone()")
}

func (g *GenClone) generateClone() {
	for mName, fields := range g.aggregateObject {
		g.gen.P("// ", mName)
		g.gen.P("func (in *", mName, ") Clone() *", mName, "{")
		g.gen.P("if in == nil {")
		g.gen.P("return nil")
		g.gen.P("}")
		g.gen.P("out := &", mName, "{}")
		for _, impl := range fields {
			impl.fieldClone(g.gen)
		}
		g.gen.P("return out")
		g.gen.P("}")
	}

	for k, v := range g.arrayFunc {
		v.cb(k, v.valueType)
	}

	for _, v := range g.mapFunc {
		for _, d := range v {
			// 过滤生成已经有的方法
			_, ok := g.mapFuncName[d.funcName]
			if ok {
				continue
			}
			g.genMapTypeFunc(d.funcName, d.keyType, d.valueType)
			g.mapFuncName[d.funcName] = struct{}{}
		}
	}

}

// funcName string,*AppleServerIn,
func (g *GenClone) genArrayTypeFunc(funcName string, valueType string) {
	var repeatedType string
	if valueType == "byte" {
		repeatedType = "[][]"
	} else {
		repeatedType = "[]"
	}
	g.gen.P()
	g.gen.P("func ", funcName, "(arr ", repeatedType, valueType, ")", repeatedType, valueType, " {")
	g.gen.P("if arr == nil {")
	g.gen.P("return nil")
	g.gen.P("}")
	if _, ok := spiltStar(valueType); ok {
		g.gen.P("a := make([]", valueType, ",len(arr), cap(arr))")
		g.gen.P("for i,v := range arr {")
		g.gen.P("a[i] = v.Clone()")
		g.gen.P("}")
	} else {
		g.gen.P("a := make(", repeatedType, valueType, ",len(arr),cap(arr))")
		g.gen.P("copy(a,", "arr)")
	}
	g.gen.P("return a")
	g.gen.P("}")
}

func (g *GenClone) genMapTypeFunc(funcName string, kt, vt string) {
	g.gen.P()
	t := "v"
	if _, ok := spiltStar(vt); ok {
		t += ".Clone()"
	}
	g.gen.P("func ", funcName, "(m map[", kt, "]", vt, ") map[", kt, "]", vt, "{")
	g.gen.P("if m == nil {")
	g.gen.P("return nil")
	g.gen.P("}")
	g.gen.P("a := make(map[", kt, "]", vt, ",len(m))")
	g.gen.P("for k,v := range m {")
	g.gen.P("a[k] = ", t)
	g.gen.P("}")
	g.gen.P("return a")
	g.gen.P("}")
}
