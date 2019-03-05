This demo show how to mange your custom configurations

it include 2 custom files.

# custom.yaml
this file's name is considered as key and content of file is considered as value
```go
archaius.AddFile("./conf/custom.yaml", 
	archaius.WithFileHandler(filesource.UseFileNameAsKeyContentAsValue))
```

# props.yaml
each of key value in file is considered as key value
```go
archaius.AddFile("./conf/props.yaml")
```

See results
```sh
go build main.go
./main
```

call http://127.0.0.1:5001/file to see custom.yaml

call http://127.0.0.1:5001/props/{key} to see props.yaml value

you can change file content of any file in runtime, 
and call API to see that config has been changed in runtime.

NOTICE: In distributed system you can change key value in a config server 
to manipulate configurations in runtime, 
just like you change value by changing file content. 
go-archaius use config client to interact with config server. 
Check https://github.com/go-chassis/go-cc-client to see config center client implementation
