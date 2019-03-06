package main

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-archaius/sources/file-source"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/archaius/local/resource"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/

func main() {
	chassis.RegisterSchema("rest", &resource.Hello{})
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}
	//Add your custom file to archaius
	if err := archaius.AddFile("./conf/custom.yaml", archaius.WithFileHandler(filesource.UseFileNameAsKeyContentAsValue)); err != nil {
		openlogging.Error("add file configurations failed." + err.Error())
		return
	}
	//Add your custom file to archaius
	if err := archaius.AddFile("./conf/props.yaml"); err != nil {
		openlogging.Error("add props configurations failed." + err.Error())
		return
	}
	chassis.Run()
}
