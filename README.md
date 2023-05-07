# GOM

> 当前为测试版本，非稳定版本

GOM 提供了比较简单的对象假数据生成方式。
通过给定一个目标结构和相关描述，可以按照描述进行数据生成。
支持全部的 `JSON` 数据类型：

- 数字类型（`Number`）: 整形（`Integer`）、浮点型（`Float`）
- 字符串类型（`String`）
- 布尔类型（`Bool`）: `True`,`False`
- 数组类型（`Array`）
- 对象类型（`Map`）

并且提供了根据指定正则逆向为某个字符串的方法（正则表达式需要遵循 `Golang-RE2`
规范，有关规范请移步：[https://golang.org/s/re2syntax](https://golang.org/s/re2syntax)）。

## 1. 功能特性

- [x] 基于`Golang-RE2`的正则反解。
- [x] 根据对象描述生成指定结构。
- [ ] 控制台工具。
- [ ] 支持 WEB-API 调用，以及本地请求方法。
- [ ] 同版本种子回放特性。

## 2. 文档链接

### 2.1. 语言

- [中文文档](./README.md)
- [EN Doc](./docs/en/README.md)

### 2.2. 模块

- [正则反解生成器]()
- [结构体模拟数据生成器]()

## 3. 使用

### 3.1. Golang 引用

> 当前为测试版本，非稳定版本。

在控制台中执行以下指令：

```bash 
go get -u github.com/uberate/gom
```

在您的代码中使用：

```go
import "github.com/uberate/gom/pkg/obj_describe"
```

更多信息请移步：[example](./example)

## 4. 构建

有关构建请移步：[构建 GOM](./docs/zhcn/build.md)