package docs

import (
	"reflect"
	"testing"
)

func BenchmarkCalls(b *testing.B) {
	arg1 := "arg1"
	arg2 := "arg2"
	expectedRes := "arg1arg2"

	b.Run("simple call", func(b *testing.B) {
		res := toCall(arg1, arg2)
		if res != expectedRes {
			b.Error("unknown")
		}
	})

	b.Run("reflect call", func(b *testing.B) {
		toCallValue := reflect.ValueOf(toCall)
		arg1Val := reflect.ValueOf(arg1)
		arg2Val := reflect.ValueOf(arg2)
		resVal := toCallValue.Call([]reflect.Value{arg1Val, arg2Val})

		if resVal[0].Interface() != expectedRes {
			b.Error("unknown")
		}
	})
}

func toCall(arg1, arg2 string) string {
	return arg1 + arg2
}
