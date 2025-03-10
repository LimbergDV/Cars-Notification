package repositroy

type INotify interface {
	NotifyOfRent(msg string, phone_number string, return_date string, name string)
	NotifyOfReturn(msg string, phone_number string, name string)
}