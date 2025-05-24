# 平台特定信息

本文档提供与本 SDK 支持的特定 LLM 平台相关的详细信息和链接。

## 火山方舟 (Volcengine Ark)

*   **官方网站:** [https://www.volcengine.com/product/ark](https://www.volcengine.com/product/ark)
*   **API 文档:** [https://www.volcengine.com/docs/82379/1330310](https://www.volcengine.com/docs/82379/1330310) (由您提供)
*   **关键 API 端点 (待识别):**
    *   文本生成: `待定`
    *   身份验证: `待定`
    *   其他模型...
*   **身份验证说明:** (关于火山方舟如何处理身份验证的详细信息，例如 IAM、API 密钥等)

## 阿里百炼 (Alibaba Bailian)

*   **官方网站:** [https://bailian.aliyun.com/](https://bailian.aliyun.com/)
*   **API 文档:** 
    *   控制台/概览: [https://bailian.console.aliyun.com/](https://bailian.console.aliyun.com/) (您之前提供的链接)
    *   具体API文档: (我们需要查找百炼针对特定模型交互的直接 API 文档，例如文本生成、向量模型等)
    *   获取API Key: [https://help.aliyun.com/zh/model-studio/get-api-key](https://help.aliyun.com/zh/model-studio/get-api-key) (由您提供)
*   **关键 API 端点 (待识别):**
    *   文本生成: `待定`
    *   身份验证: `待定`
    *   其他模型...
*   **身份验证说明:** 阿里百炼平台主要通过 API Key 进行身份验证。开发者需要在阿里云控制台模型服务灵骏中创建并获取 API Key。

## 添加新平台

此处将概述开发人员如何通过添加对新平台的支持来做出贡献。这将涉及：

1.  定义新的 `Provider` (提供商) 常量。
2.  实现特定平台的 `Handler` (处理器) 接口。
3.  更新文档。 