# 模型类型

本文档描述了 SDK 支持的每种模型类型的通用请求和响应结构。特定平台的模型 ID 或名称将尽可能映射到这些通用类型。

## 1. 文本生成 (Text Generation)

*   **描述:** 根据给定的提示生成文本的模型。
*   **通用请求 (`models.TextGenerationRequest`):**
    ```go
    type TextGenerationRequest struct {
        Prompt         string            `json:"prompt"` // 输入的提示文本
        Model          string            `json:"model,omitempty"` // 平台特定的模型 ID 或通用别名
        MaxTokens      int               `json:"max_tokens,omitempty"` // 生成文本的最大长度
        Temperature    float32           `json:"temperature,omitempty"` // 控制生成文本的随机性
        TopP           float32           `json:"top_p,omitempty"` // 控制核心采样的概率阈值
        StopSequences  []string          `json:"stop_sequences,omitempty"` // 遇到即停止生成的序列
        Stream         bool              `json:"stream,omitempty"` // 如果为 true，则响应将是流式传输
        // ... 其他通用参数
        PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
    }
    ```
*   **通用响应 (`models.TextGenerationResponse`):**
    ```go
    type TextGenerationResponse struct {
        ID             string   `json:"id"` // 请求的唯一标识符
        GeneratedText  string   `json:"generated_text"` // 生成的文本内容
        FinishReason   string   `json:"finish_reason"` // 完成原因，例如 "stop" (自然停止), "length" (达到最大长度)
        TokenUsage     struct { // Token 使用情况 (如果平台提供)
            PromptTokens     int `json:"prompt_tokens"`     // 输入提示的 Token 数量
            CompletionTokens int `json:"completion_tokens"` // 生成文本的 Token 数量
            TotalTokens      int `json:"total_tokens"`      // 总 Token 数量
        } `json:"token_usage,omitempty"`
        // ... 其他通用字段
    }
    ```
*   **通用流式块 (如果 `Stream: true`):**
    ```go
    type TextGenerationStreamChunk struct {
        ID      string `json:"id"`      // 块的唯一标识符或关联请求的ID
        Delta   string `json:"delta"`   // 生成的文本块
        IsFinal bool   `json:"is_final"` // 是否是最后一个块
        // ... 其他流特定的字段
    }
    ```

## 2. 多模态 (Multimodal)

*   **描述:** 能够处理和生成来自多种模态（例如文本、图像、音频）信息的模型。
*   **(请求/响应结构待定 - 需要针对每个平台确定更具体的用例)**

## 3. 推理模型 (Inference Models)

*   **描述:** 通用推理端点，通常用于自定义或微调模型。
*   **(请求/响应结构待定)**

## 4. 音频理解 (Audio Understanding)

*   **描述:** 转录或分析音频内容的模型。
*   **(请求/响应结构待定)**

## 5. 视频理解 (Video Understanding)

*   **描述:** 分析视频内容的模型。
*   **(请求/响应结构待定)**

## 6. 视频生成 (Video Generation)

*   **描述:** 从文本或其他输入生成视频的模型。
*   **(请求/响应结构待定)**

## 7. 图片处理 (Image Processing)

*   **描述:** 用于图像增强、滤波等任务的模型。
*   **(请求/响应结构待定)**

## 8. 图片理解 (Image Understanding)

*   **描述:** 分析和描述图像的模型（例如，图像字幕、对象检测）。
*   **(请求/响应结构待定)**

## 9. 图片生成 (Image Generation)

*   **描述:** 从文本提示生成图像的模型（文生图）。
*   **通用请求 (`models.ImageGenerationRequest`):**
    ```go
    type ImageGenerationRequest struct {
        Prompt         string            `json:"prompt"` // 输入的提示文本
        Model          string            `json:"model,omitempty"` // 平台特定的模型 ID 或通用别名
        N              int               `json:"n,omitempty"` // 要生成的图片数量
        Size           string            `json:"size,omitempty"` // 图片尺寸，例如 "1024x1024"
        Quality        string            `json:"quality,omitempty"` // 图片质量，例如 "standard", "hd"
        Style          string            `json:"style,omitempty"` // 图片风格，例如 "vivid" (鲜明), "natural" (自然)
        // ... 其他通用参数
        PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
    }
    ```
*   **通用响应 (`models.ImageGenerationResponse`):**
    ```go
    type ImageGenerationResponse struct {
        ID      string  `json:"id"`     // 请求的唯一标识符
        Images  []Image `json:"images"` // 生成的图片列表
    }

    type Image struct {
        URL     string `json:"url,omitempty"`      // 生成图片的 URL (如果可用)
        Base64  string `json:"base64,omitempty"`  // Base64 编码的图片数据 (如果可用)
        RevisedPrompt string `json:"revised_prompt,omitempty"` // 如果平台修改了原始提示，则为修改后的提示
    }
    ```

## 10. 向量模型 (Embeddings / Vector Models)

*   **描述:** 为文本或其他数据生成向量嵌入的模型。
*   **通用请求 (`models.EmbeddingRequest`):**
    ```go
    type EmbeddingRequest struct {
        Input          []string          `json:"input"` // 需要进行向量嵌入的文本列表
        Model          string            `json:"model,omitempty"` // 平台特定的模型 ID 或通用别名
        EncodingFormat string            `json:"encoding_format,omitempty"` // 编码格式，例如 "float", "base64"
        // ... 其他通用参数
        PlatformSpecificParams map[string]interface{} `json:"platform_specific_params,omitempty"` // 用于覆盖平台特定参数
    }
    ```
*   **通用响应 (`models.EmbeddingResponse`):**
    ```go
    type EmbeddingResponse struct {
        ID          string      `json:"id"`         // 请求的唯一标识符
        Embeddings  []Embedding `json:"embeddings"` // 生成的向量嵌入列表
        TokenUsage  struct { // Token 使用情况 (如果平台提供)
            PromptTokens int `json:"prompt_tokens"` // 输入提示的 Token 数量
            TotalTokens  int `json:"total_tokens"`  // 总 Token 数量
        } `json:"token_usage,omitempty"`
    }

    type Embedding struct {
        Index     int       `json:"index"`     // 对应输入列表中的索引
        Embedding []float32 `json:"embedding"` // 向量嵌入数据 (或根据 EncodingFormat 确定适当类型)
    }
    ```

## 11. 语音合成 (Text-to-Speech / TTS)

*   **描述:** 将文本转换为语音音频的模型。
*   **(请求/响应结构待定)**

## 12. 语音识别 (Audio-to-Text / ASR)

*   **描述:** 将语音音频转录为文本的模型。
*   **(请求/响应结构待定)**

## 13. 排序模型 (Ranking Models)

*   **描述:** 用于对搜索结果或其他项目进行重新排序的模型。
*   **(请求/响应结构待定)**

*(随着对特定平台 API 的深入研究，将详细说明更多模型类型。)* 