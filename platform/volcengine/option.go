package volcengine

import "time"

// Option 是用于配置 VolcengineHandler 的选项。
type Option func(*VolcengineHandler)

// WithTimeout 设置 HTTP 请求的超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(h *VolcengineHandler) {
		h.httpClient.Timeout = timeout
	}
}
