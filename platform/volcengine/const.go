package volcengine

type ModelID string

func (m ModelID) String() string {
	return string(m)
}

// TextGenerationModelID 是火山方舟平台支持的文本生成模型的ID。
type TextGenerationModelID ModelID

// Text Generation Model IDs for Volcengine
const (
	ModelDoubaoPro32k          TextGenerationModelID = "doubao-1.5-pro-32k-250115"
	ModelDoubaoPro32kCharacter TextGenerationModelID = "doubao-1.5-pro-32k-character-250228"
	ModelDoubaoPro256k         TextGenerationModelID = "doubao-1.5-pro-256k-250115"
	ModelDoubaoLite32k         TextGenerationModelID = "doubao-1.5-lite-32k-250115"
	ModelDeepseekV3_250324     TextGenerationModelID = "deepseek-v3-250324"
	ModelDeepseekV3_241226     TextGenerationModelID = "deepseek-v3-241226"
)

// ImageGenerationModelID 是火山方舟平台支持的图像生成模型的ID。
type ImageGenerationModelID ModelID

// Image Generation Model IDs for Volcengine
// doubao-seedream-3-0-t2i-250415
const (
	ModelDoubaoSeedream30T2i ImageGenerationModelID = "doubao-seedream-3-0-t2i-250415"
)

// 视频生成能力
type VideoGenerationModelID ModelID

// Video Generation Model IDs for Volcengine
/*
 * doubao-seedance-1-0-lite-t2v-250428 文生视频 图生视频-基于首帧  720p，480p 帧率：24 fps 时长：5 秒，10秒
 * wan2-1-14b-t2v-250225 文生视频 720p，480p 帧率：16 fps 时长：5 秒
 * wan2-1-14b-i2v-250225 图生视频-基于首帧  720p，480p 帧率：16 fps 时长：5 秒
 * wan2-1-14b-flf2v-250417 图生视频-基于首尾帧 720p 帧率：16 fps 时长：5 秒
 */
const (
	ModelDoubaoSeedance10LiteT2v VideoGenerationModelID = "doubao-seedance-1-0-lite-t2v-250428"
	ModelWan2114bT2v             VideoGenerationModelID = "wan2-1-14b-t2v-250225"
	ModelWan2114bI2v             VideoGenerationModelID = "wan2-1-14b-i2v-250225"
	ModelWan2114bFlf2v           VideoGenerationModelID = "wan2-1-14b-flf2v-250417"
)

// 文本向量化能力
type TextEmbeddingModelID ModelID

/*
 * doubao-embedding-large-text-240915 最高向量维度 4096
 * doubao-embedding-text-240715 最高向量维度 2560
 * doubao-embedding-text-240515 最高向量维度 2560
 */
const (
	ModelDoubaoEmbeddingLargeText240915 TextEmbeddingModelID = "doubao-embedding-large-text-240915"
	ModelDoubaoEmbeddingText240715      TextEmbeddingModelID = "doubao-embedding-text-240715"
	ModelDoubaoEmbeddingText240515      TextEmbeddingModelID = "doubao-embedding-text-240515"
)

// 图文向量化能力
type ImageTextEmbeddingModelID ModelID

/*
 * doubao-embedding-vision-250328
 * doubao-embedding-vision-241215
 */

// Image Text Embedding Model IDs for Volcengine
const (
	ModelDoubaoEmbeddingVision250328 ImageTextEmbeddingModelID = "doubao-embedding-vision-250328"
	ModelDoubaoEmbeddingVision241215 ImageTextEmbeddingModelID = "doubao-embedding-vision-241215"
)
