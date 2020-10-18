package helper

import (
	"fmt"
	"testing"
)

func Test_showJSON(t *testing.T) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"a-a": map[string]interface{}{
				"a-a-a": map[string]interface{}{
					"a-a-a-a": "4",
				},
			},
		},
		"b": map[string]interface{}{
			"b-b": map[string]interface{}{
				"b-b-b": map[string]interface{}{
					"b-b-b-b": "4",
				},
			},
		},
		"c": []int{1, 2, 3, 4, 5, 6},
		"d": []int{1, 2, 3, 4, 5, 6, 7},
	}
	fmt.Println("===============all")
	fmt.Println(ShowJSON(data))
	fmt.Println("===============0")
	fmt.Println(ShowJSON(data, 0))
	fmt.Println("===============1")
	fmt.Println(ShowJSON(data, 1))
	fmt.Println("===============2")
	fmt.Println(ShowJSON(data, 2))
	fmt.Println("===============")

}

func Test_powerfulIn(t *testing.T) {

	tests := []struct {
		name     string
		needle   interface{}
		haystack []interface{}
		want     bool
	}{
		{
			name:     "a",
			needle:   2,
			haystack: []interface{}{"1", "2", 3, 4},
			want:     true,
		},
		{
			name:   "b",
			needle: 2,
			haystack: []interface{}{
				1,
				[]string{"1", "2"},
			},
			want: true,
		},
		{
			name:   "c",
			needle: 2,
			haystack: []interface{}{
				map[string]interface{}{
					"2": 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.needle, tt.haystack...); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := append(a, 100)
	c := append(a, 200)

	t.Log("b:", b)
	t.Log("c:", c)

	var d []int
	t.Log("d:", d == nil)

}
