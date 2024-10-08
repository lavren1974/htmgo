package ui

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

type InputProps struct {
	Id             string
	Label          string
	Name           string
	Type           string
	DefaultValue   string
	ValidationPath string
	Children       []h.Ren
}

func Input(props InputProps) h.Ren {
	validation := h.If(props.ValidationPath != "", h.Children(
		h.Post(props.ValidationPath, hx.ChangeEvent),
		h.Attribute("hx-swap", "innerHTML transition:true"),
		h.Attribute("hx-target", "next div"),
	))

	input := h.Input(
		props.Type,
		h.Class("border p-2 rounded"),
		h.If(props.Id != "", h.Id(props.Id)),
		h.If(props.Name != "", h.Name(props.Name)),
		h.If(props.Children != nil, h.Children(props.Children...)),
		h.If(props.DefaultValue != "", h.Attribute("value", props.DefaultValue)),
		validation,
	)

	wrapped := h.Div(
		h.Class("flex flex-col gap-1"),
		h.If(props.Label != "", h.Label(props.Label)),
		input,
		h.Div(h.Class("text-red-500")),
	)

	return wrapped
}
