package model

type UserDetails struct {
	Id        string `json:"id"`
	FullName  string `json:"fullName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	EmailId   string `json:"emailId"`
	UserName  string `json:"userName"`
	Role      string `json:"role"`
	Status    string `json:"status"`
}
