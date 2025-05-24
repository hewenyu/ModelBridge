// models/models.go
package models

// TextGenerationRequest 定义了文本生成请求的结构。
type TextGenerationRequest struct {
	Prompt                 string                 `json:"prompt"`                             // 输入的提示文本
	Model                  string                 `json:"model,omitempty"`                    // 平台特定的模型 ID 或通用别名
	MaxTokens              int                    `json:"max_tokens,omitempty"`               // 生成文本的最大长度
	Temperature            float32                `json:"temperature,omitempty"`              // 控制生成文本的随机性
	TopP                   float32                `json:"top_p,omitempty"`                    // 控制核心采样的概率阈值
	StopSequences          []string               `json:"stop_sequences,omitempty"`           // 遇到即停止生成的序列
	Stream                 bool                   `json:"stream,omitempty"`                   // 如果为 true，则响应将是流式传输
	PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
}

// TextGenerationResponse 定义了文本生成响应的结构。
type TextGenerationResponse struct {
	ID            string   `json:"id"`             // 请求的唯一标识符
	GeneratedText string   `json:"generated_text"` // 生成的文本内容
	FinishReason  string   `json:"finish_reason"`  // 完成原因，例如 "stop" (自然停止), "length" (达到最大长度)
	TokenUsage    struct { // Token 使用情况 (如果平台提供)
		PromptTokens     int `json:"prompt_tokens"`     // 输入提示的 Token 数量
		CompletionTokens int `json:"completion_tokens"` // 生成文本的 Token 数量
		TotalTokens      int `json:"total_tokens"`      // 总 Token 数量
	} `json:"token_usage,omitempty"`
}

// TextGenerationStreamChunk 定义了文本生成流式响应的块结构。
type TextGenerationStreamChunk struct {
	ID      string `json:"id"`       // 块的唯一标识符或关联请求的ID
	Delta   string `json:"delta"`    // 生成的文本块
	IsFinal bool   `json:"is_final"` // 是否是最后一个块
}

// ImageGenerationRequest 定义了图片生成请求的结构。
type ImageGenerationRequest struct {
	Prompt                 string                 `json:"prompt"`                             // 输入的提示文本
	Model                  string                 `json:"model,omitempty"`                    // 平台特定的模型 ID 或通用别名
	N                      int                    `json:"n,omitempty"`                        // 要生成的图片数量
	Size                   string                 `json:"size,omitempty"`                     // 图片尺寸，例如 "1024x1024"
	Quality                string                 `json:"quality,omitempty"`                  // 图片质量，例如 "standard", "hd"
	Style                  string                 `json:"style,omitempty"`                    // 图片风格，例如 "vivid" (鲜明), "natural" (自然)
	PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
}

// ImageGenerationResponse 定义了图片生成响应的结构。
type ImageGenerationResponse struct {
	ID     string  `json:"id"`     // 请求的唯一标识符
	Images []Image `json:"images"` // 生成的图片列表
}

// Image 定义了生成图片的信息。
type Image struct {
	URL           string `json:"url,omitempty"`            // 生成图片的 URL (如果可用)
	Base64        string `json:"base64,omitempty"`         // Base64 编码的图片数据 (如果可用)
	RevisedPrompt string `json:"revised_prompt,omitempty"` // 如果平台修改了原始提示，则为修改后的提示
}

// EmbeddingRequest 定义了向量嵌入请求的结构。
type EmbeddingRequest struct {
	Input                  []string               `json:"input"`                              // 需要进行向量嵌入的文本列表
	Model                  string                 `json:"model,omitempty"`                    // 平台特定的模型 ID 或通用别名
	EncodingFormat         string                 `json:"encoding_format,omitempty"`          // 编码格式，例如 "float", "base64"
	PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
}

// EmbeddingResponse 定义了向量嵌入响应的结构。
type EmbeddingResponse struct {
	ID         string      `json:"id"`         // 请求的唯一标识符
	Embeddings []Embedding `json:"embeddings"` // 生成的向量嵌入列表
	TokenUsage struct {    // Token 使用情况 (如果平台提供)
		PromptTokens int `json:"prompt_tokens"` // 输入提示的 Token 数量
		TotalTokens  int `json:"total_tokens"`  // 总 Token 数量
	} `json:"token_usage,omitempty"`
}

// Embedding 定义了单个向量嵌入的数据。
type Embedding struct {
	Index     int       `json:"index"`     // 对应输入列表中的索引
	Embedding []float32 `json:"embedding"` // 向量嵌入数据 (或根据 EncodingFormat 确定适当类型)
}
