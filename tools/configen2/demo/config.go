package demo

import (
	xstr "strings"

	"github.com/bitwormhole/starter-tools/tools/configen1/demo/car"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
)

type car1 struct {
	markup.Component

	instance *car.Body `as:"instance" id:"abc"    class:"body" scope:"singleton" `

	LeftDoor  *car.Door `inject:"#door1" to:"LeftDoor"`
	RightDoor *car.Door `inject:"#door2"`
	BackDoor  *car.Door `inject:"#door3"`
}

type car2 struct {
	markup.Component

	instance *car.Body
}

type door struct {
	markup.Component

	instance *car.Door
}

type builder struct {
	markup.Component `id:"" class:""`

	// 小写开头的是特殊字段
	instance *xstr.Builder `initMethod:"start" destroyMethod:"stop" injectMethod:"config"`
	context  application.Context

	num int

	// 大写开头的是直接注入字段
	Env collection.Environment
}

func (inst *builder) inject() error {

	return nil
}
