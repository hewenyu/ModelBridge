package volcengine

import (
	"context"
	"fmt"

	"github.com/hewenyu/modelbridge/models"
)

// ImageGeneration 占位符实现
func (h *VolcengineHandler) ImageGeneration(ctx context.Context, req *models.ImageGenerationRequest) (*models.ImageGenerationResponse, error) {
	// h.logger.Println("Volcengine ImageGeneration not yet implemented.")
	return nil, fmt.Errorf("volcengine handler: ImageGeneration not yet implemented") // errors.New(errors.ErrCodeUnsupported, "ImageGeneration not yet implemented for Volcengine")
}
