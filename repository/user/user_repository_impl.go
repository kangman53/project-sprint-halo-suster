package user_repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepositoryImpl struct {
	DBpool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		DBpool: dbPool,
	}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, user user_entity.User) (user_entity.UserData, error) {
	var userId string
	query := "INSERT INTO users (name, nip, role, password, identity_card_scan_img) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	if err := repository.DBpool.QueryRow(ctx, query, user.Name, user.Nip, user.Role, user.Password, user.IdentityCardScanImg).Scan(&userId); err != nil {
		return user_entity.UserData{}, err
	}

	return user_entity.UserData{Id: userId}, nil
}

func (repository *userRepositoryImpl) Login(ctx context.Context, user user_entity.User) (user_entity.User, error) {
	query := "SELECT id, name, password, role FROM users WHERE nip = $1 AND is_deleted = false LIMIT 1"
	row := repository.DBpool.QueryRow(ctx, query, user.Nip)

	var loggedInUser user_entity.User
	err := row.Scan(&loggedInUser.Id, &loggedInUser.Name, &loggedInUser.Password, &loggedInUser.Role)
	if err != nil {
		return user_entity.User{}, err
	}

	return loggedInUser, nil
}

func (repository *userRepositoryImpl) Edit(ctx context.Context, user user_entity.User) error {
	query := "UPDATE users SET nip = $1, name = $2 WHERE id = $3 AND role = 'nurse' RETURNING id"
	if err := repository.DBpool.QueryRow(ctx, query, user.Nip, user.Name, user.Id).Scan(&user.Id); err != nil {
		return err
	}

	return nil
}

func (repository *userRepositoryImpl) Search(ctx context.Context, searchQuery user_entity.UserGetRequest) (*[]user_entity.UserResponseData, error) {
	query := `SELECT id, name, cast(nip as BIGINT) nip, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') createdAt FROM users WHERE is_deleted = false`
	var whereClause []string
	var searchParams []interface{}

	if searchQuery.Id != "" {
		whereClause = append(whereClause, fmt.Sprintf("id = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.Id)
	}
	if nip, _ := strconv.Atoi(searchQuery.Nip); nip > 0 {
		whereClause = append(whereClause, fmt.Sprintf("nip ~* $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.Nip)
	}
	if searchQuery.Name != "" {
		whereClause = append(whereClause, fmt.Sprintf("name ~* $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.Name)
	}
	if searchQuery.Role != "" {
		whereClause = append(whereClause, fmt.Sprintf("role = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.Role)
	}

	if len(whereClause) > 0 {
		query += " AND " + strings.Join(whereClause, " AND ")
	}

	query += " ORDER BY created_at"
	if strings.ToLower(searchQuery.CreatedAt) != "asc" {
		query += " DESC"
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", searchQuery.Limit, searchQuery.Offset)
	rows, err := repository.DBpool.Query(ctx, query, searchParams...)
	if err != nil {
		return &[]user_entity.UserResponseData{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[user_entity.UserResponseData])
	if err != nil {
		return &[]user_entity.UserResponseData{}, err
	}

	return &users, nil
}

func (repository *userRepositoryImpl) GiveAccess(ctx context.Context, user user_entity.User) (user_entity.User, error) {
	query := "UPDATE users SET password = $1 WHERE id = $2 AND role = 'nurse' RETURNING name, nip"
	if err := repository.DBpool.QueryRow(ctx, query, user.Password, user.Id).Scan(&user.Name, &user.Nip); err != nil {
		return user_entity.User{}, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Delete(ctx context.Context, userId string) (user_entity.User, error) {
	user := user_entity.User{
		Id:   userId,
		Role: "nurse",
	}
	query := "UPDATE users SET is_deleted = true WHERE id = $1 AND role = 'nurse' RETURNING name, nip"
	if err := repository.DBpool.QueryRow(ctx, query, user.Id).Scan(&user.Name, &user.Nip); err != nil {
		return user_entity.User{}, err
	}

	return user, nil
}
