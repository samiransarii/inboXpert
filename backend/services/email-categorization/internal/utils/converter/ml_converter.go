package converter

import (
	mlpb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
)

// ToMLRequest converts an internal MLRequest model into the mlpb.EmailRequest
// type defined in the machine learning server's protobuf specification.
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

// FromMLResponse converts a CategoryResponse from the ML service (protobuf form)
// into the internal MLResponse model. It also maps any alternative categories.
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
