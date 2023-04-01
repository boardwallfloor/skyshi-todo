package main

import "time"

func getAllActivities() ([]Activity, error) {

	activity := []Activity{
		{
			ID:        1,
			Title:     "test",
			Email:     "test@mail.com",
			CreatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
			UpdatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
		},
	}
	return activity, nil
}

func getOneActivity(id int) (Activity, error) {

	activity := Activity{
		ID:        1,
		Title:     "test",
		Email:     "test@mail.com",
		CreatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
		UpdatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
	}
	return activity, nil
}

func createActivity(title, email string) (Activity, error) {
	activity := Activity{
		ID:        1,
		Title:     "test",
		Email:     "test@mail.com",
		CreatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
		UpdatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
	}
	return activity, nil
}

func updateActivity(id int, title, email string) (Activity, error) {
	activity := Activity{
		ID:        1,
		Title:     "test",
		Email:     "test@mail.com",
		CreatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
		UpdatedAt: time.Date(2022, 11, 28, 7, 22, 24, 0, time.UTC),
	}

	return activity, nil
}

func deleteActivity(id int) (Activity, error) {
	return Activity{}, nil
}
