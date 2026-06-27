<div align="center">

# validate2zh

**Go struct validation · Chinese error messages**

[![Go Reference](https://pkg.go.dev/badge/github.com/ctenni/validate2zh.svg)](https://pkg.go.dev/github.com/ctenni/validate2zh)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ctenni/validate2zh)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/ctenni/validate2zh)](LICENSE)
[![Tag](https://img.shields.io/github/v/tag/ctenni/validate2zh)](https://github.com/ctenni/validate2zh/tags)

</div>

A wrapper around [`go-playground/validator`](https://github.com/go-playground/validator)  
that automatically registers Chinese translations, returning validation errors in **Chinese**.

> 🇨🇳 中文版请查看 [README.cn.md](README.cn.md)

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

<details>
<summary>Required family</summary>

| Tag | Description |
|-----|-------------|
| `required` | Field must not be zero value |
| `required_if` | Required if another field equals given value |
| `required_unless` | Required unless another field equals given value |
| `required_with` | Required if any of given fields are present |
| `required_with_all` | Required if all of given fields are present |
| `required_without` | Required if any of given fields are absent |
| `required_without_all` | Required if all of given fields are absent |
</details>

<details>
<summary>Excluded family</summary>

| Tag | Description |
|-----|-------------|
| `excluded_if` | Excluded if another field equals given value |
| `excluded_unless` | Excluded unless another field equals given value |
| `excluded_with` | Excluded if any of given fields are present |
| `excluded_with_all` | Excluded if all of given fields are present |
| `excluded_without` | Excluded if any of given fields are absent |
| `excluded_without_all` | Excluded if all of given fields are absent |
</details>

<details>
<summary>Comparison</summary>

| Tag | Description |
|-----|-------------|
| `eq` | Equal to given value |
| `eq_ignore_case` | Equal to given value (case-insensitive) |
| `ne` | Not equal to given value |
| `ne_ignore_case` | Not equal to given value (case-insensitive) |
| `gt` | Greater than |
| `gte` | Greater than or equal |
| `lt` | Less than |
| `lte` | Less than or equal |
</details>

<details>
<summary>Cross-field comparison</summary>

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
</details>

<details>
<summary>String content</summary>

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
</details>

<details>
<summary>Character composition</summary>

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
</details>

<details>
<summary>Numeric formats</summary>

| Tag | Description |
|-----|-------------|
| `number` | Integer string |
| `numeric` | Integer or float string |
| `boolean` | Boolean string |
| `hexadecimal` | Hexadecimal string |
| `hexcolor` | Hex color (#RRGGBB) |
| `rgb` | RGB color |
| `rgba` | RGBA color |
| `hsl` | HSL color |
| `hsla` | HSLA color |
| `cmyk` | CMYK color |
| `iscolor` | Any color format (alias) |
</details>

<details>
<summary>Network formats</summary>

| Tag | Description |
|-----|-------------|
| `ip` | IP address (v4 or v6) |
| `ipv4` | IPv4 address |
| `ipv6` | IPv6 address |
| `cidr` | CIDR notation |
| `cidrv4` | CIDR v4 |
| `cidrv6` | CIDR v6 |
| `mac` | MAC address |
| `hostname` | Hostname |
| `hostname_rfc1123` | RFC 1123 hostname |
| `hostname_port` | Hostname:port |
| `fqdn` | FQDN |
| `port` | Port number |
| `tcp4_addr` / `tcp6_addr` / `tcp_addr` | TCP address |
| `udp4_addr` / `udp6_addr` / `udp_addr` | UDP address |
| `ip4_addr` / `ip6_addr` / `ip_addr` | IP address |
| `unix_addr` | Unix address |
| `uds_exists` | Unix domain socket exists |
</details>

<details>
<summary>Format validation</summary>

| Tag | Description |
|-----|-------------|
| `email` | Email address |
| `url` | URL |
| `http_url` | HTTP/HTTPS URL |
| `https_url` | HTTPS URL |
| `uri` | URI |
| `origin` | Origin |
| `urn_rfc2141` | URN |
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
| `datetime` | Datetime with given format |
</details>

<details>
<summary>Identifier — UUID / ISBN / SSN</summary>

| Tag | Description |
|-----|-------------|
| `uuid` / `uuid3` / `uuid4` / `uuid5` | UUID |
| `uuid_rfc4122` / `uuid3_rfc4122` / `uuid4_rfc4122` / `uuid5_rfc4122` | UUID RFC 4122 |
| `ulid` | ULID |
| `isbn` / `isbn10` / `isbn13` | ISBN |
| `issn` | ISSN |
| `ssn` | SSN |
| `ein` | EIN |
| `cve` | CVE identifier |
</details>

<details>
<summary>Hash algorithms</summary>

| Tag | Description |
|-----|-------------|
| `md4` / `md5` | MD |
| `sha256` / `sha384` / `sha512` | SHA |
| `ripemd128` / `ripemd160` | RIPEMD |
| `tiger128` / `tiger160` / `tiger192` | TIGER |
</details>

<details>
<summary>Enum</summary>

| Tag | Description |
|-----|-------------|
| `oneof` | One of given values (case-sensitive) |
| `oneofci` 🔥 | One of given values (case-insensitive, **custom**) |
| `noneof` | None of given values (case-sensitive) |
| `noneofci` | None of given values (case-insensitive) |
</details>

<details>
<summary>Geo & timezone</summary>

| Tag | Description |
|-----|-------------|
| `latitude` | Latitude coordinate |
| `longitude` | Longitude coordinate |
| `timezone` | Timezone string |
</details>

<details>
<summary>Country / Currency</summary>

| Tag | Description |
|-----|-------------|
| `country_code` | Country code (alias) |
| `iso3166_1_alpha2` / `iso3166_1_alpha3` / `iso3166_1_alpha_numeric` | ISO 3166-1 |
| `iso3166_2` | ISO 3166-2 |
| `iso4217` / `iso4217_numeric` | Currency |
| `postcode_iso3166_alpha2` / `postcode_iso3166_alpha2_field` | Postal code |
| `bcp47_language_tag` / `bcp47_strict_language_tag` | Language tag |
| `eu_country_code` | EU country code (alias) |
| `bic_iso_9362_2014` / `bic` | BIC / SWIFT |
</details>

<details>
<summary>Crypto / Finance</summary>

| Tag | Description |
|-----|-------------|
| `eth_addr` / `eth_addr_checksum` | Ethereum address |
| `btc_addr` / `btc_addr_bech32` | Bitcoin address |
| `credit_card` | Credit card (Luhn) |
| `luhn_checksum` | Luhn checksum |
| `e164` | E.164 phone |
| `phone` | Phone number |
</details>

<details>
<summary>File & directory</summary>

| Tag | Description |
|-----|-------------|
| `file` | File path |
| `filepath` | File path (valid format) |
| `dir` | Directory path |
| `dirpath` | Directory path (valid format) |
| `image` | Image file |
| `mimetype` | MIME type |
</details>

<details>
<summary>Other</summary>

| Tag | Description |
|-----|-------------|
| `unique` | Unique values in slice/map |
| `semver` | Semantic version |
| `cron` | Cron expression |
| `mongodb` / `mongodb_connection_string` | MongoDB |
| `dns_rfc1035_label` | DNS label |
| `spicedb` | SpiceDB |
| `isdefault` | Is default value |
| `skip_unless` | Skip unless condition met |
</details>

> `oneofci` is a custom validator registered via `RegisterValidation` — supports case-insensitive enum matching.

## 🧪 Testing

```bash
go test -v ./v2zh/...
```

## 📄 License

[MIT](LICENSE)
