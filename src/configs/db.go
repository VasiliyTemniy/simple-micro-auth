package configs

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

func initializeTokenHandler() {
	tokenHandler = NewTokenHandler()
}

var tokenHandler ITokenHandler

func init() {
	initializeTokenHandler()
}

type dbHandler interface {
	CreateAuth(auth AuthDTO) *AuthResponse
	UpdateAuth(auth AuthDTOUpdate) *AuthResponse
	DeleteAuth(lookupHash string) error
	CompareAuth(auth AuthDTO) *AuthResponse
	checkAuth(auth AuthDTO) bool
}

type dbHandlerImpl struct {
	db *sql.DB
}

// CreateAuth implements dbHandler.
func (handler *dbHandlerImpl) CreateAuth(auth AuthDTO) *AuthResponse {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), EnvBcryptCost)

	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with creating this " + auth.Password + " password hash with bcrypt: " + err.Error(),
		}
	}

	_, err = handler.db.Exec(`INSERT INTO auth(lookup_hash, password_hash)
	VALUES($1, $2)`, auth.LookupHash, passwordHash)

	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with creating new auth: " + err.Error(),
		}
	}

	tokenString, err := tokenHandler.CreateToken(auth.Id)
	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with signing token: " + err.Error(),
		}
	}

	return &AuthResponse{
		Id:    auth.Id,
		Token: tokenString,
		Error: "",
	}
}

// UpdateAuth implements dbHandler.
func (handler *dbHandlerImpl) UpdateAuth(auth AuthDTOUpdate) *AuthResponse {

	authDTOToCompare := AuthDTO{
		Id:         auth.Id,
		LookupHash: auth.LookupHash,
		Password:   auth.OldPassword,
	}

	if !handler.checkAuth(authDTOToCompare) {
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Incorrect password",
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.NewPassword), EnvBcryptCost)
	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with creating this " + auth.NewPassword + " password hash with bcrypt: " + err.Error(),
		}
	}

	_, err = handler.db.Exec(`UPDATE auth SET password_hash = $1 WHERE lookup_hash = $2`, passwordHash, auth.LookupHash)
	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with updating password_hash in db while this lookup_hash was found: " + auth.LookupHash + " error: " + err.Error(),
		}
	}

	tokenString, err := tokenHandler.CreateToken(auth.Id)
	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with signing token: " + err.Error(),
		}
	}

	return &AuthResponse{
		Id:    auth.Id,
		Token: tokenString,
		Error: "",
	}
}

// DeleteAuth implements dbHandler.
func (handler *dbHandlerImpl) DeleteAuth(lookupHash string) error {

	_, err := handler.db.Exec(`DELETE FROM auth WHERE lookup_hash = $1`, lookupHash)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

// CompareAuth implements dbHandler.
func (handler *dbHandlerImpl) CompareAuth(auth AuthDTO) *AuthResponse {

	if !handler.checkAuth(auth) {
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Incorrect password",
		}
	}

	tokenString, err := tokenHandler.CreateToken(auth.Id)
	if err != nil {
		log.Println(err)
		return &AuthResponse{
			Id:    auth.Id,
			Token: "",
			Error: "Problem with signing token: " + err.Error(),
		}
	}

	return &AuthResponse{
		Id:    auth.Id,
		Token: tokenString,
		Error: "",
	}
}

// CheckAuth implements dbHandler.
func (handler *dbHandlerImpl) checkAuth(auth AuthDTO) bool {

	storedPasswordHashRow, err := handler.db.Query(`SELECT password_hash FROM auth WHERE lookup_hash = $1`, auth.LookupHash)
	if err != nil {
		log.Println(err)
		return false
	}

	var storedPasswordHash string

	if storedPasswordHashRow.Next() {
		err = storedPasswordHashRow.Scan(&storedPasswordHash)
		if err != nil {
			log.Println(err)
			return false
		}
	}

	return bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(auth.Password)) == nil
}

func NewDBHandler() dbHandler {

	postgresConfig := EnvPostgresConfig

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.DBName,
		postgresConfig.SslMode,
	)

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatal("Error opening db")
		log.Fatal("DB connection string: ", postgresqlDbInfo)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging db")
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS auth (
		lookup_hash varchar(100) UNIQUE NOT NULL PRIMARY KEY,
		password_hash varchar(100) NOT NULL
	)`)
	if err != nil {
		log.Fatal("Error creating auth table")
		panic(err)
	} else {
		log.Println("auth table exists or created")
	}

	return &dbHandlerImpl{db}
}
