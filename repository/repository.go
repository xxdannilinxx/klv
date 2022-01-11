package repository

import (
	"database/sql"
	"fmt"

	"github.com/xxdannilinxx/klv/entity"
	"github.com/xxdannilinxx/klv/utils"
)

type CryptoCurrencyRepository interface {
	GetMostVoted() (*entity.CryptoCurrency, error)
	GetById(id int64) (*entity.CryptoCurrency, error)
	Save(cc *entity.CryptoCurrency) (*entity.CryptoCurrency, error)
	Update(cc *entity.CryptoCurrency) (*entity.CryptoCurrency, error)
	Delete(id int64) (*entity.CryptoCurrency, error)
	UpVote(id int64) (*entity.CryptoCurrency, error)
	DownVote(id int64) (*entity.CryptoCurrency, error)
}

type Repository struct{}

var (
	db        *sql.DB
	tableName = "cryptocurrencies"
)

// New cryptocurrency repository office
func NewCryptoCurrencyRepository(dbConn *sql.DB) *Repository {
	db = dbConn
	return &Repository{}
}

// Get the most voted cryptocurrency from the repository
func (r *Repository) GetMostVoted() (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	query := fmt.Sprintf(`SELECT * FROM "%s" ORDER BY "votes" DESC LIMIT 1;`, tableName)
	row := db.QueryRow(query)

	crypto := &entity.CryptoCurrency{}

	err = row.Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)
	switch err {
	case sql.ErrNoRows:
	case nil:
	default:
		utils.CheckError(err)
	}

	return crypto, err
}

// Get the cryptocurrency from the repository through the id
func (r *Repository) GetById(id int64) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE "id" = $1;`, tableName)
	row := db.QueryRow(query, id)

	crypto := &entity.CryptoCurrency{}

	err = row.Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)
	switch err {
	case sql.ErrNoRows:
	case nil:
	default:
		utils.CheckError(err)
	}

	return crypto, err
}

// Save a new cryptocurrency to the repository
func (r *Repository) Save(cc *entity.CryptoCurrency) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	crypto := &entity.CryptoCurrency{}

	query := fmt.Sprintf(`INSERT INTO "%s" ("name", "token") VALUES ($1, $2) RETURNING id, name, token, votes;`, tableName)
	err = db.QueryRow(query, cc.Name, cc.Token).Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)

	return crypto, err
}

// Change an existing cryptocurrency in the repository
func (r *Repository) Update(cc *entity.CryptoCurrency) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(cc.Id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "name" = $2, "token" = $3 WHERE "id" = $1;`, tableName)
	_, err = db.Exec(query, cc.Id, cc.Name, cc.Token)
	crypto = cc

	return crypto, err
}

// Delete an existing cryptocurrency from the repository
func (r *Repository) Delete(id int64) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`DELETE FROM "%s" WHERE "id" = $1;`, tableName)
	_, err = db.Exec(query, id)

	return crypto, err
}

// Add a vote on an existing cryptocurrency in the repository
func (r *Repository) UpVote(id int64) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "votes" = "votes" + 1 WHERE "id" = $1;`, tableName)
	_, err = db.Exec(query, id)

	return crypto, err
}

// Add a negative vote on an existing cryptocurrency in the repository
func (r *Repository) DownVote(id int64) (*entity.CryptoCurrency, error) {
	err := db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "votes" = "votes" - 1 WHERE "id" = $1;`, tableName)
	_, err = db.Exec(query, id)

	return crypto, err
}
