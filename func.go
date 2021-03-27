package base_type

// 最基本的算子类型，而且不封装类型，就是原生的func(interface{}) interface{}
type Func = func(...interface{}) interface{}
