package services

import (
	"cars_notify/src/application/repositroy"
	"cars_notify/src/domain"
)


type FetchAPIService struct {
	f repositroy.IFetchAPI
}

func NewFetchAPIService (f repositroy.IFetchAPI) *FetchAPIService {
	return &FetchAPIService{f:f}
}

func (s *FetchAPIService) Run(id_customer int) (domain.Response, error) {
	return s.f.FetchAPI(id_customer)
}