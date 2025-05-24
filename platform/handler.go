// platform/handler.go
package platform

import (
	"context"

	"github.com/hewenyu/modelbridge/models" // 假设的 module 路径，请根据实际情况修改
)

// PlatformHandler 定义了与特定大模型平台交互的通用接口。
// 每个受支持的平台都需要实现此接口。
type PlatformHandler interface {
	// TextGeneration 执行文本生成任务。
	TextGeneration(ctx context.Context, req *models.TextGenerationRequest) (*models.TextGenerationResponse, error)

	// ImageGeneration 执行图片生成任务。
	ImageGeneration(ctx context.Context, req *models.ImageGenerationRequest) (*models.ImageGenerationResponse, error)

	// Embedding 执行向量嵌入任务。
	Embedding(ctx context.Context, req *models.EmbeddingRequest) (*models.EmbeddingResponse, error)

	// ... 未来可以根据需要添加更多模型操作的方法，例如：
	// AudioTranscription(ctx context.Context, req *models.AudioTranscriptionRequest) (*models.AudioTranscriptionResponse, error)
	// TextToSpeech(ctx context.Context, req *models.TTSRequest) (*models.TTSResponse, error)

	// GetPlatformInfo 返回平台相关信息，例如平台名称、支持的模型等。
	// GetPlatformInfo() PlatformInfo // PlatformInfo 结构体待定义
}

// Provider 是用于标识不同大模型平台的类型。
type Provider string

const (
	ProviderVolcengine Provider = "volcengine"
	ProviderAlibaba    Provider = "alibaba"
	// 可以根据需要添加更多平台
)

// PlatformConfig 用于配置特定平台的客户端。
// 它在 `GETTING_STARTED.md` 中已有初步定义。
type PlatformConfig struct {
	Provider    Provider          `json:"provider"`    // 平台提供商
	Credentials map[string]string `json:"credentials"` // 平台凭证，例如 API Key, Secret Key 等
	// 其他平台特定的配置项，例如 Region, Endpoint 等
	// SpecificConfig map[string]interface{} `json:"specific_config,omitempty"`
}

// (可以考虑将 Provider 和 PlatformConfig 移至 client 包或一个更通用的 config 包，
// 因为它们是客户端初始化时直接使用的，而 PlatformHandler 是内部接口)
