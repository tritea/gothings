package dataspec

// DataSpec 数据规格接口
type DataSpec interface {

	// 验证接口，验证数据是否符合这个接口规格
	Validate(v interface{}) (bool, error)
}
