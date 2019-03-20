This demo show how to mange your custom configurations for archaius

archaius support json file and key: value to mange you custom . Notice here you customized
keys are not GO-Chassis governance

## json file
 JSON file:  batch mange you custom configurations

```json
{
    "a.user":"peter",
    "its.pwd":"123"
}
```

## key:value
use key:value  one-by-one mange you custom configurations
```text
"a.user":"peter"
```
## Archaius

how to use archaius for huawei cloud

### json

* Click the button `Import`
![]( img/cse-1-json.png)

* choose json file and then click the button `Upload` to upload file
![]( img/cse-2-json.png)

* uploaded json file successfully , you can see config
![]( img/cse-3-json.png)
* call the api `/props/its.pwd` will reply `123`

### key:value

* Click the button `Create Configuration`
![]( img/cse-1-kv.png)

* input `key` any `value`
![]( img/cse-2-kv.png)

* set config successfully , you can see this
![]( img/cse-3-kv.png)

* call api `/props/c.name` success will reply `test-name`