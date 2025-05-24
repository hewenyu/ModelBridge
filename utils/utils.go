// utils/utils.go
package utils

import (
	"fmt"
	"net/http"
	// "YOUR_PROJECT_PATH/errors" // 考虑引入自定义错误
	// "YOUR_PROJECT_PATH/client" // 可能需要 Logger
)

// HTTPClient 是一个简单的 HTTP 客户端接口，方便测试时 mock。
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient 是 HTTPClient 接口的默认实现，使用 http.DefaultClient。
var DefaultHTTPClient HTTPClient = http.DefaultClient

// MakeHTTPRequest 是一个辅助函数，用于创建和执行 HTTP 请求。
// method: HTTP 方法 (GET, POST, etc.)
// url: 请求的 URL
// headers: 请求头 (可以为 nil)
// requestBody: 请求体 (可以为 nil, 对于 GET 请求通常为 nil)
// responseBody: 用于 unmarshal 响应体的目标结构体指针 (如果不需要解析响应体，可以为 nil)
// client: 用于执行请求的 HTTPClient (如果为 nil, 使用 DefaultHTTPClient)
// logger: (可选) 用于记录日志的 logger
// func MakeHTTPRequest(
// 	ctx context.Context,
// 	method string,
// 	url string,
// 	headers map[string]string,
// 	requestBody interface{},
// 	responseBody interface{},
// 	httpClient HTTPClient,
// 	// logger client.Logger, // 暂时注释，避免循环依赖
// ) error {
// 	if httpClient == nil {
// 		httpClient = DefaultHTTPClient
// 	}

// 	var reqBodyBytes []byte
// 	var err error

// 	if requestBody != nil {
// 		reqBodyBytes, err = json.Marshal(requestBody)
// 		if err != nil {
// 			// if logger != nil { logger.Printf("Error marshalling request body: %v", err) }
// 			return fmt.Errorf("failed to marshal request body: %w", err) // errors.Wrap(err, errors.ErrCodeInternal, "failed to marshal request body")
// 		}
// 	}

// 	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBodyBytes))
// 	if err != nil {
// 		// if logger != nil { logger.Printf("Error creating new HTTP request: %v", err) }
// 		return fmt.Errorf("failed to create HTTP request: %w", err) // errors.Wrap(err, errors.ErrCodeInternal, "failed to create HTTP request")
// 	}

// 	// 设置请求头
// 	if len(reqBodyBytes) > 0 {
// 		req.Header.Set("Content-Type", "application/json")
// 	}
// 	req.Header.Set("Accept", "application/json") // 通常期望 JSON 响应
// 	for k, v := range headers {
// 		req.Header.Set(k, v)
// 	}

// 	// if logger != nil { logger.Printf("Making HTTP request: %s %s", method, url) }
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		// if logger != nil { logger.Printf("Error performing HTTP request to %s: %v", url, err) }
// 		return fmt.Errorf("HTTP request failed: %w", err) // errors.Wrap(err, errors.ErrCodePlatformError, "HTTP request failed")
// 	}
// 	defer resp.Body.Close()

// 	// if logger != nil { logger.Printf("Received HTTP response: %s %s, Status: %s", method, url, resp.Status) }

// 	respBodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		// if logger != nil { logger.Printf("Error reading response body from %s: %v", url, err) }
// 		return fmt.Errorf("failed to read response body: %w", err) // errors.Wrap(err, errors.ErrCodePlatformError, "failed to read response body")
// 	}

// 	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 		// if logger != nil { logger.Printf("HTTP Error: Status %d, Body: %s", resp.StatusCode, string(respBodyBytes)) }
// 		// 可以尝试解析平台特定的错误结构
// 		// platformErr := errors.New(errors.ErrCodePlatformError, fmt.Sprintf("platform request failed with status %d", resp.StatusCode))
// 		// platformErr.PlatformDetails = map[string]interface{}{"status_code": resp.StatusCode, "response_body": string(respBodyBytes)}
// 		// return platformErr
// 		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBodyBytes))
// 	}

// 	if responseBody != nil && len(respBodyBytes) > 0 {
// 		if err := json.Unmarshal(respBodyBytes, responseBody); err != nil {
// 			// if logger != nil { logger.Printf("Error unmarshalling response body from %s: %v. Body: %s", url, err, string(respBodyBytes)) }
// 			return fmt.Errorf("failed to unmarshal response body: %w. Body: %s", err, string(respBodyBytes)) // errors.Wrap(err, errors.ErrCodePlatformError, "failed to unmarshal response body")
// 		}
// 	}

// 	return nil
// }

// TODO:
// - 完善 MakeHTTPRequest 函数，包括错误处理 (使用自定义 errors 包)、日志记录的集成。
// - 考虑添加重试逻辑、超时控制等。
// - 提供更细致的 HTTP 错误解析。
// - 添加其他有用的工具函数，例如生成 UUID, 处理时间等。

// Placeholder function to avoid empty package issues if MakeHTTPRequest is fully commented out
func Placeholder() {
	fmt.Println("Utils placeholder")
}
