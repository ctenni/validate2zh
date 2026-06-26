<div align="center">

# validate2zh

**Go struct validation · Chinese error messages**

[![Go Reference](https://pkg.go.dev/badge/github.com/ctenni/validate2zh.svg)](https://pkg.go.dev/github.com/ctenni/validate2zh)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ctenni/validate2zh)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/ctenni/validate2zh)](LICENSE)
[![Tag](https://img.shields.io/github/v/tag/ctenni/validate2zh)](https://github.com/ctenni/validate2zh/tags)

A wrapper around [`go-playground/validator`](https://github.com/go-playground/validator)  
that automatically registers Chinese translations, returning validation errors in **Chinese**.

</div>

---

## 📦 Installation

```bash
go get github.com/ctenni/validate2zh@latest
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/ctenni/validate2zh/v2zh"
)

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"required,min=1,max=150"`
}

func main() {
    user := User{}
    err := v2zh.ValidateStruct(user)
    if err != nil {
        fmt.Println(err) // Name为必填字段
    }
}
```

## 📖 API

### `ValidateStruct(s any) error`

Validates the struct and returns the **first** Chinese error message, or `nil` on success.

```go
err := v2zh.ValidateStruct(user)
// → "Name为必填字段"
```

### `ValidateStructAll(s any) []error`

Validates the struct and returns **all** Chinese error messages, or `nil` on success.

```go
errs := v2zh.ValidateStructAll(user)
// → ["Name为必填字段", "Email必须是一个有效的邮箱", "Age必须大于或等于18"]
```

## 🏷️ Supported Validation Tags

| Category | Tags | Description |
|----------|------|-------------|
| Required | `required`, `required_if`, `required_unless`, `required_with`, `required_without` etc. | Required family |
| Range | `min`, `max`, `len`, `gt`, `gte`, `lt`, `lte` | Numeric & length |
| Enum | `oneof`, `oneofci` 🔥 | Enum validation; `oneofci` is case-insensitive |
| Comparison | `eqfield`, `nefield` | Cross-field comparison |
| Format | `email`, `url`, `uuid`, `ip`, `ipv4`, `mac`, `datetime` | String format |
| Numeric | `boolean`, `number`, `numeric` | Numeric types |
| Character | `alpha`, `alphanum` | Character composition |

> `oneofci` is a custom validator that supports case-insensitive enum matching.

## 🧪 Testing

```bash
go test -v ./v2zh/...
```

## 📄 License

[MIT](LICENSE)

---

<div align="center">

# validate2zh

**Go 结构体校验 · 中文错误提示**

基于 [`go-playground/validator`](https://github.com/go-playground/validator) 封装，  
自动注册中文翻译器，让校验错误信息直接返回**中文友好提示**。

</div>

---

## 📦 安装

```bash
go get github.com/ctenni/validate2zh@latest
```

## 🚀 快速开始

```go
package main

import (
    "fmt"
    "github.com/ctenni/validate2zh/v2zh"
)

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"required,min=1,max=150"`
}

func main() {
    user := User{}
    err := v2zh.ValidateStruct(user)
    if err != nil {
        fmt.Println(err) // Name为必填字段
    }
}
```

## 📖 API

### `ValidateStruct(s any) error`

校验结构体，返回**首条**中文错误信息。校验通过返回 `nil`。

```go
err := v2zh.ValidateStruct(user)
// → "Name为必填字段"
```

### `ValidateStructAll(s any) []error`

校验结构体，返回**所有**中文错误信息。校验通过返回 `nil`。

```go
errs := v2zh.ValidateStructAll(user)
// → ["Name为必填字段", "Email必须是一个有效的邮箱", "Age必须大于或等于18"]
```

## 🏷️ 支持的校验规则

| 分类 | 标签 | 说明 |
|------|------|------|
| 必填 | `required` / `required_if` / `required_unless` / `required_with` / `required_without` 等 | required 家族 |
| 范围 | `min` / `max` / `len` / `gt` / `gte` / `lt` / `lte` | 数值与长度 |
| 枚举 | `oneof` / `oneofci` 🔥 | 枚举值校验，`oneofci` 忽略大小写 |
| 比较 | `eqfield` / `nefield` | 跨字段比较 |
| 格式 | `email` / `url` / `uuid` / `ip` / `ipv4` / `mac` / `datetime` | 字符串格式 |
| 数值 | `boolean` / `number` / `numeric` | 数值类型 |
| 字符 | `alpha` / `alphanum` | 字符组成 |

> `oneofci` 是自定义校验器，支持大小写不敏感的枚举匹配。

## 🧪 测试

```bash
go test -v ./v2zh/...
```

## 📄 License

[MIT](LICENSE)
