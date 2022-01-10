package cryptocurrency

import (
	"database/sql"
	"fmt"

	"github.com/xxdannilinxx/klv/utils"
)

type CryptoCurrencyRepository struct {
	db *sql.DB
}

var (
	tableName = "cryptocurrencies"
)

// New cryptocurrency repository office
func NewCryptoCurrencyRepository(db *sql.DB) *CryptoCurrencyRepository {
	return &CryptoCurrencyRepository{db}
}

// Get the most voted cryptocurrency from the repository
func (r *CryptoCurrencyRepository) GetMostVoted() (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	query := fmt.Sprintf(`SELECT * FROM "%s" ORDER BY "votes" DESC LIMIT 1;`, tableName)
	row := r.db.QueryRow(query)

	crypto := &CryptoCurrency{}

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
func (r *CryptoCurrencyRepository) GetById(id int64) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE "id" = $1;`, tableName)
	row := r.db.QueryRow(query, id)

	crypto := &CryptoCurrency{}

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
func (r *CryptoCurrencyRepository) Save(cc *CryptoCurrency) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	crypto := &CryptoCurrency{}

	query := fmt.Sprintf(`INSERT INTO "%s" ("name", "token") VALUES ($1, $2) RETURNING id, name, token, votes;`, tableName)
	err = r.db.QueryRow(query, cc.Name, cc.Token).Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)

	return crypto, err
}

// Change an existing cryptocurrency in the repository
func (r *CryptoCurrencyRepository) Update(cc *CryptoCurrency) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(cc.Id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "name" = $2, "token" = $3 WHERE "id" = $1;`, tableName)
	_, err = r.db.Exec(query, cc.Id, cc.Name, cc.Token)
	crypto = cc

	return crypto, err
}

// Delete an existing cryptocurrency from the repository
func (r *CryptoCurrencyRepository) Delete(id int64) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`DELETE FROM "%s" WHERE "id" = $1;`, tableName)
	_, err = r.db.Exec(query, id)

	return crypto, err
}

// Add a vote on an existing cryptocurrency in the repository
func (r *CryptoCurrencyRepository) UpVote(id int64) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "votes" = "votes" + 1 WHERE "id" = $1;`, tableName)
	_, err = r.db.Exec(query, id)

	return crypto, err
}

// Add a negative vote on an existing cryptocurrency in the repository
func (r *CryptoCurrencyRepository) DownVote(id int64) (*CryptoCurrency, error) {
	err := r.db.Ping()
	utils.CheckError(err)

	crypto, err := r.GetById(id)
	if err != nil {
		return crypto, err
	}

	query := fmt.Sprintf(`UPDATE "%s" SET "votes" = "votes" - 1 WHERE "id" = $1;`, tableName)
	_, err = r.db.Exec(query, id)

	return crypto, err
}
