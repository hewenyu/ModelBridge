// client/client.go
package client

import (
	"context"
	"fmt"
	"log" // 标准库 log
	"os"  // 用于默认 logger

	"github.com/hewenyu/modelbridge/models"   // 假设的 module 路径
	"github.com/hewenyu/modelbridge/platform" // 假设的 module 路径
	// 计划在这里导入具体的平台实现，例如：
	// "github.com/hewenyu/modelbridge/platform/alibaba"
	// "github.com/hewenyu/modelbridge/platform/volcengine"
)

// Logger 是一个简单的日志接口，允许用户提供自定义的日志实现。
type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Client 是与大模型平台交互的统一客户端。
type Client struct {
	handler platform.PlatformHandler // 内部持有一个特定平台的处理器
	logger  Logger                   // 添加 logger 字段
}

// defaultLogger 是一个使用标准库 log.Logger 的默认实现。
type defaultLogger struct {
	stdlog *log.Logger
}

func (l *defaultLogger) Printf(format string, v ...interface{}) {
	l.stdlog.Printf(format, v...)
}

func (l *defaultLogger) Println(v ...interface{}) {
	l.stdlog.Println(v...)
}

// NewDefaultLogger 创建一个默认的 logger。
func NewDefaultLogger() Logger {
	return &defaultLogger{stdlog: log.New(os.Stderr, "[ModelBridgeSDK] ", log.LstdFlags|log.Lshortfile)}
}

// Option 是用于配置 Client 的函数选项类型。
type Option func(*Client) error

// WithLogger 设置自定义的 logger。
func WithLogger(logger Logger) Option {
	return func(c *Client) error {
		if logger == nil {
			return fmt.Errorf("logger cannot be nil") // 或者可以允许 nil 并内部处理
		}
		c.logger = logger
		return nil
	}
}

// NewClient 根据提供的平台配置创建一个新的客户端实例。
// 可以通过传入 Option 函数来定制客户端，例如 WithLogger。
func NewClient(config *platform.PlatformConfig, opts ...Option) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("platform config cannot be nil") // 可以使用 errors.New SDK 错误
	}

	c := &Client{
		// 设置默认 logger，后续可以被 Option覆盖
		logger: NewDefaultLogger(),
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err) // 可以使用 errors.Wrap
		}
	}

	var handler platform.PlatformHandler
	var err error

	c.logger.Printf("Initializing new client for provider: %s", config.Provider)

	switch config.Provider {
	case platform.ProviderAlibaba:
		// handler, err = alibaba.NewHandler(config, c.logger) // 传递 logger 给 handler
		c.logger.Println("Alibaba provider selected (not yet implemented)")
		// TODO: 实现阿里百炼平台的 Handler 初始化
		err = fmt.Errorf("alibaba provider not yet implemented")
	case platform.ProviderVolcengine:
		// handler, err = volcengine.NewHandler(config, c.logger) // 传递 logger 给 handler
		c.logger.Println("Volcengine provider selected (not yet implemented)")
		// TODO: 实现火山方舟平台的 Handler 初始化
		err = fmt.Errorf("volcengine provider not yet implemented")
	default:
		err = fmt.Errorf("unsupported provider: %s", config.Provider)
	}

	if err != nil {
		c.logger.Printf("Failed to create platform handler for %s: %v", config.Provider, err)
		return nil, fmt.Errorf("failed to create platform handler for %s: %w", config.Provider, err) // 可以使用 errors.Wrap
	}
	// 理论上，如果上面的 case 没有正确实现，handler 可能依然是 nil
	// 当具体 handler 实现后，这部分逻辑需要调整
	if handler == nil && err == nil { // 确保如果没错误，handler 必须被设置
		err = fmt.Errorf("handler not initialized for provider %s despite no error", config.Provider)
		c.logger.Printf("Error: %v", err)
		return nil, err
	}

	c.handler = handler
	c.logger.Println("Client initialized successfully.")
	return c, nil
}

// TextGeneration 使用配置的平台执行文本生成任务。
func (c *Client) TextGeneration(ctx context.Context, req *models.TextGenerationRequest) (*models.TextGenerationResponse, error) {
	if c.handler == nil {
		err := fmt.Errorf("client not properly initialized or no handler set")
		c.logger.Printf("Error in TextGeneration: %v", err)
		return nil, err // 可以返回自定义的 SDK Error
	}
	c.logger.Printf("Executing TextGeneration for model '%s' with prompt: \"%s...\"", req.Model, truncateForLog(req.Prompt, 30))
	resp, err := c.handler.TextGeneration(ctx, req)
	if err != nil {
		c.logger.Printf("Error from platform handler in TextGeneration: %v", err)
	}
	return resp, err
}

// ImageGeneration 使用配置的平台执行图片生成任务。
func (c *Client) ImageGeneration(ctx context.Context, req *models.ImageGenerationRequest) (*models.ImageGenerationResponse, error) {
	if c.handler == nil {
		err := fmt.Errorf("client not properly initialized or no handler set")
		c.logger.Printf("Error in ImageGeneration: %v", err)
		return nil, err
	}
	c.logger.Printf("Executing ImageGeneration for model '%s' with prompt: \"%s...\"", req.Model, truncateForLog(req.Prompt, 30))
	resp, err := c.handler.ImageGeneration(ctx, req)
	if err != nil {
		c.logger.Printf("Error from platform handler in ImageGeneration: %v", err)
	}
	return resp, err
}

// Embedding 使用配置的平台执行向量嵌入任务。
func (c *Client) Embedding(ctx context.Context, req *models.EmbeddingRequest) (*models.EmbeddingResponse, error) {
	if c.handler == nil {
		err := fmt.Errorf("client not properly initialized or no handler set")
		c.logger.Printf("Error in Embedding: %v", err)
		return nil, err
	}
	inputCount := len(req.Input)
	firstInput := ""
	if inputCount > 0 {
		firstInput = req.Input[0]
	}
	c.logger.Printf("Executing Embedding for model '%s' with %d inputs, first input: \"%s...\"", req.Model, inputCount, truncateForLog(firstInput, 30))
	resp, err := c.handler.Embedding(ctx, req)
	if err != nil {
		c.logger.Printf("Error from platform handler in Embedding: %v", err)
	}
	return resp, err
}

// truncateForLog is a helper function to truncate strings for logging.
// Note: This is a simple byte-wise truncation and may cut multi-byte characters.
func truncateForLog(s string, maxLen int) string {
	if len(s) > maxLen {
		if maxLen > 3 { // Ensure "..." fits
			return s[:maxLen-3] + "..."
		}
		return s[:maxLen]
	}
	return s
}

// Add more methods here to expose other functionalities from PlatformHandler as needed.
