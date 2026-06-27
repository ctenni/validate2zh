<div align="center">

# validate2zh

**Go 结构体校验 · 中文错误提示**

[![Go Reference](https://pkg.go.dev/badge/github.com/ctenni/validate2zh.svg)](https://pkg.go.dev/github.com/ctenni/validate2zh)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ctenni/validate2zh)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/ctenni/validate2zh)](LICENSE)
[![Tag](https://img.shields.io/github/v/tag/ctenni/validate2zh)](https://github.com/ctenni/validate2zh/tags)

</div>

基于 [`go-playground/validator`](https://github.com/go-playground/validator) 封装，  
自动注册中文翻译器，让校验错误信息直接返回**中文友好提示**。
多个字段同时校验，给出全面的错误提示。

> 🇬🇧 English version: [README.md](README.md)

---

## 📦 安装

```bash
go get github.com/ctenni/validate2zh@latest
```

## 🚀 快速开始

```go

import (
"testing"

"github.com/ctenni/validate2zh/v2zh"
)

type RequiredTest1 struct {
Name string `validate:"required"`
}

type AllPassTest struct {
Name  string `validate:"required"`
Email string `validate:"required,email"`
Age   int    `validate:"required,min=1,max=150"`
}


func Test_require1(t *testing.T) {
rt := &RequiredTest1{}
err := v2zh.ValidateStruct(rt)
t.Error(err)

err = v2zh.ValidateStructAll(AllPassTest{
Email: "invalid_email",
Age:   151,
})
t.Error(err)
}

```
- 输出
```markdown
    v2zh_test.go:22: Name为必填字段
    v2zh_test.go:28: Name为必填字段; Email必须是一个有效的邮箱; Age必须小于或等于150
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

<details>
<summary>required 家族 — 必填校验</summary>

| 标签 | 说明 |
|------|------|
| `required` | 字段不能为零值 |
| `required_if` | 当另一个字段等于指定值时必填 |
| `required_unless` | 除非另一个字段等于指定值，否则必填 |
| `required_with` | 当任意指定字段存在时必填 |
| `required_with_all` | 当所有指定字段都存在时必填 |
| `required_without` | 当任意指定字段不存在时必填 |
| `required_without_all` | 当所有指定字段都不存在时必填 |
</details>

<details>
<summary>excluded 家族 — 排除校验</summary>

| 标签 | 说明 |
|------|------|
| `excluded_if` | 当另一个字段等于指定值时排除 |
| `excluded_unless` | 除非另一个字段等于指定值，否则排除 |
| `excluded_with` | 当任意指定字段存在时排除 |
| `excluded_with_all` | 当所有指定字段都存在时排除 |
| `excluded_without` | 当任意指定字段不存在时排除 |
| `excluded_without_all` | 当所有指定字段都不存在时排除 |
</details>

<details>
<summary>比较校验</summary>

| 标签 | 说明 |
|------|------|
| `eq` | 等于指定值 |
| `eq_ignore_case` | 等于指定值（忽略大小写） |
| `ne` | 不等于指定值 |
| `ne_ignore_case` | 不等于指定值（忽略大小写） |
| `gt` | 大于 |
| `gte` | 大于等于 |
| `lt` | 小于 |
| `lte` | 小于等于 |
</details>

<details>
<summary>跨字段比较</summary>

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
</details>

<details>
<summary>字符串校验</summary>

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
</details>

<details>
<summary>字符组成</summary>

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
</details>

<details>
<summary>数值校验</summary>

| 标签 | 说明 |
|------|------|
| `number` | 整数字符串 |
| `numeric` | 整数或浮点数字符串 |
| `boolean` | 布尔字符串 |
| `hexadecimal` | 十六进制字符串 |
| `hexcolor` | 十六进制颜色 |
| `rgb` | RGB 颜色 |
| `rgba` | RGBA 颜色 |
| `hsl` | HSL 颜色 |
| `hsla` | HSLA 颜色 |
| `cmyk` | CMYK 颜色 |
| `iscolor` | 任意颜色格式（别名） |
</details>

<details>
<summary>网络格式</summary>

| 标签 | 说明 |
|------|------|
| `ip` | IP 地址（v4 或 v6） |
| `ipv4` | IPv4 地址 |
| `ipv6` | IPv6 地址 |
| `cidr` | CIDR 表示法 |
| `cidrv4` | CIDR v4 |
| `cidrv6` | CIDR v6 |
| `mac` | MAC 地址 |
| `hostname` | 主机名 |
| `hostname_rfc1123` | RFC 1123 主机名 |
| `hostname_port` | 主机名:端口 |
| `fqdn` | 完整域名 |
| `port` | 端口号 |
| `tcp4_addr` / `tcp6_addr` / `tcp_addr` | TCP 地址 |
| `udp4_addr` / `udp6_addr` / `udp_addr` | UDP 地址 |
| `ip4_addr` / `ip6_addr` / `ip_addr` | IP 地址 |
| `unix_addr` | Unix 地址 |
| `uds_exists` | Unix 域套接字存在 |
</details>

<details>
<summary>格式校验</summary>

| 标签 | 说明 |
|------|------|
| `email` | 邮箱地址 |
| `url` | URL |
| `http_url` | HTTP/HTTPS URL |
| `https_url` | HTTPS URL |
| `uri` | URI |
| `origin` | Origin |
| `urn_rfc2141` | URN |
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
</details>

<details>
<summary>标识符校验 — UUID / ISBN / SSN</summary>

| 标签 | 说明 |
|------|------|
| `uuid` / `uuid3` / `uuid4` / `uuid5` | UUID |
| `uuid_rfc4122` / `uuid3_rfc4122` / `uuid4_rfc4122` / `uuid5_rfc4122` | UUID RFC 4122 |
| `ulid` | ULID |
| `isbn` / `isbn10` / `isbn13` | ISBN |
| `issn` | ISSN |
| `ssn` | 社会安全号码 |
| `ein` | 雇主识别号 |
| `cve` | CVE 标识符 |
</details>

<details>
<summary>哈希校验</summary>

| 标签 | 说明 |
|------|------|
| `md4` / `md5` | MD |
| `sha256` / `sha384` / `sha512` | SHA |
| `ripemd128` / `ripemd160` | RIPEMD |
| `tiger128` / `tiger160` / `tiger192` | TIGER |
</details>

<details>
<summary>枚举校验</summary>

| 标签 | 说明 |
|------|------|
| `oneof` | 枚举值之一（大小写敏感） |
| `oneofci` 🔥 | 枚举值之一（忽略大小写，**自定义**） |
| `noneof` | 不是枚举值之一（大小写敏感） |
| `noneofci` | 不是枚举值之一（忽略大小写） |
</details>

<details>
<summary>地理位置</summary>

| 标签 | 说明 |
|------|------|
| `latitude` | 纬度 |
| `longitude` | 经度 |
| `timezone` | 时区 |
</details>

<details>
<summary>国家与货币</summary>

| 标签 | 说明 |
|------|------|
| `country_code` | 国家代码（别名） |
| `iso3166_1_alpha2` / `iso3166_1_alpha3` / `iso3166_1_alpha_numeric` | ISO 3166-1 |
| `iso3166_2` | ISO 3166-2 |
| `iso4217` / `iso4217_numeric` | 货币代码 |
| `postcode_iso3166_alpha2` / `postcode_iso3166_alpha2_field` | 邮政编码 |
| `bcp47_language_tag` / `bcp47_strict_language_tag` | 语言标签 |
| `eu_country_code` | 欧盟国家代码（别名） |
| `bic_iso_9362_2014` / `bic` | BIC / SWIFT 代码 |
</details>

<details>
<summary>加密与金融</summary>

| 标签 | 说明 |
|------|------|
| `eth_addr` / `eth_addr_checksum` | 以太坊地址 |
| `btc_addr` / `btc_addr_bech32` | 比特币地址 |
| `credit_card` | 信用卡号（Luhn） |
| `luhn_checksum` | Luhn 校验和 |
| `e164` | E.164 电话 |
| `phone` | 电话号码 |
</details>

<details>
<summary>文件与目录</summary>

| 标签 | 说明 |
|------|------|
| `file` | 文件路径 |
| `filepath` | 文件路径格式 |
| `dir` | 目录路径 |
| `dirpath` | 目录路径格式 |
| `image` | 图片文件 |
| `mimetype` | MIME 类型 |
</details>

<details>
<summary>其他</summary>

| 标签 | 说明 |
|------|------|
| `unique` | 切片/Map 中唯一值 |
| `semver` | 语义化版本 |
| `cron` | Cron 表达式 |
| `mongodb` / `mongodb_connection_string` | MongoDB |
| `dns_rfc1035_label` | DNS 标签 |
| `spicedb` | SpiceDB |
| `isdefault` | 是否为零值 |
| `skip_unless` | 条件不满足时跳过校验 |
</details>

> `oneofci` 是通过 `RegisterValidation` 注册的自定义校验器，支持大小写不敏感的枚举匹配。

## 🧪 测试

```bash
go test -v ./v2zh/...
```

## 📄 License

[MIT](LICENSE)
