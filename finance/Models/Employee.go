package Models

import "errors"

func (e *Employee) Validate() error {

	if e.Name == "" {
		return errors.New("name is required")
	}

	if e.Age < 18 || e.Age > 100 {
		return errors.New("invalid age")
	}

	return nil
}

func GetAllEmployee(e *[]Employee) (err error) {
	return
}
