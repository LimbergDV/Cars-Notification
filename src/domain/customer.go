package domain

type Customer struct {
	Id int
	Name string
	Last_name string
	Phone_number string
	Curp string
	Number_license string
}

type Links struct {
	Self string
}

type Response struct {
	Data   []Customer
	Links  Links
	Available bool
}
