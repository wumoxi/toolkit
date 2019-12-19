# Golang开发常用函数工具箱

## 安装

```shell
go get github.com/wumoxi/toolkit
```

## 随机数

### 生成指定范围内的随机数

```go
// r: 可能的随机值722479
r := toolbox.Random(100000, 999999)
```

## Slice

### 生成指定范围内的有序整型切片

```go
// s: [1 4 7 10 13 16 19]
s := toolbox.GenerateSectionIntSliceOfOrderly(1, 20, 3)
```

### 生成指定范围内的无序整型切片

```go
// s: 可能是这样的[4 18 16 7 8 4 19 1 16 9 13 9 1 4 16 15 13 8 14]
s := toolbox.GenerateSectionIntSliceOfDisorderly(1, 20)
```

### 使用可变参数连接字符串切片

```go
// s: 中华人民共和国
s := toolbox.JoinItemOfStringSlice("", "中", "华", "人", "民", "共", "和", "国")
```

### 生成指定范围内的斐波那契数列

```go
// s: [1 1 2 3 5 8 13 21 34 55 89]
s := toolbox.FibSliceOfInt(100)
```

### 计算正整数阶乘

```go
// i: 362880
i := toolbox.Factorial(10)
```

### 生成指定范围内的正整数阶乘切片

```go
// s: [1 1 2 6 24 120 720 5040 40320 362880]
s := toolbox.FactorialSliceOfUint64(10)
```

## 时间

### 计算函数运行时间

可以使用 `toolkit.FunctionStart()` 和 `duration := toolkit.FunctionEnd()` 函数对待检测运行时间的函数进行包裹，`FunctionEnd()` 返回一个 `Duration` 值。

```go
func main() {
	toolkit.FunctionStart()
	sum(number int64) (sum int64) {
    	var i int64
    	for i = 0; i < number; i++ {
    		sum += i
    	}
    	return sum
    }
	duration := toolkit.FunctionEnd()
	// 具体的运行时间根据执行环境而定
	// 102.202207ms
	fmt.Println(duration)
}
```

## 结构体

### 获取结构字段名和标记映射

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

### 是从标签值中获取字段名的数组

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

// actual: ["Address", "Age", "Email", "Name", "Phone", "Sex"] 
actual, err := GetFieldNameByTagValue(student{}, "change")
if err != nil {
    t.Errorf("get field name slice by tag value error: %s\n", err)
}

// actual: ["Address", "Age"] 
actual, err = GetFieldNameByTagValue(student{}, "modify")
if err != nil {
    t.Errorf("get field name slice by tag value error: %s\n", err)
}
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

