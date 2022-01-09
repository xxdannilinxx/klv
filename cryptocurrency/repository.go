package cryptocurrency

import (
	"database/sql"

	"github.com/xxdannilinxx/klv/utils"
)

type CryptoCurrencyRepository struct {
	db *sql.DB
}

func (r *CryptoCurrencyRepository) GetMostVoted() (*CryptoCurrency, error) {
	row := r.db.QueryRow(`SELECT * FROM "cryptocurrencies" ORDER BY "votes" DESC LIMIT 1;`)

	crypto := &CryptoCurrency{}

	err := row.Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)
	switch err {
	case sql.ErrNoRows:
	case nil:
	default:
		utils.CheckError(err)
	}

	return crypto, err
}

func (r *CryptoCurrencyRepository) Get(id int64) (*CryptoCurrency, error) {
	row := r.db.QueryRow(`SELECT * FROM "cryptocurrencies" WHERE "id" = $1;`, id)

	crypto := &CryptoCurrency{}

	err := row.Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)
	switch err {
	case sql.ErrNoRows:
	case nil:
	default:
		utils.CheckError(err)
	}

	return crypto, err
}

func (r *CryptoCurrencyRepository) Insert(cc *CryptoCurrency) (*CryptoCurrency, error) {
	crypto := &CryptoCurrency{}
	err := r.db.QueryRow(`INSERT INTO "cryptocurrencies" ("name", "token") VALUES ($1, $2) RETURNING id, name, token, votes;`, cc.Name, cc.Token).Scan(&crypto.Id, &crypto.Name, &crypto.Token, &crypto.Votes)
	return crypto, err
}

func (r *CryptoCurrencyRepository) Update(cc *CryptoCurrency) (*CryptoCurrency, error) {
	crypto, err := r.Get(cc.Id)
	if err != nil {
		return crypto, err
	}

	_, err = r.db.Exec(`UPDATE "cryptocurrencies" SET "name" = $1, "token" = $2;`, cc.Name, cc.Token)
	crypto = cc

	return crypto, err
}

func (r *CryptoCurrencyRepository) Delete(id int64) (*CryptoCurrency, error) {
	crypto, err := r.Get(id)
	if err != nil {
		return crypto, err
	}

	_, err = r.db.Exec(`DELETE FROM "cryptocurrencies" WHERE "id" = $1;`, id)

	return crypto, err
}

func (r *CryptoCurrencyRepository) UpVote(id int64) (*CryptoCurrency, error) {
	crypto, err := r.Get(id)
	if err != nil {
		return crypto, err
	}

	_, err = r.db.Exec(`UPDATE "cryptocurrencies" SET "votes" = "votes" + 1 WHERE "id" = $1;`, id)

	return crypto, err
}

func (r *CryptoCurrencyRepository) DownVote(id int64) (*CryptoCurrency, error) {
	crypto, err := r.Get(id)
	if err != nil {
		return crypto, err
	}

	_, err = r.db.Exec(`UPDATE "cryptocurrencies" SET "votes" = "votes" - 1 WHERE "id" = $1;`, id)

	return crypto, err
}
