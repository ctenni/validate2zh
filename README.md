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
        fmt.Println(err)
        // Name为必填字段
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

### `ValidateStructAll(s any) error`

Validates the struct and returns **all** Chinese error messages joined by `; `, or `nil` on success.

```go
err := v2zh.ValidateStructAll(user)
// → "Name为必填字段; Email必须是一个有效的邮箱; Age必须大于或等于18"
```

## 🏷️ Supported Validation Tags

### Required — 必填校验

| Tag | Description |
|-----|-------------|
| `required` | Field must not be zero value |
| `required_if` | Required if another field equals given value |
| `required_unless` | Required unless another field equals given value |
| `required_with` | Required if any of given fields are present |
| `required_with_all` | Required if all of given fields are present |
| `required_without` | Required if any of given fields are absent |
| `required_without_all` | Required if all of given fields are absent |

### Excluded — 排除校验

| Tag | Description |
|-----|-------------|
| `excluded_if` | Excluded if another field equals given value |
| `excluded_unless` | Excluded unless another field equals given value |
| `excluded_with` | Excluded if any of given fields are present |
| `excluded_with_all` | Excluded if all of given fields are present |
| `excluded_without` | Excluded if any of given fields are absent |
| `excluded_without_all` | Excluded if all of given fields are absent |

### Comparison — 比较校验

| Tag | Description |
|-----|-------------|
| `eq` | Equal to given value |
| `eq_ignore_case` | Equal to given value (case-insensitive) |
| `ne` | Not equal to given value |
| `ne_ignore_case` | Not equal to given value (case-insensitive) |
| `gt` | Greater than given value |
| `gte` | Greater than or equal to given value |
| `lt` | Less than given value |
| `lte` | Less than or equal to given value |

### Cross-Field — 跨字段比较

| Tag | Description |
|-----|-------------|
| `eqfield` | Equal to another field |
| `nefield` | Not equal to another field |
| `gtfield` | Greater than another field |
| `gtefield` | Greater than or equal to another field |
| `ltfield` | Less than another field |
| `ltefield` | Less than or equal to another field |
| `eqcsfield` | Equal to another field (case-sensitive) |
| `necsfield` | Not equal to another field (case-sensitive) |
| `gtcsfield` | Greater than another field (case-sensitive) |
| `gtecsfield` | Greater than or equal to another field (case-sensitive) |
| `ltcsfield` | Less than another field (case-sensitive) |
| `ltecsfield` | Less than or equal to another field (case-sensitive) |

### String — 字符串校验

| Tag | Description |
|-----|-------------|
| `len` | Exact length |
| `min` | Minimum length / value |
| `max` | Maximum length / value |
| `contains` | Contains substring |
| `containsany` | Contains any of given characters |
| `containsrune` | Contains given rune |
| `excludes` | Does not contain substring |
| `excludesall` | Does not contain any of given characters |
| `excludesrune` | Does not contain given rune |
| `startswith` | Starts with prefix |
| `endswith` | Ends with suffix |
| `startsnotwith` | Does not start with prefix |
| `endsnotwith` | Does not end with suffix |
| `fieldcontains` | Contains value of another field |
| `fieldexcludes` | Does not contain value of another field |

### Character — 字符组成

| Tag | Description |
|-----|-------------|
| `alpha` | Only letters (a-z, A-Z) |
| `alphaspace` | Letters and spaces |
| `alphanum` | Letters and digits |
| `alphanumspace` | Letters, digits and spaces |
| `alphaunicode` | Unicode letters |
| `alphanumunicode` | Unicode letters and digits |
| `ascii` | ASCII characters only |
| `printascii` | Printable ASCII characters |
| `multibyte` | Multi-byte characters |
| `lowercase` | Lowercase only |
| `uppercase` | Uppercase only |

### Numeric — 数值校验

| Tag | Description |
|-----|-------------|
| `number` | Integer string |
| `numeric` | Integer or float string |
| `boolean` | Boolean string ("true"/"false") |
| `hexadecimal` | Hexadecimal string |
| `hexcolor` | Hex color (#RRGGBB) |
| `rgb` | RGB color (rgb(r,g,b)) |
| `rgba` | RGBA color (rgba(r,g,b,a)) |
| `hsl` | HSL color (hsl(h,s,l)) |
| `hsla` | HSLA color (hsla(h,s,l,a)) |
| `cmyk` | CMYK color |
| `iscolor` | Any color format (alias for hexcolor\|rgb\|rgba\|hsl\|hsla\|cmyk) |

### Network — 网络格式

| Tag | Description |
|-----|-------------|
| `ip` | IP address (v4 or v6) |
| `ipv4` | IPv4 address |
| `ipv6` | IPv6 address |
| `cidr` | CIDR notation (v4 or v6) |
| `cidrv4` | CIDR v4 |
| `cidrv6` | CIDR v6 |
| `mac` | MAC address |
| `hostname` | Hostname |
| `hostname_rfc1123` | RFC 1123 hostname |
| `hostname_port` | Hostname:port |
| `fqdn` | Fully qualified domain name |
| `port` | Port number |
| `tcp4_addr` | TCP4 address |
| `tcp6_addr` | TCP6 address |
| `tcp_addr` | TCP address |
| `udp4_addr` | UDP4 address |
| `udp6_addr` | UDP6 address |
| `udp_addr` | UDP address |
| `ip4_addr` | IP4 address |
| `ip6_addr` | IP6 address |
| `ip_addr` | IP address |
| `unix_addr` | Unix address |
| `uds_exists` | Unix domain socket exists |

### Format — 格式校验

| Tag | Description |
|-----|-------------|
| `email` | Email address |
| `url` | URL |
| `http_url` | HTTP/HTTPS URL |
| `https_url` | HTTPS URL |
| `uri` | URI |
| `origin` | Origin |
| `urn_rfc2141` | URN (RFC 2141) |
| `base32` | Base32 string |
| `base64` | Base64 string |
| `base64url` | Base64 URL-safe string |
| `base64rawurl` | Base64 raw URL-safe string |
| `datauri` | Data URI |
| `json` | JSON string |
| `jwt` | JWT token |
| `html` | HTML string |
| `html_encoded` | HTML-encoded string |
| `url_encoded` | URL-encoded string |
| `datetime` | Datetime matching given format |

### Identifier — 标识符校验

| Tag | Description |
|-----|-------------|
| `uuid` | UUID |
| `uuid3` | UUID v3 |
| `uuid4` | UUID v4 |
| `uuid5` | UUID v5 |
| `uuid_rfc4122` | UUID RFC 4122 |
| `uuid3_rfc4122` | UUID v3 RFC 4122 |
| `uuid4_rfc4122` | UUID v4 RFC 4122 |
| `uuid5_rfc4122` | UUID v5 RFC 4122 |
| `ulid` | ULID |
| `isbn` | ISBN |
| `isbn10` | ISBN-10 |
| `isbn13` | ISBN-13 |
| `issn` | ISSN |
| `ssn` | SSN |
| `ein` | EIN (Employer Identification Number) |
| `cve` | CVE identifier |

### Hash — 哈希校验

| Tag | Description |
|-----|-------------|
| `md4` | MD4 hash |
| `md5` | MD5 hash |
| `sha256` | SHA-256 hash |
| `sha384` | SHA-384 hash |
| `sha512` | SHA-512 hash |
| `ripemd128` | RIPEMD-128 hash |
| `ripemd160` | RIPEMD-160 hash |
| `tiger128` | Tiger-128 hash |
| `tiger160` | Tiger-160 hash |
| `tiger192` | Tiger-192 hash |

### Enum — 枚举校验

| Tag | Description |
|-----|-------------|
| `oneof` | One of given values (case-sensitive) |
| `oneofci` 🔥 | One of given values (case-insensitive, custom) |
| `noneof` | None of given values (case-sensitive) |
| `noneofci` | None of given values (case-insensitive) |

### Location — 地理位置

| Tag | Description |
|-----|-------------|
| `latitude` | Latitude coordinate |
| `longitude` | Longitude coordinate |
| `timezone` | Timezone string |

### Country / Currency — 国家与货币

| Tag | Description |
|-----|-------------|
| `country_code` | Country code (alias for iso3166_1_alpha2\|iso3166_1_alpha3\|iso3166_1_alpha_numeric) |
| `iso3166_1_alpha2` | ISO 3166-1 alpha-2 |
| `iso3166_1_alpha3` | ISO 3166-1 alpha-3 |
| `iso3166_1_alpha_numeric` | ISO 3166-1 numeric |
| `iso3166_2` | ISO 3166-2 |
| `iso4217` | Currency code (ISO 4217) |
| `iso4217_numeric` | Currency numeric code (ISO 4217) |
| `postcode_iso3166_alpha2` | Postal code by country |
| `postcode_iso3166_alpha2_field` | Postal code by country from field |
| `bcp47_language_tag` | BCP 47 language tag |
| `bcp47_strict_language_tag` | BCP 47 strict language tag |
| `eu_country_code` | EU country code (alias) |
| `bic_iso_9362_2014` | BIC / SWIFT code |
| `bic` | BIC code (alias) |

### Crypto / Finance — 加密与金融

| Tag | Description |
|-----|-------------|
| `eth_addr` | Ethereum address |
| `eth_addr_checksum` | Ethereum address (checksummed) |
| `btc_addr` | Bitcoin address |
| `btc_addr_bech32` | Bitcoin Bech32 address |
| `credit_card` | Credit card number (Luhn) |
| `luhn_checksum` | Luhn checksum |
| `e164` | E.164 phone number |
| `phone` | Phone number |

### File / Directory — 文件与目录

| Tag | Description |
|-----|-------------|
| `file` | File path |
| `filepath` | File path (valid format) |
| `dir` | Directory path |
| `dirpath` | Directory path (valid format) |
| `image` | Image file |

### Other — 其他

| Tag | Description |
|-----|-------------|
| `mimetype` | MIME type |
| `unique` | Unique values in slice/map |
| `semver` | Semantic version |
| `cron` | Cron expression |
| `mongodb` | MongoDB ObjectID |
| `mongodb_connection_string` | MongoDB connection string |
| `dns_rfc1035_label` | DNS label (RFC 1035) |
| `spicedb` | SpiceDB |
| `isdefault` | Is default value |
| `skip_unless` | Skip validation unless condition met |
| `validateFn` | Custom validation function |

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
        fmt.Println(err)
        // Name为必填字段
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

### `ValidateStructAll(s any) error`

校验结构体，返回**所有**中文错误信息（`; ` 拼接），校验通过返回 `nil`。

```go
err := v2zh.ValidateStructAll(user)
// → "Name为必填字段; Email必须是一个有效的邮箱; Age必须大于或等于18"
```

## 🏷️ 支持的校验规则

### required 家族 — 必填校验

| 标签 | 说明 |
|------|------|
| `required` | 字段不能为零值 |
| `required_if` | 当另一个字段等于指定值时必填 |
| `required_unless` | 除非另一个字段等于指定值，否则必填 |
| `required_with` | 当任意指定字段存在时必填 |
| `required_with_all` | 当所有指定字段都存在时必填 |
| `required_without` | 当任意指定字段不存在时必填 |
| `required_without_all` | 当所有指定字段都不存在时必填 |

### excluded 家族 — 排除校验

| 标签 | 说明 |
|------|------|
| `excluded_if` | 当另一个字段等于指定值时排除 |
| `excluded_unless` | 除非另一个字段等于指定值，否则排除 |
| `excluded_with` | 当任意指定字段存在时排除 |
| `excluded_with_all` | 当所有指定字段都存在时排除 |
| `excluded_without` | 当任意指定字段不存在时排除 |
| `excluded_without_all` | 当所有指定字段都不存在时排除 |

### 比较校验

| 标签 | 说明 |
|------|------|
| `eq` | 等于指定值 |
| `eq_ignore_case` | 等于指定值（忽略大小写） |
| `ne` | 不等于指定值 |
| `ne_ignore_case` | 不等于指定值（忽略大小写） |
| `gt` | 大于指定值 |
| `gte` | 大于等于指定值 |
| `lt` | 小于指定值 |
| `lte` | 小于等于指定值 |

### 跨字段比较

| 标签 | 说明 |
|------|------|
| `eqfield` | 等于另一个字段 |
| `nefield` | 不等于另一个字段 |
| `gtfield` | 大于另一个字段 |
| `gtefield` | 大于等于另一个字段 |
| `ltfield` | 小于另一个字段 |
| `ltefield` | 小于等于另一个字段 |
| `eqcsfield` | 等于另一个字段（大小写敏感） |
| `necsfield` | 不等于另一个字段（大小写敏感） |
| `gtcsfield` | 大于另一个字段（大小写敏感） |
| `gtecsfield` | 大于等于另一个字段（大小写敏感） |
| `ltcsfield` | 小于另一个字段（大小写敏感） |
| `ltecsfield` | 小于等于另一个字段（大小写敏感） |

### 字符串校验

| 标签 | 说明 |
|------|------|
| `len` | 精确长度 |
| `min` | 最小长度/值 |
| `max` | 最大长度/值 |
| `contains` | 包含子串 |
| `containsany` | 包含任意字符 |
| `containsrune` | 包含指定字符 |
| `excludes` | 不包含子串 |
| `excludesall` | 不包含任意字符 |
| `excludesrune` | 不包含指定字符 |
| `startswith` | 以指定前缀开头 |
| `endswith` | 以指定后缀结尾 |
| `startsnotwith` | 不以指定前缀开头 |
| `endsnotwith` | 不以指定后缀结尾 |
| `fieldcontains` | 包含另一个字段的值 |
| `fieldexcludes` | 不包含另一个字段的值 |

### 字符组成

| 标签 | 说明 |
|------|------|
| `alpha` | 仅字母 |
| `alphaspace` | 字母和空格 |
| `alphanum` | 字母和数字 |
| `alphanumspace` | 字母、数字和空格 |
| `alphaunicode` | Unicode 字母 |
| `alphanumunicode` | Unicode 字母和数字 |
| `ascii` | 仅 ASCII 字符 |
| `printascii` | 可打印 ASCII 字符 |
| `multibyte` | 多字节字符 |
| `lowercase` | 仅小写 |
| `uppercase` | 仅大写 |

### 数值校验

| 标签 | 说明 |
|------|------|
| `number` | 整数字符串 |
| `numeric` | 整数或浮点数字符串 |
| `boolean` | 布尔字符串（"true"/"false"） |
| `hexadecimal` | 十六进制字符串 |
| `hexcolor` | 十六进制颜色（#RRGGBB） |
| `rgb` | RGB 颜色 |
| `rgba` | RGBA 颜色 |
| `hsl` | HSL 颜色 |
| `hsla` | HSLA 颜色 |
| `cmyk` | CMYK 颜色 |
| `iscolor` | 任意颜色格式（hexcolor\|rgb\|rgba\|hsl\|hsla\|cmyk） |

### 网络格式

| 标签 | 说明 |
|------|------|
| `ip` | IP 地址（v4 或 v6） |
| `ipv4` | IPv4 地址 |
| `ipv6` | IPv6 地址 |
| `cidr` | CIDR（v4 或 v6） |
| `cidrv4` | CIDR v4 |
| `cidrv6` | CIDR v6 |
| `mac` | MAC 地址 |
| `hostname` | 主机名 |
| `hostname_rfc1123` | RFC 1123 主机名 |
| `hostname_port` | 主机名:端口 |
| `fqdn` | 完整域名 |
| `port` | 端口号 |
| `tcp4_addr` | TCP4 地址 |
| `tcp6_addr` | TCP6 地址 |
| `tcp_addr` | TCP 地址 |
| `udp4_addr` | UDP4 地址 |
| `udp6_addr` | UDP6 地址 |
| `udp_addr` | UDP 地址 |
| `ip4_addr` | IP4 地址 |
| `ip6_addr` | IP6 地址 |
| `ip_addr` | IP 地址 |
| `unix_addr` | Unix 地址 |
| `uds_exists` | Unix 域套接字存在 |

### 格式校验

| 标签 | 说明 |
|------|------|
| `email` | 邮箱地址 |
| `url` | URL |
| `http_url` | HTTP/HTTPS URL |
| `https_url` | HTTPS URL |
| `uri` | URI |
| `origin` | Origin |
| `urn_rfc2141` | URN（RFC 2141） |
| `base32` | Base32 字符串 |
| `base64` | Base64 字符串 |
| `base64url` | Base64 URL 安全字符串 |
| `base64rawurl` | Base64 原始 URL 安全字符串 |
| `datauri` | Data URI |
| `json` | JSON 字符串 |
| `jwt` | JWT 令牌 |
| `html` | HTML 字符串 |
| `html_encoded` | HTML 编码字符串 |
| `url_encoded` | URL 编码字符串 |
| `datetime` | 日期时间格式校验 |

### 标识符校验

| 标签 | 说明 |
|------|------|
| `uuid` | UUID |
| `uuid3` | UUID v3 |
| `uuid4` | UUID v4 |
| `uuid5` | UUID v5 |
| `uuid_rfc4122` | UUID RFC 4122 |
| `uuid3_rfc4122` | UUID v3 RFC 4122 |
| `uuid4_rfc4122` | UUID v4 RFC 4122 |
| `uuid5_rfc4122` | UUID v5 RFC 4122 |
| `ulid` | ULID |
| `isbn` | ISBN |
| `isbn10` | ISBN-10 |
| `isbn13` | ISBN-13 |
| `issn` | ISSN |
| `ssn` | 社会安全号码 |
| `ein` | 雇主识别号 |
| `cve` | CVE 标识符 |

### 哈希校验

| 标签 | 说明 |
|------|------|
| `md4` | MD4 哈希 |
| `md5` | MD5 哈希 |
| `sha256` | SHA-256 哈希 |
| `sha384` | SHA-384 哈希 |
| `sha512` | SHA-512 哈希 |
| `ripemd128` | RIPEMD-128 哈希 |
| `ripemd160` | RIPEMD-160 哈希 |
| `tiger128` | Tiger-128 哈希 |
| `tiger160` | Tiger-160 哈希 |
| `tiger192` | Tiger-192 哈希 |

### 枚举校验

| 标签 | 说明 |
|------|------|
| `oneof` | 枚举值之一（大小写敏感） |
| `oneofci` 🔥 | 枚举值之一（忽略大小写，自定义） |
| `noneof` | 不是枚举值之一（大小写敏感） |
| `noneofci` | 不是枚举值之一（忽略大小写） |

### 地理位置

| 标签 | 说明 |
|------|------|
| `latitude` | 纬度 |
| `longitude` | 经度 |
| `timezone` | 时区 |

### 国家与货币

| 标签 | 说明 |
|------|------|
| `country_code` | 国家代码（别名） |
| `iso3166_1_alpha2` | ISO 3166-1 alpha-2 |
| `iso3166_1_alpha3` | ISO 3166-1 alpha-3 |
| `iso3166_1_alpha_numeric` | ISO 3166-1 数字代码 |
| `iso3166_2` | ISO 3166-2 |
| `iso4217` | 货币代码（ISO 4217） |
| `iso4217_numeric` | 货币数字代码 |
| `postcode_iso3166_alpha2` | 邮政编码（按国家） |
| `postcode_iso3166_alpha2_field` | 邮政编码（从字段读取国家） |
| `bcp47_language_tag` | BCP 47 语言标签 |
| `bcp47_strict_language_tag` | BCP 47 严格语言标签 |
| `eu_country_code` | 欧盟国家代码（别名） |
| `bic_iso_9362_2014` | BIC/SWIFT 代码 |
| `bic` | BIC 代码（别名） |

### 加密与金融

| 标签 | 说明 |
|------|------|
| `eth_addr` | 以太坊地址 |
| `eth_addr_checksum` | 以太坊地址（校验和） |
| `btc_addr` | 比特币地址 |
| `btc_addr_bech32` | 比特币 Bech32 地址 |
| `credit_card` | 信用卡号（Luhn 算法） |
| `luhn_checksum` | Luhn 校验和 |
| `e164` | E.164 电话号码 |
| `phone` | 电话号码 |

### 文件与目录

| 标签 | 说明 |
|------|------|
| `file` | 文件路径 |
| `filepath` | 文件路径格式 |
| `dir` | 目录路径 |
| `dirpath` | 目录路径格式 |
| `image` | 图片文件 |

### 其他

| 标签 | 说明 |
|------|------|
| `mimetype` | MIME 类型 |
| `unique` | 切片/Map 中唯一值 |
| `semver` | 语义化版本 |
| `cron` | Cron 表达式 |
| `mongodb` | MongoDB ObjectID |
| `mongodb_connection_string` | MongoDB 连接字符串 |
| `dns_rfc1035_label` | DNS 标签（RFC 1035） |
| `spicedb` | SpiceDB |
| `isdefault` | 是否为零值 |
| `skip_unless` | 条件不满足时跳过校验 |

> `oneofci` 是自定义校验器，支持大小写不敏感的枚举匹配。

## 🧪 测试

```bash
go test -v ./v2zh/...
```

## 📄 License

[MIT](LICENSE)
