package main

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-archaius/source/util"
	"github.com/go-chassis/go-chassis-examples/archaius/resource"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/

func main() {
	chassis.RegisterSchema("rest", &resource.Hello{})
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	//Add your custom file to archaius
	if err := archaius.AddFile("./conf/custom.yaml", archaius.WithFileHandler(util.UseFileNameAsKeyContentAsValue)); err != nil {
		openlog.Error("add file configurations failed." + err.Error())
		return
	}
	//Add your custom file to archaius
	if err := archaius.AddFile("./conf/props.yaml"); err != nil {
		openlog.Error("add props configurations failed." + err.Error())
		return
	}
	chassis.Run()
}
