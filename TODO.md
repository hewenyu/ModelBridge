# TODO - 通用大模型平台 Go SDK 开发规划

本文档用于跟踪和规划通用大模型平台 Go SDK 的开发任务。

## 优先级 P0 - 核心功能与可用性

这些任务是确保 SDK 具备基本可用性的核心功能。

*   **[X] 完善项目结构和基础代码 (P0)**
    *   [X] 确定并创建核心包结构 (`client`, `models`, `platform`, `auth`, `utils`, `errors`)。
    *   [X] 实现 `client.NewClient()` 构造函数 (初步完成，Handler 初始化待具体平台实现)。
    *   [X] 定义核心的 `PlatformHandler` 接口。
    *   [X] 实现基本的错误处理机制 (`errors` 包) 和日志记录功能 (集成到 `Client`)。
    *   [X] 创建 `auth` 包的基础结构/占位符 (例如 `auth/auth.go`)。
    *   [X] 创建 `utils` 包的基础结构/占位符 (例如 `utils/utils.go`)。
*   **[ ] 完善平台支持 - 火山方舟 (P0)**
    *   [ ] 在 `platform/volcengine` 下创建 Handler 骨架。
    *   [ ] 实现火山方舟的身份验证逻辑 (根据 `GETTING_STARTED.md` 和官方文档)，可能在 `auth` 包或 `platform/volcengine` 中。
    *   [ ] 对接火山方舟的文本生成 API，并映射到 `models.TextGenerationRequest` 和 `models.TextGenerationResponse`。
    *   [ ] 对接火山方舟的图片生成 API (如果支持)，并映射到 `models.ImageGenerationRequest` 和 `models.ImageGenerationResponse`。
    *   [ ] 对接火山方舟的向量模型 API (如果支持)，并映射到 `models.EmbeddingRequest` 和 `models.EmbeddingResponse`。
    *   [ ] 添加火山方舟相关的单元测试和集成测试（需要模拟或真实凭证）。
*   **[ ] 完善平台支持 - 阿里百炼 (P0)**
    *   [ ] 在 `platform/alibaba` 下创建 Handler 骨架。
    *   [ ] 实现阿里百炼的身份验证逻辑 (根据 `GETTING_STARTED.md` 和官方文档)，可能在 `auth` 包或 `platform/alibaba` 中。
    *   [ ] 对接阿里百炼的文本生成 API，并映射到 `models.TextGenerationRequest` 和 `models.TextGenerationResponse`。
    *   [ ] 对接阿里百炼的图片生成 API (如果支持)，并映射到 `models.ImageGenerationRequest` 和 `models.ImageGenerationResponse`。
    *   [ ] 对接阿里百炼的向量模型 API (如果支持)，并映射到 `models.EmbeddingRequest` 和 `models.EmbeddingResponse`。
    *   [ ] 添加阿里百炼相关的单元测试和集成测试（需要模拟或真实凭证）。
*   **[ ] 完善文档 (P0)**
    *   [ ] 更新 `doc/README.md` 中的安装和快速开始示例 (待核心功能可用后)。
    *   [ ] 更新 `doc/GETTING_STARTED.md` 中的安装步骤和完整的认证配置示例 (待核心功能可用后)。
    *   [ ] 补充 `doc/PLATFORMS.md` 中火山方舟和阿里百炼的关键 API 端点信息 (在平台对接过程中进行)。
    *   [ ] 为已实现的核心功能和模型类型提供代码注释。
*   **[ ] 基础构建和测试 (P0)**
    *   [ ] 配置好 CI/CD 流程（例如 GitHub Actions），至少包含 `gofmt`, `go vet`, `go test ./...`。

## 优先级 P1 - 功能扩展与模型支持

在核心功能稳定后，可以扩展支持更多的模型类型和平台。

*   **[ ] 完善模型类型定义 (P1)**
    *   [ ] 调研并确定以下模型类型的通用请求/响应结构，并更新 `doc/MODEL_TYPES.md`：
        *   [ ] 多模态 (Multimodal)
        *   [ ] 推理模型 (Inference Models) - 考虑通用性
        *   [ ] 音频理解 (Audio Understanding) - 例如 ASR
        *   [ ] 语音合成 (Text-to-Speech / TTS)
    *   [ ] 为新定义的模型类型在 `PlatformHandler` 接口中添加相应方法。
*   **[ ] 实现流式响应 (P1)**
    *   [ ] 为文本生成等适用场景实现流式数据传输 (`Stream: true` 的处理)。
    *   [ ] 在平台 Handler 中实现对流式 API 的支持。
*   **[ ] 平台特定参数支持 (P1)**
    *   [ ] 确保 `PlatformSpecificParams` 能够正确传递和处理，允许用户覆盖或指定平台独有的参数。
*   **[ ] 扩展模型支持 - 火山方舟 (P1)**
    *   [ ] 根据 `doc/MODEL_TYPES.md` 中定义的其他模型类型，调研火山方舟是否支持，并实现对接。
*   **[ ] 扩展模型支持 - 阿里百炼 (P1)**
    *   [ ] 根据 `doc/MODEL_TYPES.md` 中定义的其他模型类型，调研阿里百炼是否支持，并实现对接。
*   **[ ] 示例代码 (P1)**
    *   [ ] 提供更丰富的示例代码，覆盖所有支持的模型类型和平台。
    *   [ ] 考虑创建一个 `examples` 目录。

## 优先级 P2 - 进阶功能与生态

*   **[ ] 进一步完善模型类型定义 (P2)**
    *   [ ] 调研并确定以下模型类型的通用请求/响应结构，并更新 `doc/MODEL_TYPES.md`：
        *   [ ] 视频理解 (Video Understanding)
        *   [ ] 视频生成 (Video Generation)
        *   [ ] 图片处理 (Image Processing)
        *   [ ] 图片理解 (Image Understanding)
        *   [ ] 排序模型 (Ranking Models)
*   **[ ] 支持更多平台 (P2)**
    *   [ ] 调研并选择下一个要支持的 LLM 平台。
    *   [ ] 按照 `doc/CONTRIBUTING.md` 中的指南添加新平台支持。
*   **[ ] 完善错误处理 (P2)**
    *   [ ] 在代码中全面使用 `errors` 包定义的结构化错误。
    *   [ ] 确保各平台返回的错误能够被合理地转换为通用 SDK `Error`。
*   **[ ] 配置管理 (P2)**
    *   [ ] 考虑更灵活的配置方式，例如从环境变量、配置文件加载凭证。
*   **[ ] 文档完善 (P2)**
    *   [ ] 撰写更详细的开发者文档，说明如何扩展和贡献。
    *   [ ] 补充 `CODE_OF_CONDUCT.md`。

## 待讨论/未来考虑

*   [ ] SDK 的版本管理策略。
*   [ ] 异步 API 支持（如果平台提供且有需求）。
*   [ ] 客户端级别的重试、超时机制。
*   [ ] 更高级的日志管理（例如集成流行的日志库 `logrus` 或 `zap`）。

---

请将此内容保存为项目根目录下的 `TODO.md` 文件。您可以根据项目的进展和实际需求随时更新此文档。 