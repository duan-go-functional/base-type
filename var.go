package base_type

func Var(v interface{}) Func {
	return func(...interface{}) interface{} {
		return v
	}
}
