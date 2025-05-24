package volcengine

import (
	"context"
	"fmt"

	"github.com/hewenyu/modelbridge/models"
)

// Embedding 占位符实现
func (h *VolcengineHandler) Embedding(ctx context.Context, req *models.EmbeddingRequest) (*models.EmbeddingResponse, error) {
	// h.logger.Println("Volcengine Embedding not yet implemented.")
	return nil, fmt.Errorf("volcengine handler: Embedding not yet implemented") // errors.New(errors.ErrCodeUnsupported, "Embedding not yet implemented for Volcengine")
}
