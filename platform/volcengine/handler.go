package volcengine

import (
	"fmt"
	"net/http" // Added for strings.TrimPrefix
	"time"

	"github.com/hewenyu/modelbridge/platform"
	// "github.com/hewenyu/modelbridge/errors" // SDK自定义错误，后续会用到
	// "github.com/hewenyu/modelbridge/client" // client.Logger，后续会用到
)

const (
	volcengineAPIKeyName         = "apiKey" // 与 GETTING_STARTED.md 中定义的凭证 key 一致
	volcengineChatCompletionsURL = "https://ark.cn-beijing.volces.com/api/v3/chat/completions"
	sseDataPrefix                = "data: "
	sseDataPrefixNoSpace         = "data:"
	sseDoneMessage               = "[DONE]"
	DefaultTimeout               = 10 * time.Second
)

// VolcengineHandler 实现了 PlatformHandler 接口，用于与火山方舟平台交互。
type VolcengineHandler struct {
	apiKey     string
	httpClient *http.Client
	// logger     client.Logger // 后续添加
}

// NewHandler 创建一个新的 VolcengineHandler 实例。
// config 参数用于传递平台特定的配置，例如 API Key。
// logger 参数用于日志记录。
func NewHandler(config *platform.PlatformConfig, opts ...Option) (*VolcengineHandler, error) {
	if config == nil {
		return nil, fmt.Errorf("volcengine handler: platform config cannot be nil") //  errors.New(errors.ErrCodeConfiguration, "volcengine handler: platform config cannot be nil")
	}

	apiKey, ok := config.Credentials[volcengineAPIKeyName]
	if !ok || apiKey == "" {
		return nil, fmt.Errorf("volcengine handler: API key not found or empty in credentials") // errors.New(errors.ErrCodeConfiguration, "volcengine handler: API key not found or empty in credentials")
	}

	handler := &VolcengineHandler{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: DefaultTimeout}, // 使用默认的 http.Client，后续可以配置超时等
		// logger:     logger,
	}

	for _, opt := range opts {
		opt(handler)
	}

	return handler, nil
}

// compile-time check to ensure VolcengineHandler implements PlatformHandler
var _ platform.PlatformHandler = (*VolcengineHandler)(nil)

func (e *volcengineError) Error() string {
	return fmt.Sprintf("volcengine API error: code=%s, message=%s, type=%s", e.Code, e.Message, e.Type)
}

// init registers the VolcengineHandler with the platform registry.
func init() {
	constructor := func(config *platform.PlatformConfig /*, logger client.Logger */) (platform.PlatformHandler, error) {
		return NewHandler(config /*, logger */) // *VolcengineHandler implements PlatformHandler
	}
	platform.RegisterHandler("volcengine", constructor)
}
