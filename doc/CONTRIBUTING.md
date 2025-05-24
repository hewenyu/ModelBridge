# 贡献指南

感谢您对本项目有兴趣并考虑做出贡献！我们欢迎各种形式的贡献，包括但不限于：

*   报告 Bug
*   提交功能建议
*   编写和改进文档
*   提交代码修复和新功能

## 如何贡献

### 报告 Bug

如果您在项目中发现了 Bug，请通过提交一个 Issue 来报告它。在您的 Issue 中，请尽可能详细地描述以下信息：

*   您使用的 Go 版本。
*   您使用的 SDK 版本。
*   清晰描述 Bug 的复现步骤。
*   您期望的结果是什么。
*   实际发生了什么，包括任何错误信息和堆栈跟踪。

### 提交功能建议

如果您有关于新功能或改进现有功能的建议，也请通过提交一个 Issue 来进行讨论。请在建议中阐述：

*   您希望解决的问题或达成的目标。
*   您的具体建议和它可能带来的好处。
*   （可选）您对如何实现这个功能的初步想法。

### 提交合并请求 (Pull Requests)

我们非常欢迎代码贡献！如果您希望提交代码，请遵循以下步骤：

1.  **Fork 本仓库**：点击仓库页面右上角的 "Fork" 按钮。
2.  **Clone 您的 Fork**：`git clone https://github.com/hewenyu/modelbridge.git` (请将 `hewenyu` 替换为您的 GitHub 用户名)。
3.  **创建新分支**：`git checkout -b feature/your-feature-name` 或 `bugfix/issue-number`。
4.  **进行修改**：按照您的想法进行代码修改。请确保遵循下述的编码风格和测试要求。
5.  **提交您的修改**：`git commit -m "feat: 描述您的修改"` (请参考 [Conventional Commits](https://www.conventionalcommits.org/) 规范进行提交信息的编写)。
6.  **Push 到您的 Fork**：`git push origin feature/your-feature-name`。
7.  **创建 Pull Request**：回到您 Fork 的 GitHub 仓库页面，点击 "New pull request" 按钮，并选择合适的分支进行比较和提交。

## 开发环境设置

*   **Go 版本**：请确保您的 Go 版本与 `go.mod` 文件中指定的版本一致或更新。 (当前 `go.mod` 指定为 `go 1.23.6`)
*   **依赖管理**：本项目使用 Go Modules 管理依赖。您可以通过 `go get` 或 `go mod tidy` 来管理依赖项。

## 编码风格

*   请遵循 Go 语言的官方编码规范 ([Effective Go](https://go.dev/doc/effective_go))。
*   使用 `gofmt` 或 `goimports` 格式化您的代码。
*   代码应清晰、易读、可维护，并添加必要的注释。
*   借鉴 `go-general-expert` 规则中提到的设计原则，关注代码的可扩展性、可靠性、可维护性和安全性。

## 测试

*   对于任何代码更改，请添加相应的单元测试。
*   确保所有测试都通过 (`go test ./...`)。
*   对于核心功能的修改或新平台的添加，可能需要集成测试。

## 添加新的大模型平台支持

如果您希望为 SDK 添加对新的大模型平台的支持，通常需要以下步骤：

1.  在 `client` (或其他相关包) 中定义新的 `Provider` 常量。
2.  实现一个新的平台特定的 `Handler` 接口，该接口将处理与新平台 API 的所有交互（认证、请求构建、响应解析等）。
3.  根据新平台支持的模型类型，在 `models` 包中可能需要适配或扩展现有的请求/响应结构体，或者定义新的结构体。
4.  更新 `doc/PLATFORMS.md` 文档，添加关于新平台的信息、API 文档链接和认证说明。
5.  更新 `doc/MODEL_TYPES.md` 文档（如果适用），说明新平台如何支持各种模型类型。
6.  添加相关的单元测试和（如果可能）集成测试。

## 行为准则

我们致力于为所有贡献者和用户提供一个友好、尊重和无骚扰的环境。所有参与本项目的人员都应遵守[贡献者行为准则](LINK_TO_CODE_OF_CONDUCT.md) (请替换为实际的行为准则文件链接，如果暂时没有，可以先留空或指向一个通用的模板，如 Contributor Covenant)。

感谢您的贡献！ 