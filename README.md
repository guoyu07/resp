# RESP

redis协议解析类库，并支持redis inline command。

## Parse/Encode Redis command

redis的command分为两种，inline command与bulkString。

在io.Reader中读取一个命令：

```
//ex: parse command
body := []byte("setex name 10 walu\r\n") //其实inline command对最后的\r\n没有要求
r := bytes.NewReader(body)

cmd, err := ReadCommand(r)

//so:
//cmd.Name() == "setex"
//cmd.String(1) == "name"
//cmd.Integer(2) == 10 (int64)
//cmd.String(3)  == "walu"

//----------------------------
//encode command
cmd := NewCommand("setex", "name", "10", "walu")
body := cmd.Format()
fmt.Println(body)
```

## Parse/Encode Redis Data

redis的通讯协议resp支持五种数据类型：SimpleString、Error、BulkString、Interger、Array


resp使用了Data结构体来表示这五种数据
```
type Data struct {
	T byte 		//表示类型
	String []byte	//SimpleString、Error、BulkString 使用这个属性取值
	Integer int64	//Interger使用这个属性取值
	Array []*Data	//Array使用这个属性取值
	IsNil bool	//
}
```

在io.Reader中读取一个Data：

```
//ex: parse command
body := []byte("+I'm simple string\r\n")
r := bytes.NewReader(body)

data, err := ReadCommand(r)
```