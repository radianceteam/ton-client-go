package spec

type TypeName string

const (
	Ref          TypeName = "Ref"
	Optional     TypeName = "Optional"
	String       TypeName = "String"
	Number       TypeName = "Number"
	Value        TypeName = "Value"
	Struct       TypeName = "Struct"
	None         TypeName = "None"
	BigInt       TypeName = "BigInt"
	Boolean      TypeName = "Boolean"
	Array        TypeName = "Array"
	Generic      TypeName = "Generic"
	EnumOfTypes  TypeName = "EnumOfTypes"
	EnumOfConsts TypeName = "EnumOfConsts"
)

var ignoredTypesByName = map[string]bool{
	"ResultOfGetApiReference": true,
	"ClientError":             true,
	"ClientConfig":            true,
}

var ignoredFunctionsByName = map[string]bool{
	"get_api_reference": true,
}