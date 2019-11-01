### 开发常用函数工具箱(Develop a common function toolkit)

#### 安装(Install)

```shell
go get github.com/wumoxi/toolkit
```

#### 生成指定范围内的随机数

```go
r := toolbox.Random(100000, 999999)
```

#### 生成指定范围内的有序整型切片

```go
s := toolbox.GenerateSectionIntSliceOfOrderly(1, 20, 3)
```


#### 生成指定范围内的无序整型切片

```go
s := toolbox.GenerateSectionIntSliceOfDisorderly(1, 20)
```

