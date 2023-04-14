package helper

import (
	"api/models"
	"context"
	"errors"
)

func RetrieveUser(c context.Context) (models.User, error) {
	user, ok := c.Value("user").(models.User)

	if !ok {
		return models.User{}, errors.New("could not cast user")
	}

	return user, nil
}
