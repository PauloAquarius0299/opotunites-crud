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
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Salary >= 0 { 
		return fmt.Errorf("request cannot be nil")
	}
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

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//if any field is provided, at least one must be non-empty
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Salary >= 0 && !r.Remote {
		return nil
	}
	// Validate each field if it is provided
	return fmt.Errorf("at least one field must be provided for update")
}
