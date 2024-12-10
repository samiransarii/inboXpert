package converter

import (
	mlpb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
)

func ToMLRequest(req *models.MLRequest) *mlpb.EmailRequest {
	return &mlpb.EmailRequest{
		Id:         req.ID,
		Subject:    req.Subject,
		Body:       req.Body,
		Sender:     req.Sender,
		Recipients: req.Recipents,
		Headers:    req.Headers,
	}
}

func FromMLResponse(resp *mlpb.CategoryResponse) *models.MLResponse {
	alternatives := make([]models.Alternative, len(resp.Alternatives))
	for i, alt := range resp.Alternatives {
		alternatives[i] = models.Alternative{
			Category:        alt.Category,
			ConfidenceScore: alt.Confindence,
		}
	}

	return &models.MLResponse{
		ID:              resp.Id,
		Category:        resp.Category,
		ConfidenceScore: resp.Confidence,
		Alternatives:    alternatives,
	}
}
