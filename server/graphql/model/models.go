// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewSetting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Return struct {
	Ok   bool   `json:"ok"`
	Info string `json:"info"`
}

type Setting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Admin bool   `json:"admin"`
}