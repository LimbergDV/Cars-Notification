package repositroy

import "cars_notify/src/domain"

type IFetchAPI interface {
	FetchAPI(id_customer int) (domain.Response, error)
}