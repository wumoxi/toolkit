# Golang开发常用函数工具箱

## 安装

```shell
go get github.com/wumoxi/toolkit
```

## 随机数

### 生成指定范围内的随机数

```go
// r: 可能的随机值722479
r := toolkit.Random(100000, 999999)
```

## Slice

### 生成指定范围内的有序整型切片

```go
// s: [1 4 7 10 13 16 19]
s := toolkit.GenerateSectionIntSliceOfOrderly(1, 20, 3)
```

### 生成指定范围内的无序整型切片

```go
// s: 可能是这样的[4 18 16 7 8 4 19 1 16 9 13 9 1 4 16 15 13 8 14]
s := toolkit.GenerateSectionIntSliceOfDisorderly(1, 20)
```

### 使用可变参数连接字符串切片

```go
// s: 中华人民共和国
s := toolkit.JoinItemOfStringSlice("", "中", "华", "人", "民", "共", "和", "国")
```

### 生成指定范围内的斐波那契数列

```go
// s: [1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765]
s := toolkit.GenerateSectionFibSlice(20)
```

### 计算正整数阶乘

```go
// i: 362880
i := toolkit.Factorial(10)
```

### 生成指定范围内的正整数阶乘切片

```go
// s: [1 1 2 6 24 120 720 5040 40320 362880]
s := toolkit.FactorialSliceOfUint64(10)
```

## 时间

### 计算函数运行时间

可以使用 `toolkit.FunctionStart()` 和 `duration := toolkit.FunctionEnd()` 函数对待检测运行时间的函数进行包裹，`FunctionEnd()` 返回一个 `Duration` 值。

```go
func main() {
	toolkit.FunctionStart()
	total := func(number int64) (sum int64) {
		var i int64
		for i = 0; i < number; i++ {
			sum += i
		}
		return sum
	}(100000000)
	duration := toolkit.FunctionEnd()
	// 具体的运行时间根据执行环境而定
	// 102.202207ms
	fmt.Println("duration:", duration, ", total:", total)
}
```

## 运行时

### 获取当前程序占用内存大小

```go
// such as: info = 128.00KB
info := toolkit.MemAllocated()
```

## 反射

### 获取运行时函数名称

```go
// funcName = "fmt.Printf"
funcName, err := toolkit.GetFuncName(fmt.Printf)
if err != nil {
	panic(err)
}
```

### 获取结构体字段名和标签映射

```go
// student struct type
type student struct {
	Id      int    `json:"id" canChangeMethod:"-"`
	Name    string `json:"name" canChangeMethod:"change"`
	Email   string `json:"email" canChangeMethod:"change"`
	Phone   string `json:"phone" canChangeMethod:"change"`
	Sex     string `json:"sex" canChangeMethod:"change"`
	Age     int    `json:"age" canChangeMethod:"change,modify"`
	Address string `json:"address" canChangeMethod:"change,modify"`
}

// actual := map[string][]map[string]string{
//     "Id":      {{"json": "id"}},
//     "Name":    {{"json": "name"}, {"canChangeMethod": "change"}},
//     "Email":   {{"json": "email"}, {"canChangeMethod": "change"}},
//     "Phone":   {{"json": "phone"}, {"canChangeMethod": "change"}},
//     "Sex":     {{"json": "sex"}, {"canChangeMethod": "change"}},
//     "Age":     {{"json": "age"}, {"canChangeMethod": "change,modify"}},
//     "Address": {{"json": "address"}, {"canChangeMethod": "change,modify"}},
// }
actual, err := toolkit.GetAllStructTags(student{})
if err != nil {
	panic(err)
}
```

## 结构体

### 通过标签值获取字段名列表

```go
// student struct type
type student struct {
	Id      int    `json:"id" canChangeMethod:"-"`
	Name    string `json:"name" canChangeMethod:"change"`
	Email   string `json:"email" canChangeMethod:"change"`
	Phone   string `json:"phone" canChangeMethod:"change"`
	Sex     string `json:"sex" canChangeMethod:"change"`
	Age     int    `json:"age" canChangeMethod:"change,modify"`
	Address string `json:"address" canChangeMethod:"change,modify"`
}

// actual := [Address Age Email Name Phone Sex]
actual, err := toolkit.GetFieldsNameByTag(student{}, "change")
if err != nil {
	panic(err)
}

// actual = [Address Age]
actual, err = toolkit.GetFieldsNameByTag(student{}, "modify")
if err != nil {
	panic(err)
}

// actual = [Phone]
actual, err = toolkit.GetFieldsNameByTag(student{}, "phone")
if err != nil {
	panic(err)
}
```

## 类型判断

### 是否是布尔值

只有当值为布尔类型值时，其结果值为 `true`，否则为 `false`

```go
// r the value is true
r := toolkit.IsBool(false)

// r the value is true
r := toolkit.IsBool(true)
```

### 是否是float32类型值

只有当值为float32类型值时，其结果值为 `true`，否则为 `false`

```go
// r the value is true
r := toolkit.IsFloat32(3.1415926)
```

### 是否是float64类型值

只有当值为float64类型值时，其结果值为 `true`，否则为 `false`

```go
// r the value is true
r := toolkit.IsFloat64(3.1415926)
```

### 是否是uint8类型值

只有当值为uint8类型值时，其结果值为 `true`，否则为 `false`

```go
var v uint8 = 100

// r the value is true
r := toolkit.IsUint8(v)
```

### 是否是uint16类型值

只有当值为uint16类型值时，其结果值为 `true`，否则为 `false`

```go
var v uint16 = 100

// r the value is true
r := toolkit.IsUint16(v)
```


### 是否是uint32类型值

只有当值为uint32类型值时，其结果值为 `true`，否则为 `false`

```go
var v uint32 = 100

// r the value is true
r := toolkit.IsUint32(v)
```

### 是否是uint64类型值

只有当值为uint64类型值时，其结果值为 `true`，否则为 `false`

```go
var v uint64 = 100

// r the value is true
r := toolkit.IsUint64(v)
```

### 是否是uint类型值

只有当值为uint类型值时，其结果值为 `true`，否则为 `false`

```go
var v uint = 100

// r the value is true
r := toolkit.IsUint(v)
```

### 是否是int8类型值

只有当值为int8类型值时，其结果值为 `true`，否则为 `false`

```go
var v int8 = 100

// r the value is true
r := toolkit.IsInt8(v)
```

### 是否是int16类型值

只有当值为int16类型值时，其结果值为 `true`，否则为 `false`

```go
var v int16 = 100

// r the value is true
r := toolkit.IsInt16(v)
```

### 是否是int32类型值

只有当值为int32类型值时，其结果值为 `true`，否则为 `false`

```go
var v int32 = 100

// r the value is true
r := toolkit.IsInt32(v)
```

### 是否是int64类型值

只有当值为int64类型值时，其结果值为 `true`，否则为 `false`

```go
var v int64 = 100

// r the value is true
r := toolkit.IsInt64(v)
```


### 是否是int类型值

只有当值为int类型值时，其结果值为 `true`，否则为 `false`

```go
var v int = 100

// r the value is true
r := toolkit.IsInt(v)
```

### 是否是uintptr类型值

只有当值为uintptr类型值时，其结果值为 `true`，否则为 `false`

```go
var v uintptr = 7788556699

// r the value is true
r := toolkit.IsUintptr(v)
```

## 发送邮件

```go
mailer := toolkit.NewEmail(&toolkit.MailerParams{
    ServerHost:   "smtp.qq.com",
    ServerPort:   465,
    FromEmail:    "shaohua@foxmail.com",
    FromPassword: "mmooqssdsuwjskssddthpubddf",
    FromName:     "武沫汐",
    Toers:        []string{"warner@126.com", "warner@139.com", "contact.warner@gmail.com"},
    CCers:        []string{"warner@163.com"},
})

send, err := mailer.Send("Golang邮件发送", `中华人民共和国 - Golang邮件发送`, "text/plain")
if err != nil && !send {
    t.Fatal(err)
} else {
    t.Log("发送成功")
}
```

