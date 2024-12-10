package converter

import (
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
)

// ToProtoEmail converts internal Email model to protobuf Email
func ToProtoEmail(email *models.Email) *pb.Email {
	if email == nil {
		return nil
	}
	return &pb.Email{
		Id:         email.ID,
		Subject:    email.Subject,
		Body:       email.Body,
		Sender:     email.Sender,
		Recipients: email.Recipients,
		Headers:    email.Headers,
	}
}

// FromProtoEmail converts protobuf Email to internal Email model
func FromProtoEmail(pbEmail *pb.Email) *models.Email {
	if pbEmail == nil {
		return nil
	}
	return &models.Email{
		ID:         pbEmail.Id,
		Subject:    pbEmail.Subject,
		Body:       pbEmail.Body,
		Sender:     pbEmail.Sender,
		Recipients: pbEmail.Recipients,
		Headers:    pbEmail.Headers,
	}
}

// ToProtoCategoryResult converts internal CategoryResult to protobuf CategoryResult
func ToProtoCategoryResult(result *models.CategoryResult) *pb.CategoryResult {
	if result == nil {
		return nil
	}
	return &pb.CategoryResult{
		Id:              result.EmailID,
		Categories:      result.Categories,
		ConfidenceScore: result.ConfidenceScore,
	}
}

// FromProtoCategoryResult converts protobuf CategoryResult to internal CategoryResult
func FromProtoCategoryResult(pbResult *pb.CategoryResult) *models.CategoryResult {
	if pbResult == nil {
		return nil
	}
	return &models.CategoryResult{
		EmailID:         pbResult.Id,
		Categories:      pbResult.Categories,
		ConfidenceScore: pbResult.ConfidenceScore,
	}
}