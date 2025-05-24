# 快速入门指南

本文档提供详细说明，指导您如何设置并开始使用通用大模型平台 Go SDK。

## 先决条件

*   Go (版本 X.Y.Z 或更高版本 - 请指定您计划支持的版本，例如，与 `go.mod` 中的版本一致 `1.23.6`)
*   目标 LLM 平台的访问凭证 (API 密钥等)

## 安装

详细的安装步骤将在此处提供。目前，假设您将使用 `go get`。

```bash
# (仓库公开后待添加)
go get github.com/hewenyu/modelbridge # 占位符，请替换为您的实际仓库地址
```

## 身份验证

每个平台都需要特定的身份验证机制。SDK 提供了一种统一的方式来配置这些机制。

### 火山方舟 (Volcengine Ark)

*   **认证方法:** (例如：API 密钥, Access Key/Secret Key)
*   **配置示例:**
    ```go
    // 火山方舟配置示例
    volcConfig := &client.PlatformConfig{
        Provider: client.ProviderVolcengine, // 假设 client 包中定义了 ProviderVolcengine
        Credentials: map[string]string{
            "apiKey": "YOUR_VOLCENGINE_API_KEY", // 或其他相关密钥
            // "secretKey": "YOUR_VOLCENGINE_SECRET_KEY",
            // "accessKey": "YOUR_VOLCENGINE_ACCESS_KEY",
        },
    }
    ```

### 阿里百炼 (Alibaba Bailian)

*   **认证方法:** API 密钥 (API Key)。请参考[如何获取API Key](https://help.aliyun.com/zh/model-studio/get-api-key?spm=a2c4g.11186623.0.0.78d84823OWXAx8)进行获取。
*   **配置示例:**
    ```go
    // 阿里百炼配置示例
    aliConfig := &client.PlatformConfig{
        Provider: client.ProviderAlibaba, // 假设 client 包中定义了 ProviderAlibaba
        Credentials: map[string]string{
            "apiKey": "YOUR_ALIBABA_BAILIAN_API_KEY",
        },
    }
    ```

## 基本用法

以下是如何使用 SDK 的概念性示例。具体的模型交互示例将在 `MODEL_TYPES.md` (模型类型文档) 中提供。

```go
package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/hewenyu/modelbridge/client" // 占位符
	// "github.com/hewenyu/modelbridge/models" // 占位符
)

func main() {
	ctx := context.Background()

	// 初始化火山方舟客户端 (示例)
	/*
	volcConfig := &client.PlatformConfig{
	    Provider: client.ProviderVolcengine,
	    Credentials: map[string]string{
	        "apiKey": "YOUR_VOLCENGINE_API_KEY",
	    },
	}
	volcClient, err := client.NewClient(volcConfig) // NewClient 是假设的构造函数
	if err != nil {
		log.Fatalf("创建火山方舟客户端失败: %v", err)
	}

	// 示例：使用火山方舟进行文本生成 (概念性)
	textGenRequest := &models.TextGenerationRequest{
		Prompt: "将 'hello world' 翻译成法语。",
		Model:  "volcengine-specific-model-id", // 或使用通用模型别名
	}
	resp, err := volcClient.TextGeneration(ctx, textGenRequest) // TextGeneration 是假设的客户端方法
	if err != nil {
		log.Fatalf("文本生成失败: %v", err)
	}
	fmt.Printf("火山方舟响应: %s\n", resp.GeneratedText)
	*/

	fmt.Println("SDK 快速入门 - 更多示例敬请期待！")
}
``` 