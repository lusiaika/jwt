package database

import (
	"context"
	"database/sql"
	"log"
	"tugasregisterjwt/entity"
)

func (s *Database) Login(ctx context.Context, i string) (*entity.User, error) {
	result := &entity.User{}

	rows, err := s.SqlDb.QueryContext(ctx, "select id, username, email, password, age, createdat, updatedat from users where username = @username",
		sql.Named("username", i))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.Id,
			&result.Username,
			&result.Email,
			&result.Password,
			&result.Age,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return result, nil
}
