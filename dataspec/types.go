package dataspec

// DataType 数据类型
type DataType string

const (
	StringType  DataType = "string"
	IntegerType DataType = "integer"
	NumberType  DataType = "number"
	BooleanType DataType = "boolean"
	ArrayType   DataType = "array"
	StructType  DataType = "struct"
	VoidType    DataType = "void"
)
