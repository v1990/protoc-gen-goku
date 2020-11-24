package goku

import "text/template"

type Data map[string]interface{}

func (data Data) Copy() Data {
	out := make(Data)
	for k, v := range data {
		out[k] = v
	}
	return out
}
func (data Data) DoMerge(other Data) {
	for k, v := range other {
		data[k] = v
	}
}

type FuncMap template.FuncMap

func (m FuncMap) Copy() FuncMap {
	out := make(FuncMap)
	for k, v := range m {
		out[k] = v
	}
	return out
}

func (m FuncMap) DoMerge(other FuncMap) {
	for k, v := range other {
		m[k] = v
	}
}
