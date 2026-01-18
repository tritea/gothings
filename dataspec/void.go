package dataspec

// VoidDataSpec 空数据描述，当为该类型时，代表不需要传入数据，void的spec将被忽略
type VoidDataSpec struct {
}

func (n *VoidDataSpec) Validate(v interface{}) (bool, error) {
	return true, nil
}
