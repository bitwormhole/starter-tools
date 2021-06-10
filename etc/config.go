package etc

import (
	"github.com/bitwormhole/starter-tools/cmd"
	"github.com/bitwormhole/starter-tools/tools/configen2/templates"
	"github.com/bitwormhole/starter-tools/tools/configenx"
	"github.com/bitwormhole/starter-tools/tools/help"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}

	cb.AddComponent(&config.ComInfo{
		ID: "commands",
		OnNew: func() lang.Object {
			return &cmd.CommandManager{}
		},
		OnInject: func(obj lang.Object, context application.Context) error {

			helpHandler := &help.CommandHandler{}
			helpHandler.Init(context)

			com := obj.(*cmd.CommandManager)
			com.SetDefaultHandler(helpHandler)
			return com.Inject(context)
		},
		OnInit: func(o lang.Object) error {
			com := o.(*cmd.CommandManager)
			return com.Init()
		},
	})

	cb.AddComponent(&config.ComInfo{
		ID: "cmd-about",
		OnNew: func() lang.Object {
			return &cmd.CommandRegistration{
				Name: "about",
			}
		},
		OnInject: func(obj lang.Object, context application.Context) error {
			reg := obj.(*cmd.CommandRegistration)
			reg.Handler = &help.CommandHandler{}
			return reg.Handler.Init(context)
		},
	})

	cb.AddComponent(&config.ComInfo{
		ID: "cmd-help",
		OnNew: func() lang.Object {
			return &cmd.CommandRegistration{
				Name: "help",
			}
		},
		OnInject: func(obj lang.Object, context application.Context) error {
			reg := obj.(*cmd.CommandRegistration)
			reg.Handler = &help.CommandHandler{}
			return reg.Handler.Init(context)
		},
	})

	cb.AddComponent(&config.ComInfo{
		ID: "cmd-version",
		OnNew: func() lang.Object {
			return &cmd.CommandRegistration{
				Name: "version",
			}
		},
		OnInject: func(obj lang.Object, context application.Context) error {
			reg := obj.(*cmd.CommandRegistration)
			reg.Handler = &help.CommandHandler{}
			return reg.Handler.Init(context)
		},
	})

	// configen1
	cominfobuilder.Reset()
	cominfobuilder.ID("configen1").Class("").Scope("singleton").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &cmd.CommandRegistration{
			Name:    "configen1",
			Handler: &configenx.CommandHandler{},
		}
	})
	cominfobuilder.OnInject(func(obj lang.Object, context application.Context) error {
		reg := obj.(*cmd.CommandRegistration)
		return reg.Handler.Init(context)
	})
	cominfobuilder.CreateTo(cb)

	// configen2
	cominfobuilder.Reset()
	cominfobuilder.ID("configen2").Class("").Scope("singleton").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &cmd.CommandRegistration{
			Name:    "configen2",
			Handler: &configenx.CommandHandler{},
		}
	})
	cominfobuilder.OnInject(func(obj lang.Object, context application.Context) error {
		reg := obj.(*cmd.CommandRegistration)
		return reg.Handler.Init(context)
	})
	cominfobuilder.CreateTo(cb)

	// configenX
	cominfobuilder.Reset()
	cominfobuilder.ID("configen-x").Class("").Scope("singleton").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &cmd.CommandRegistration{
			Name:    "configen",
			Handler: &configenx.CommandHandler{},
		}
	})
	cominfobuilder.OnInject(func(obj lang.Object, context application.Context) error {
		reg := obj.(*cmd.CommandRegistration)
		return reg.Handler.Init(context)
	})
	cominfobuilder.CreateTo(cb)

	// configen2main-template
	cominfobuilder.Reset()
	cominfobuilder.ID("configen2-main-template-factory").Class("").Scope("singleton").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &templates.MainTemplateFactory{}
	})
	cominfobuilder.CreateTo(cb)

	return nil
}
