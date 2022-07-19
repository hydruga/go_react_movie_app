package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get Returns one movie and error if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating,
		created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// get the genres
	genres := make(map[int]string)
	query = `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name 
				from movies_genres as mg 
				left join genres g on (g.id = mg.genre_id)
				where
				mg.movie_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName

	}
	movie.MovieGenre = genres

	return &movie, nil
}

// GetAll returns all movies and error if any
func (m *DBModel) GetAll() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating,
	created_at, updated_at from movies order by title`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		genreQuery := `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name 
				from movies_genres as mg 
				left join genres g on (g.id = mg.genre_id)
				where
				mg.movie_id = $1`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)
		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			genreErr := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if genreErr != nil {
				return nil, err
			}
			genres[mg.ID] = mg.Genre.GenreName

		}
		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}
