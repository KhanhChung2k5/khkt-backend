package services

import (
	"khkt-backend/models"
	"khkt-backend/request"
)

type ChemistryService interface {
	GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error)
	GetReferenceDocument(chemistry *request.GetRefDocument) ([]*models.ReferenceDocument, error)
	CreateChildren(typeChemistry string) []*MenuResponse
	GetMenu(req *request.GetMenu) ([]*MenuResponse, error)
}
