package etc

import (
	xstr "strings"

	"github.com/bitwormhole/starter-tools/tools/configen/demo/car"
	"github.com/bitwormhole/starter/application"
)

func car1(t *car.Body, context application.RuntimeContext) error {

	// [component]
	// id=abc

	return nil
}

func car2(t *car.Body, context application.RuntimeContext) error {

	// [component]
	// id=body2
	// class=body

	return nil
}

func door(t *car.Door, context application.RuntimeContext) error {

	// [component]
	// id=door1
	// class=door

	return nil
}

func builder(t *xstr.Builder, context application.RuntimeContext) error {

	// [component]
	// id=strbuilder1
	// class=strBuilder

	return nil
}
