package volcengine

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hewenyu/modelbridge/models"
)

// volcengineChatRequest 是火山方舟对话 API 的请求体结构。
type volcengineChatRequest struct {
	Model         volcengineModel          `json:"model"`
	Messages      []volcengineChatMessage  `json:"messages"`
	Stream        bool                     `json:"stream,omitempty"`
	StreamOptions *volcengineStreamOptions `json:"stream_options,omitempty"`
	User          string                   `json:"user,omitempty"`
	MaxTokens     int                      `json:"max_tokens,omitempty"`
	Temperature   float32                  `json:"temperature,omitempty"`
	TopP          float32                  `json:"top_p,omitempty"`
	Stop          []string                 `json:"stop,omitempty"`
	// Tools          []volcengineTool          `json:"tools,omitempty"` // 暂时不支持
}

type volcengineModel struct {
	ModelID string `json:"model_id"`
	Version string `json:"version,omitempty"`
}

type volcengineChatMessage struct {
	Role    string `json:"role"` // user, assistant, system
	Content string `json:"content"`
}

type volcengineStreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

// volcengineChatResponse 是火山方舟对话 API 的响应体结构 (非流式)。
type volcengineChatResponse struct {
	ID      string               `json:"id"`
	Object  string               `json:"object"`
	Created int64                `json:"created"`
	Model   string               `json:"model"`
	Choices []volcengineChoice   `json:"choices"`
	Usage   volcengineTokenUsage `json:"usage"`
	Error   *volcengineError     `json:"error,omitempty"` // 火山方舟特定的错误结构
}

// volcengineStreamChatCompletionChunk 是火山方舟对话 API 流式响应中每个 chunk 的结构。
type volcengineStreamChatCompletionChunk struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"` // e.g., "chat.completion.chunk"
	Created int64                    `json:"created"`
	Model   string                   `json:"model"`
	Choices []volcengineStreamChoice `json:"choices"`
	Usage   *volcengineTokenUsage    `json:"usage,omitempty"` // Typically null until the last chunk if include_usage is true
	Error   *volcengineError         `json:"error,omitempty"`
}

type volcengineStreamChoice struct {
	Index        int                   `json:"index"`
	Delta        volcengineChatMessage `json:"delta"`                   // Contains the incremental content
	FinishReason *string               `json:"finish_reason,omitempty"` // Null until the last chunk for a choice
	// Logprobs     interface{} `json:"logprobs,omitempty"` // Not handled yet
}

type volcengineChoice struct {
	Index        int                   `json:"index"`
	Message      volcengineChatMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}

type volcengineTokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// volcengineError 定义了火山方舟 API 返回的错误信息结构。
type volcengineError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type,omitempty"`
}

// TextGeneration 实现文本生成逻辑。
func (h *VolcengineHandler) TextGeneration(ctx context.Context, req *models.TextGenerationRequest) (*models.TextGenerationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("volcengine handler: text generation request cannot be nil") // errors.New(errors.ErrCodeInvalidRequest, "volcengine handler: text generation request cannot be nil")
	}

	// 1. 将 models.TextGenerationRequest 转换为 volcengineChatRequest
	volcReq := volcengineChatRequest{
		Model: volcengineModel{
			ModelID: req.Model, // 假设用户直接提供火山方舟的模型 ID
			// Version: "...", //  如果需要，可以从 req.PlatformSpecificParams 获取
		},
		Messages: []volcengineChatMessage{
			{
				Role:    "user", // 默认将 prompt 作为 user message
				Content: req.Prompt,
			},
		},
		Stream:      req.Stream,
		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
		TopP:        req.TopP,
		Stop:        req.StopSequences,
		// User: 从 req.PlatformSpecificParams 获取,
	}

	// 如果是流式请求且需要在最后包含用量信息
	if req.Stream {
		// 示例：从 PlatformSpecificParams 获取 stream_options.include_usage
		if includeUsage, ok := req.PlatformSpecificParams["volc_stream_options_include_usage"].(bool); ok && includeUsage {
			volcReq.StreamOptions = &volcengineStreamOptions{IncludeUsage: true}
		}
	}

	// TODO: 处理 req.PlatformSpecificParams 中更复杂的 messages 结构 (system, assistant roles)
	// TODO: 处理 Tools (如果未来支持)

	// 2. 序列化请求体
	reqBodyBytes, err := json.Marshal(volcReq)
	if err != nil {
		// h.logger.Printf("Error marshalling volcengine request: %v", err)
		return nil, fmt.Errorf("volcengine handler: failed to marshal request body: %w", err) // errors.Wrap(err, errors.ErrCodeInternal, "volcengine handler: failed to marshal request body")
	}

	// 3. 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, volcengineChatCompletionsURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		// h.logger.Printf("Error creating volcengine HTTP request: %v", err)
		return nil, fmt.Errorf("volcengine handler: failed to create HTTP request: %w", err) // errors.Wrap(err, errors.ErrCodeInternal, "volcengine handler: failed to create HTTP request")
	}

	// 4. 设置请求头
	httpReq.Header.Set("Authorization", "Bearer "+h.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")
	if req.Stream {
		httpReq.Header.Set("Accept", "text/event-stream")
	} else {
		httpReq.Header.Set("Accept", "application/json")
	}

	// 5. 发送请求
	// h.logger.Printf("Sending TextGeneration request to Volcengine for model: %s", req.Model)
	httpResp, err := h.httpClient.Do(httpReq)
	if err != nil {
		// h.logger.Printf("Error sending request to Volcengine: %v", err)
		return nil, fmt.Errorf("volcengine handler: failed to send HTTP request: %w", err) // errors.Wrap(err, errors.ErrCodePlatformError, "volcengine handler: failed to send HTTP request")
	}
	defer httpResp.Body.Close()

	// 6. 处理响应
	if httpResp.StatusCode != http.StatusOK {
		respBodyBytes, _ := io.ReadAll(httpResp.Body) // Try to read body for error details
		// h.logger.Printf("Volcengine API error: status code %d, body: %s", httpResp.StatusCode, string(respBodyBytes))
		var volcErrResp volcengineChatResponse // Try to parse as non-stream error first
		if json.Unmarshal(respBodyBytes, &volcErrResp) == nil && volcErrResp.Error != nil {
			return nil, fmt.Errorf("volcengine API error: status %d, code %s, message: %s", httpResp.StatusCode, volcErrResp.Error.Code, volcErrResp.Error.Message)
		}
		var volcStreamErrResp volcengineStreamChatCompletionChunk // Try to parse as stream error
		if json.Unmarshal(respBodyBytes, &volcStreamErrResp) == nil && volcStreamErrResp.Error != nil {
			return nil, fmt.Errorf("volcengine API stream error: status %d, code %s, message: %s", httpResp.StatusCode, volcStreamErrResp.Error.Code, volcStreamErrResp.Error.Message)
		}
		return nil, fmt.Errorf("volcengine API error: status code %d, response: %s", httpResp.StatusCode, string(respBodyBytes))
		// errors.New(errors.ErrCodePlatformError, fmt.Sprintf("volcengine API error: status code %d, response: %s\", httpResp.StatusCode, string(respBodyBytes)))
	}

	if req.Stream {
		// 处理流式响应
		var fullTextBuilder strings.Builder
		var finalResponseID string
		var finalFinishReason string
		var finalTokenUsage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		}

		scanner := bufio.NewScanner(httpResp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" { // Skip empty lines
				continue
			}

			var dataStr string
			if strings.HasPrefix(line, sseDataPrefix) {
				dataStr = strings.TrimPrefix(line, sseDataPrefix)
			} else if strings.HasPrefix(line, sseDataPrefixNoSpace) { // Handle case without space after "data:"
				dataStr = strings.TrimPrefix(line, sseDataPrefixNoSpace)
			} else {
				// h.logger.Printf("Volcengine stream: received non-SSE line: %s", line)
				continue // Ignore lines not starting with "data: " or "data:"
			}

			if dataStr == sseDoneMessage {
				// h.logger.Println("Volcengine stream: received [DONE] message.")
				break // End of stream
			}

			var chunk volcengineStreamChatCompletionChunk
			if err := json.Unmarshal([]byte(dataStr), &chunk); err != nil {
				// h.logger.Printf("Error unmarshalling Volcengine stream chunk: %v. Data: %s", err, dataStr)
				return nil, fmt.Errorf("volcengine handler: failed to unmarshal stream chunk: %w. Data: %s", err, dataStr)
			}

			if chunk.Error != nil {
				// h.logger.Printf("Volcengine stream error in chunk: code %s, message %s", chunk.Error.Code, chunk.Error.Message)
				return nil, fmt.Errorf("volcengine stream error: code %s, message: %s", chunk.Error.Code, chunk.Error.Message)
			}

			if finalResponseID == "" {
				finalResponseID = chunk.ID
			}

			if len(chunk.Choices) > 0 {
				choice := chunk.Choices[0]
				fullTextBuilder.WriteString(choice.Delta.Content)
				if choice.FinishReason != nil {
					finalFinishReason = *choice.FinishReason
				}
			}

			if chunk.Usage != nil {
				finalTokenUsage.PromptTokens = chunk.Usage.PromptTokens
				finalTokenUsage.CompletionTokens = chunk.Usage.CompletionTokens
				finalTokenUsage.TotalTokens = chunk.Usage.TotalTokens
			}
		}

		if err := scanner.Err(); err != nil {
			// h.logger.Printf("Error reading Volcengine stream: %v", err)
			return nil, fmt.Errorf("volcengine handler: error reading stream: %w", err)
		}

		// If usage was specifically requested via stream_options but not found in a chunk,
		// and the [DONE] message might have it. This part is tricky as [DONE] itself is not JSON.
		// The Volcengine Go SDK example suggests usage comes in the last data event *before* [DONE]
		// or sometimes within a chunk that also has delta content. The current loop structure should catch it if `include_usage` is true.

		return &models.TextGenerationResponse{
			ID:            finalResponseID,
			GeneratedText: fullTextBuilder.String(),
			FinishReason:  finalFinishReason,
			TokenUsage: struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				TotalTokens      int `json:"total_tokens"`
			}{
				PromptTokens:     finalTokenUsage.PromptTokens,
				CompletionTokens: finalTokenUsage.CompletionTokens,
				TotalTokens:      finalTokenUsage.TotalTokens,
			},
		}, nil

	} else {
		// 处理非流式响应
		respBodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			// h.logger.Printf("Error reading Volcengine response body: %v", err)
			return nil, fmt.Errorf("volcengine handler: failed to read response body: %w", err) // errors.Wrap(err, errors.ErrCodePlatformError, "volcengine handler: failed to read response body")
		}

		var volcResp volcengineChatResponse
		if err := json.Unmarshal(respBodyBytes, &volcResp); err != nil {
			// h.logger.Printf("Error unmarshalling Volcengine response: %v. Body: %s", err, string(respBodyBytes))
			return nil, fmt.Errorf("volcengine handler: failed to unmarshal response body: %w. Body: %s", err, string(respBodyBytes))
			// errors.Wrap(err, errors.ErrCodePlatformError, fmt.Sprintf("volcengine handler: failed to unmarshal response body. Body: %s\", string(respBodyBytes)))
		}

		if volcResp.Error != nil { // Check for API error in the non-stream response body
			return nil, fmt.Errorf("volcengine API error: code %s, message: %s", volcResp.Error.Code, volcResp.Error.Message)
		}

		if len(volcResp.Choices) == 0 {
			// h.logger.Println("Volcengine response contained no choices.")
			return nil, fmt.Errorf("volcengine handler: no choices found in response") // errors.New(errors.ErrCodePlatformError, "volcengine handler: no choices found in response")
		}

		choice := volcResp.Choices[0]

		sdkResp := &models.TextGenerationResponse{
			ID:            volcResp.ID,
			GeneratedText: choice.Message.Content,
			FinishReason:  choice.FinishReason,
			TokenUsage: struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				TotalTokens      int `json:"total_tokens"`
			}{
				PromptTokens:     volcResp.Usage.PromptTokens,
				CompletionTokens: volcResp.Usage.CompletionTokens,
				TotalTokens:      volcResp.Usage.TotalTokens,
			},
		}
		// h.logger.Printf("Successfully received and parsed TextGeneration response from Volcengine for ID: %s\", sdkResp.ID)
		return sdkResp, nil
	}
}
