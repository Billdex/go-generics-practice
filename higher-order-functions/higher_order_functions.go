package higher_order_functions

// FilterSlice 对切片数据进行过滤
func FilterSlice[T any](array []T, callback func(T) bool) (newArray []T) {
	for i := range array {
		if callback(array[i]) {
			newArray = append(newArray, array[i])
		}
	}
	return newArray
}

// MapSlice 对切片元素实现映射
func MapSlice[T any, M any](array []T, callback func(T) M) []M {
	newArray := make([]M, len(array))
	for i := range array {
		newArray[i] = callback(array[i])
	}
	return newArray
}

// ReduceSlice 对切片按照reducer()进行迭代，最终输出计算结果
func ReduceSlice[T any, M any](array []T, initialVal M, reducer func(prev M, current T) M) M {
	var result M = initialVal
	for i := range array {
		result = reducer(result, array[i])
	}
	return result
}
