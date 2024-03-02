// collections.go is simply like python's collections library
// to furnish many convenient and powerful tools to
// help us to develop more fast.

// iterables: string, map, slice

package goTools

import (
	"fmt"
	"reflect"
	"sort"
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

func getOrderedKeys(data map[any]int) []any {
	keys := make([]any, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		dataI, dataJ := data[keys[i]], data[keys[j]]

		return dataI > dataJ
	})

	return keys
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
				v := value.MapIndex(key).Interface()
				if vv, ok := v.(int); !ok {
					fmt.Println(fmt.Errorf("[Error]:the value of map must be int, got: %v, type: %T", v, v).Error())
					return nil
				} else {
					summary[key.Interface()] = vv
				}
			}
		default:
			fmt.Println(fmt.Errorf("[Error]:the parameter must be one of string, slice, map[any]int, got: %v, type: %T", elem, elem).Error())
			return nil
		}
	}

	return &Counter{summary}
}

func (c *Counter) Elements() []any {
	elems := make([]any, 0)

	for k, v := range c.data {
		tmp := make([]any, v)

		for i := 0; i < v; i++ {
			tmp[i] = k
		}
		elems = append(elems, tmp...)
	}

	return elems
}

func (c *Counter) Total() int {
	_total := 0
	for _, v := range c.data {
		_total += v
	}

	return _total
}

func (c *Counter) Clear() {
	c.data = nil
}

func (c *Counter) MostCommon(n int) (result []CTuple) {
	ordered := getOrderedKeys(c.data)

	for i := 0; i < n; i++ {
		result = append(result, CTuple{
			elem:  ordered[i],
			times: c.data[ordered[i]],
		})
	}

	return result
}

func (c *Counter) Get(elem any) int {
	if v, ok := c.data[elem]; ok {
		return v
	} else {
		fmt.Println(fmt.Errorf("[Error]:the key is not existence, got: %v, data: %v", v, c.data).Error())
		return 0
	}
}
