package spec

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func genEnum(t Type) string {
	for i := range t.EnumConsts {
		t.EnumConsts[i].GoComment = t.EnumConsts[i].ToComment()
		t.EnumConsts[i].ConstName = toGoName(strcase.ToSnake(t.EnumConsts[i].ConstName)) + t.Name
	}
	var tpl bytes.Buffer
	if err := enumTmpl.Execute(&tpl, t); err != nil {
		panic(err)
	}

	return tpl.String()
}

func prepareMultiline(data string) string {
	parts := strings.Split(data, "\n")

	for i := range parts {
		parts[i] = strings.Trim(parts[i], " ")
	}

	return strings.Join(parts, "\n// ")
}

// ExceptionForCamelCase - camel case names exception.
var ExceptionForCamelCase = map[string]string{
	"id":                        "ID",
	"xprv":                      "XPrv",
	"url":                       "URL",
	"account_id_address_string": "AccountIDAddressString",
}

func toGoName(name string) string {
	if cameled, ok := ExceptionForCamelCase[name]; ok {
		return cameled
	}

	cameled := strcase.ToCamel(name)
	if strings.Contains(name, "id") {
		cameled = strings.ReplaceAll(cameled, "Id", "ID")
	}

	return cameled
}

func toTypeName(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) != 1 {
		return parts[1]
	}

	return name
}

func genFunc(m Module, f Function) string {
	params := make([]Type, 0, 2)
	for _, p := range f.Params {
		if p.Name == "context" || p.Name == "_context" {
			continue
		}
		params = append(params, p)
	}
	if len(params) > 1 {
		fmt.Println("WARNING: ignored function", len(params), m.Name, f.Name)

		return ""
	}

	var b bytes.Buffer
	content := funcContent{
		Name:       strcase.ToCamel(m.Name + "_" + f.Name),
		MethodName: m.Name + "." + f.Name,
		ResultType: toTypeName(f.Result.GenericArgs[0].RefName),
	}

	var tmpl *template.Template
	if len(params) == 0 {
		tmpl = withoutParamFunc
	} else if len(params) == 1 {
		content.ParamType = toTypeName(params[0].RefName)
		if content.ResultType != "" {
			tmpl = singleParamFunc
		} else {
			tmpl = singleParamWithoutResultFunc
		}
	}

	if err := tmpl.Execute(&b, content); err != nil {
		panic(err)
	}

	return "\n" + f.ToComment() + b.String()
}

func GenModule(dir string, m Module) error {
	file, err := os.Create(dir + "/mod_" + m.Name + ".gen.go")
	if err != nil {
		return err
	}
	m.Description.GoComment = m.Description.ToComment()
	if err = headerTmpl.Execute(file, m); err != nil {
		return err
	}
	for _, t := range m.Types {
		if ignoredTypesByName[t.Name] {
			continue
		}
		_, err = file.WriteString(GenerateAnyType(m, t))
		if err != nil {
			return err
		}
	}
	for _, f := range m.Functions {
		if ignoredFunctionsByName[f.Name] {
			continue
		}
		_, err = file.WriteString(genFunc(m, f))
		if err != nil {
			return err
		}
	}

	return nil
}

// genStruct - generates struct with each field specified.
func genStruct(m Module, t Type) string {
	r := "type " + toTypeName(t.Name) + " struct {\n"
	for _, f := range t.StructFields {
		if f.Name == "" {
			fmt.Println("WARNING: add struct field with empty name", t.Type, t.Name, f)
			f.Name = "value"
		}
		r += "\t" + f.ToComment() + "	" + toGoName(f.Name) + " " + GenerateAnyType(m, f)
		if f.Type == Optional {
			r += " `json:\"" + f.Name + "\"` // optional \n"
		} else {
			r += " `json:\"" + f.Name + "\"`\n"
		}
	}
	r += "}\n\n"

	return r
}

// genEnumOfTypes - generates enum of types with special enumOfConsts helper type.
func genEnumOfTypes(m Module, t Type) string {
	enumHelperType := Type{Type: EnumOfConsts, Description: Description{Name: t.Name + "Type"}}
	fields := make(map[string]int)
	structFields := make([]Type, 1)
	structFields[0] = Type{Type: Ref, RefName: enumHelperType.Name, Description: Description{Name: "type"}}
	fields["type"] = 0
	for _, et := range t.EnumTypes {
		enumHelperType.EnumConsts = append(enumHelperType.EnumConsts, Description{
			Name:        et.Name,
			Summary:     et.Summary,
			Description: et.Description.Description,
			ConstName:   et.Name,
		})
		if et.Type != Struct {
			panic("EnumOfTypes only supports structs " + et.Name)
		}
		for _, sf := range et.StructFields {
			index, ok := fields[sf.Name]
			if ok && structFields[index].Type != sf.Type {
				panic("type mismatch for enum " + et.Name + ":" + sf.Name)
			}

			if !ok {
				index = len(structFields)
				fields[sf.Name] = index
				structFields = append(structFields, sf)
				structFields[index].Description.Description += " presented in types:"
			}
			structFields[index].Description.Description += "\n\"" + et.Name + "\""
		}
	}
	t.StructFields = structFields

	return genEnum(enumHelperType) + "\n" + genStruct(m, t)
}

var optionNumberToGoType = map[string]string{
	"Int8":    "null.Int8",
	"Int16":   "null.Int16",
	"Int32":   "null.Int32",
	"Int64":   "null.Int64",
	"UInt8":   "null.Uint8",
	"UInt16":  "null.Uint16",
	"UInt32":  "null.Uint32",
	"UInt64":  "null.Uint64",
	"Float32": "null.Float32",
}

var numberToGoType = map[string]string{
	"Int8":   "int8",
	"Int16":  "int16",
	"Int32":  "int32",
	"Int64":  "int64",
	"UInt8":  "uint8",
	"UInt16": "uint16",
	"UInt32": "uint32",
	"UInt64": "uint64",
}

func genNumber(t Type, isOptional bool) string {
	if isOptional {
		numberType, found := optionNumberToGoType[t.NumberType+strconv.Itoa(t.NumberSize)]
		if found {
			return numberType
		}
		panic(fmt.Sprintf("opt number type not found %s %d", t.NumberType, t.NumberSize))
	}
	numberType, found := numberToGoType[t.NumberType+strconv.Itoa(t.NumberSize)]
	if found {
		return numberType
	}
	panic(fmt.Sprintf("req number type not found %s %d", t.NumberType, t.NumberSize))
}

// GenerateOptionalType - pointer or null simple type.
func GenerateOptionalType(m Module, t Type) string {
	switch t.Type { // nolint exhaustive
	case Ref:
		if t.RefName == "Value" {
			return emptyInterface
		}

		return "*" + toTypeName(t.RefName)
	case String:
		return "null.String"
	case Value:
		return emptyInterface
	case Number:
		return genNumber(t, true)
	case BigInt:
		return "*big.Int"
	case None:
		return ""
	case Array:
		return "[]" + GenerateAnyType(m, *t.ArrayItem)
	case Boolean:
		return "null.Bool"
	default:
		return "* " + GenerateAnyType(m, t)
	}
}

// GenerateAnyType - root generator function for type.
func GenerateAnyType(m Module, t Type) string {
	r := "NotFound::" + string(t.Type) + "::"
	switch t.Type {
	case Ref:
		if t.RefName == "Value" {
			r = emptyInterface
		} else {
			r = toTypeName(t.RefName)
		}
	case Optional:
		r = GenerateOptionalType(m, *t.OptionalInner)
	case String:
		r = "string"
	case Value:
		r = emptyInterface
	case Number:
		r = genNumber(t, false)
	case None:
		r = ""
	case Struct:
		r = genStruct(m, t)
	case BigInt:
		r = "big.Int"
	case Array:
		r = "[]" + GenerateAnyType(m, *t.ArrayItem)
	case Boolean:
		r = "bool"
	case Generic:
		r = "GENERIC"
	case EnumOfTypes:
		r = genEnumOfTypes(m, t)
	case EnumOfConsts:
		for i := range t.EnumConsts {
			t.EnumConsts[i].ConstName = strcase.ToCamel(t.EnumConsts[i].Name)
		}
		r = genEnum(t)
	}

	return r
}
