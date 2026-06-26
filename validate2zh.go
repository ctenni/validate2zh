// Package validate2zh 提供结构体验证与中文错误翻译功能。
//
// 基于 go-playground/validator 封装，自动注册中文翻译器，
// 让校验错误信息直接返回中文友好提示。
//
// 基本用法:
//
//	import validate2zh "github.com/ctenni/validate2zh/v1"
//
//	err := validate2zh.ValidateStruct(myStruct)
package validate2zh

import "github.com/ctenni/validate2zh/v2zh"

// ValidateStruct 校验结构体，返回首条中文错误信息。
// 如果校验通过则返回 nil。
func ValidateStruct(s any) error {
	return v2zh.ValidateStruct(s)
}

// ValidateStructAll 校验结构体，返回所有中文错误信息。
// 如果校验通过则返回 nil。
func ValidateStructAll(s any) []error {
	return v2zh.ValidateStructAll(s)
}
