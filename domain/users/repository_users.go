package users

import (
	"log"

	"github.com/fernandojec/h_serviceuser/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewRepo_users(db *sqlx.DB) usersRepo {
	return usersRepo{db}
}

// saveRegistrasiUsers implements repo_users.
func (r usersRepo) saveRegistrasiUsers(data tableUsersModel) (user_id string, err error) {
	//query := "Insert into auths (email,password,username,first_name,last_name,is_active,created_at) values ($1,$2,$3,$4,$5,$6,$7) returning user_id"
	//query := "insert into users (user_id, first_name, last_name, password, email, is_active, user_create, create_at ) values ($1,$2,$3,$4,$5,true,'ricky',  NOW()) returning user_id"
	query := `insert into users (user_id, first_name, last_name, password, email, is_active, user_create, create_at ) 
	values ($1,$2,$3,$4,$5,true, 'ricky',  NOW()) `

	stmn, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}

	log.Printf("%+v", data)

	_, err = stmn.Exec(
		utils.NewSQLNullString(data.UserID),
		utils.NewSQLNullString(data.FirstName),
		utils.NewSQLNullString(data.LastName),
		utils.NewSQLNullString(data.Password),
		utils.NewSQLNullString(data.Email),
	)
	if err != nil {
		log.Printf("%+v", err)
		log.Printf("%+v", query)
		return "", err
	}

	log.Printf("%+v", data.UserID)
	user_id = data.UserID
	return
}

func (r usersRepo) getbyUserID(sUserID string) (data verifyUsers, err error) {
	query := `select user_id, first_name, last_name from Users where user_id=$1`

	_, err = r.db.Prepare(query)
	if err != nil {
		return verifyUsers{}, err
	}
	err = r.db.Get(&data, query, sUserID)

	return
}
