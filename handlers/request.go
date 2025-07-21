package handlers

import "fmt"

func errParmIsRequired(name, typ string) error {
	return fmt.Errorf("parameter %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParmIsRequired("Role", "string")
	}
	if r.Company == "" {
		return errParmIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParmIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParmIsRequired("Link", "string")
	}
	if r.Salary <= 0 {
		return errParmIsRequired("Salary", "int64")
	}
	if r.Remote {
		return errParmIsRequired("Remote", "bool")
	}
	return nil
}