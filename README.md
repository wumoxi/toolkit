### 开发常用函数工具箱(Develop a common function toolkit)

#### 安装(Install)

```shell
go get github.com/wumoxi/toolkit
```

#### 生成指定范围内的随机数

> Call

```go
r := toolbox.Random(100000, 999999)
```

> Result

```text
// This is a random value
722479
```

#### 生成指定范围内的有序整型切片

> Call

```go
s := toolbox.GenerateSectionIntSliceOfOrderly(1, 20, 3)
```

> Result

```text
[1 4 7 10 13 16 19]
```


#### 生成指定范围内的无序整型切片

> Call

```go
s := toolbox.GenerateSectionIntSliceOfDisorderly(1, 20)
```

> Result

```text
// It could be something like this
[15 19 19 15 18]
```

#### 使用可变参数连接字符串切片

> Call

```go
s := toolbox.JoinItemOfStringSlice("", "中", "华", "人", "民", "共", "和", "国")
```

> Result

```text
中华人民共和国
```

