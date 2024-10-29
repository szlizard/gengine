package base

import (
	"fmt"
	"reflect"
	"strings"
)

func evaluateInOperator(left reflect.Value, right reflect.Value) (bool, error) {
	// 优化性能，减少重复反射调用
	leftInterface := left.Interface()
	leftKind := left.Kind()

	switch right.Kind() {
	case reflect.String:
		if leftKind == reflect.String {
			return strings.Contains(right.String(), left.String()), nil
		}
		return false, fmt.Errorf("left operand must be a string when right operand is a string")
	case reflect.Slice:
		for i := 0; i < right.Len(); i++ {
			rightElement := right.Index(i).Interface()
			rightValue := reflect.ValueOf(rightElement)
			// 优化性能：减少反射类型转换次数
			if rightValue.Kind() == leftKind || rightValue.Type().ConvertibleTo(left.Type()) {
				if rightValue.Type().ConvertibleTo(left.Type()) {
					rightValue = rightValue.Convert(left.Type())
				}

				// 使用严格的反射比较，确保浮点数一致
				if leftKind == reflect.Float64 && rightValue.Kind() == reflect.Float64 {
					if left.Float() == rightValue.Float() {
						return true, nil
					}
				} else {
					// 对非浮点数类型进行比较
					if reflect.DeepEqual(leftInterface, rightValue.Interface()) {
						return true, nil
					}
				}
			}
		}
		return false, nil
	case reflect.Map:
		if right.MapIndex(left).IsValid() {
			return true, nil
		}
		return false, nil
	default:
		return false, fmt.Errorf("unsupported type for right operand in `in` operator")
	}
}
