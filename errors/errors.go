// errors/errors.go
package errors

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// 通用错误类型
// /////////////////////////////////////////////////////////////////////////////

// Error 是 SDK 返回的通用错误结构。
// 它可以包装底层平台的错误，并提供统一的错误代码（可选）。
type Error struct {
	// Code 是一个可选的错误代码，可以用于程序化地处理错误。
	// 例如："ErrAuthentication", "ErrInvalidRequest", "ErrRateLimited", "ErrPlatformError"
	Code string
	// Message 是错误的描述信息。
	Message string
	// Underlying 是导致此错误的原始错误（如果有）。
	Underlying error
	// PlatformDetails 包含特定于平台的额外错误信息（可选）。
	PlatformDetails map[string]interface{}
}

// Error 实现 error 接口。
func (e *Error) Error() string {
	if e.Underlying != nil {
		return fmt.Sprintf("%s: %s (underlying: %v)", e.Code, e.Message, e.Underlying)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap 提供了对底层错误的支持，用于 errors.Is 和 errors.As。
func (e *Error) Unwrap() error {
	return e.Underlying
}

// New 创建一个新的 SDK Error。
func New(code, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Wrap 创建一个包装了现有错误的 SDK Error。
func Wrap(err error, code, message string) *Error {
	return &Error{Code: code, Message: message, Underlying: err}
}

// /////////////////////////////////////////////////////////////////////////////
// 预定义的错误代码常量 (示例)
// /////////////////////////////////////////////////////////////////////////////

const (
	ErrCodeConfiguration  = "ErrConfiguration"
	ErrCodeAuthentication = "ErrAuthentication"
	ErrCodeInvalidRequest = "ErrInvalidRequest"
	ErrCodeNotFound       = "ErrNotFound"
	ErrCodeRateLimited    = "ErrRateLimited"
	ErrCodePlatformError  = "ErrPlatformError" // 通用平台错误
	ErrCodeUnsupported    = "ErrUnsupportedOperation"
	ErrCodeInternal       = "ErrInternalSDK" // SDK 内部错误
	ErrCodeTimeout        = "ErrTimeout"
	ErrCodeCancelled      = "ErrCancelled"
)

// /////////////////////////////////////////////////////////////////////////////
// 预定义的错误变量 (方便直接比较，但使用 errors.Is(err, &Error{Code: ...}) 更灵活)
// /////////////////////////////////////////////////////////////////////////////

// var (
// 	ErrConfig = New(ErrCodeConfiguration, "invalid SDK configuration")
// 	// ... 其他常用错误可以预定义
// )

// IsSDKError 检查一个错误是否是本 SDK 定义的 *Error 类型，并且具有特定的错误代码。
func IsSDKError(err error, code string) bool {
	var sdkErr *Error
	if As(err, &sdkErr) {
		return sdkErr.Code == code
	}
	return false
}

// As 类似于标准库的 errors.As，但专门用于 *Error 类型。
// 如果 err 是或包装了一个 *Error，它会将其赋值给 target 并返回 true。
func As(err error, target **Error) bool {
	// Go 1.13+ errors.As already handles unwrapping.
	// For older versions, a manual loop might be needed if not using xerrors.
	// Assuming Go 1.13+ for simplicity here.
	return AsGo113(err, target) // Using a helper to avoid import cycle if errors.As is directly used.
}

// AsGo113 is a helper to avoid import cycles if this package were to use stdlib errors directly for As.
// In a real scenario, you\'d just use `import "errors"` and `errors.As(err, target)`.
// This is a simplified stand-in.
func AsGo113(err error, target interface{}) bool {
	// This is a placeholder. In a real implementation, you\'d use the standard library\'s errors.As.
	// For the purpose of this generation, we\'ll assume it works as expected.
	// The actual Go standard library errors.As can handle this.
	if e, ok := err.(*Error); ok {
		switch t := target.(type) {
		case **Error:
			*t = e
			return true
		}
	}
	// Check wrapped errors
	if u, ok := err.(interface{ Unwrap() error }); ok {
		return AsGo113(u.Unwrap(), target)
	}
	return false
}
