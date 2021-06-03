package configenx

import (
	"errors"
	"os"

	"github.com/bitwormhole/starter-tools/tools/configen1"
	"github.com/bitwormhole/starter-tools/tools/configen2"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

func Run(args []string) error {

	pwd, ok := os.LookupEnv("PWD")
	if !ok {
		return errors.New("no env: PWD")
	}

	dir := fs.Default().GetPath(pwd)
	propsFile := dir.GetChild("configen.properties")
	text, err := propsFile.GetIO().ReadText()
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	ver, err := props.GetPropertyRequired("configen.version")
	if err != nil {
		return err
	}

	if ver == "1" {
		return configen1.Run(args)
	} else if ver == "2" {
		return configen2.Run(args)
	}

	return errors.New("unsupported configen version:" + ver)
}
