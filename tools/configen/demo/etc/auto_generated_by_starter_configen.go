// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	car_63674827 "github.com/bitwormhole/starter-tools/tools/configen/demo/car"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	strings_873be872 "strings"
)

func Config(cb application.ConfigBuilder) error {

    // builder
    cb.AddComponent(&config.ComInfo{
		ID: "strbuilder1",
		Class: "strBuilder",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &strings_873be872.Builder{}
		},
    })

    // builder3
    cb.AddComponent(&config.ComInfo{
		ID: "builder3",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{"a1","a2","a3"},
		OnNew: func() lang.Object {
		    return &strings_873be872.Builder{}
		},
    })

    // car1
    cb.AddComponent(&config.ComInfo{
		ID: "abc",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &car_63674827.Body{}
		},
    })

    // car11
    cb.AddComponent(&config.ComInfo{
		ID: "car11",
		Class: "abc",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &car_63674827.Body{}
		},
    })

    // car2
    cb.AddComponent(&config.ComInfo{
		ID: "body2",
		Class: "body",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &car_63674827.Body{}
		},
    })

    // car22
    cb.AddComponent(&config.ComInfo{
		ID: "body2",
		Class: "body",
		Scope: application.ScopeSingleton,
		Aliases: []string{"c22","c33","c44","c50","c666"},
		OnNew: func() lang.Object {
		    return &car_63674827.Body{}
		},
		OnInit: func(obj lang.Object) error {
		    target := obj.(*car_63674827.Body)
		    return target.Start()
		},
		OnDestroy: func(obj lang.Object) error {
		    target := obj.(*car_63674827.Body)
		    return target.Stop()
		},
    })

    // door
    cb.AddComponent(&config.ComInfo{
		ID: "door1",
		Class: "door",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &car_63674827.Door{}
		},
    })

    return nil
}
