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
[4 18 16 7 8 4 19 1 16 9 13 9 1 4 16 15 13 8 14]
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

#### 生成指定范围内的斐波那契数列

> Call

```go
s := toolbox.FibSliceOfInt(100)
```

> Result

```text
[1 1 2 3 5 8 13 21 34 55 89]
```

#### 计算正整数阶乘

> Call

```go
s := toolbox.Factorial(10)
```

> Result

```text
362880
```

#### 生成指定范围内的正整数阶乘切片

> Call

```go
s := toolbox.FactorialSliceOfUint64(10)
```

> Result

```text
[1 1 2 6 24 120 720 5040 40320 362880]
```

#### 计算函数运行时间

可以使用 `toolkit.FunctionStart()` 和 `duration := toolkit.FunctionEnd()` 函数对待检测运行时间的函数进行包裹，`FunctionEnd()` 返回一个 `Duration` 值。

> Call

```go
func main() {
	toolkit.FunctionStart()
	sum(100000000)
	duration := toolkit.FunctionEnd()
	fmt.Println(duration)
}

func sum(number int64) (sum int64) {
	var i int64
	for i = 0; i < number; i++ {
		sum += i
	}
	return sum
}
```

> Result

```text
// 具体的运行时间根据执行环境而定
// The specific runtime depends on the execution environment
102.202207ms
```

