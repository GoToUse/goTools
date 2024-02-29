// collections.go is simply like python's collections library
// to furnish many convenient and powerful tools to
// help us to develop more fast.

// iterables: string, map, slice

package goTools

import (
	"fmt"
	"reflect"
)

type CTuple struct {
	elem  any
	times int
}

type CounterInterface interface {
	Elements() []any
	MostCommon(n int) []CTuple
	Total() int
	Clear()
	Get(elem any) int
}

type Counter struct {
	data map[any]int
}

func isSlice(v any) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Slice
}

func isMap(v any) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Map
}

func NewCounter(elem any) *Counter {
	summary := make(map[any]int)

	switch elem.(type) {
	case string:
		for _, v := range (elem).(string) {
			summary[string(v)]++
		}
	case map[any]int:
		summary = (elem).(map[any]int)

	default:
		switch {
		case isSlice(elem):
			value := reflect.ValueOf(elem)
			for i := 0; i < value.Len(); i++ {
				v := value.Index(i).Interface()
				summary[v]++
			}
		case isMap(elem):
			value := reflect.ValueOf(elem)

			for _, key := range value.MapKeys() {
				v := value.MapIndex(key).Interface().(int)
				summary[key.Interface()] = v
			}
		default:
			fmt.Println(fmt.Errorf("the parameter must be one of string, slice, map[any]int, got: %v, type: %T", elem, elem).Error())
			return nil
		}
	}

	fmt.Println("summary", summary)
	return &Counter{summary}
}

func (c *Counter) Elements() []any {
	return []any{}
}
