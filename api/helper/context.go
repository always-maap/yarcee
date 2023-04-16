package helper

import (
	"api/model"
	"context"
	"errors"
)

func RetrieveUser(c context.Context) (model.User, error) {
	user, ok := c.Value("user").(model.User)

	if !ok {
		return model.User{}, errors.New("could not cast user")
	}

	return user, nil
}
