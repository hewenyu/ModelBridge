// auth/auth.go
package auth

// Authenticator 接口定义了获取认证凭证的方法。
// 具体实现将根据不同平台的认证机制而有所不同。
type Authenticator interface {
	// GetCredentials 返回认证所需的凭证，可能是一个 map[string]string 或特定的结构体。
	// 例如，对于 API Key 认证，可能返回 {"Authorization": "Bearer YOUR_API_KEY"}。
	// 对于需要签名的请求，此方法可能更复杂，或由平台特定的 Handler 直接处理。
	GetCredentials() (map[string]string, error)
}

// TODO:
// - 根据平台需求，细化 Authenticator 接口或提供多种认证器实现。
// - 考虑凭证的缓存和刷新机制。
// - 实现对 platform.PlatformConfig 中 Credentials 字段的解析和使用。
