## IO

go 语言的 fmt 跟 C 标准库中的 `stdio` 作用是差不多的，甚至命名都很相似。比如 [Scanning](https://pkg.go.dev/fmt#hdr-Scanning)：

> An analogous set of functions scans formatted text to yield values. Scan, Scanf and Scanln read from os.Stdin; Fscan, Fscanf and Fscanln read from a specified io.Reader; Sscan, Sscanf and Sscanln read from an argument string.  
Scan, Fscan, Sscan treat newlines in the input as spaces.

如果我们想要做到 Python 中 `input` 函数的效果，即在检测到换行符时候停止 Scanning，那么我们应该使用的是 `fmt.Scanln`，就像这样：

```go
	var name string
	fmt.Scanln(&name)
```

