package main

import "github.com/nu7hatch/gouuid"

func generateGuid() (*uuid.UUID, error) {
	u, err := uuid.NewV4()
	return u, err
}

func stringGuid(u *uuid.UUID) string {
	return u.String()
}
