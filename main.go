package main

import (
	"log"
	"net/http"
	"os"

	"saturday/domain"
	"saturday/handlers"
	"saturday/infra"
	"saturday/interfaces"
	"saturday/repo"
	"saturday/utils"

	"github.com/alexsasharegan/dotenv"
	"github.com/gorilla/mux"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Conf := domain.SectionService{
		DB: domain.DBAccount{
			Username:     os.Getenv("DATABASE_READ_USERNAME"),
			Password:     os.Getenv("DATABASE_READ_PASSWORD"),
			URL:          os.Getenv("DATABASE_READ_URL"),
			Port:         os.Getenv("DATABASE_READ_PORT"),
			DBName:       os.Getenv("DATABASE_READ_DB_NAME"),
			MaxIdleConns: utils.GetInt(os.Getenv("DATABASE_READ_MAXIDLECONNS")),
			MaxOpenConns: utils.GetInt(os.Getenv("DATABASE_READ_MAXOPENCONNS")),
			MaxLifeTime:  utils.GetInt(os.Getenv("DATABASE_READ_MAXLIFETIME")),
			Timeout:      os.Getenv("DATABASE_READ_TIMEOUT"),
		},
	}

	db := infra.DBHandler{}
	db.ConnectDB(&Conf.DB)

	dbList := make(map[string]interfaces.Database)
	dbList["DB"] = &db

	orgsRepo := &repo.OrgsRepo{
		DB: dbList["DB"],
	}

	orgsHandler := &handlers.OrgsHandler{
		Repo: orgsRepo,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/orgs/{orgs}/comments", orgsHandler.SaveCommentHandler).Methods(http.MethodPost)
	router.HandleFunc("/orgs/{orgs}/comments", orgsHandler.GetCommentHandler).Methods(http.MethodGet)
	router.HandleFunc("/orgs/{orgs}/comments", orgsHandler.DeleteCommentHandler).Methods(http.MethodDelete)

	router.HandleFunc("/orgs/{orgs}/members", orgsHandler.GetMemberHandler).Methods(http.MethodGet)

	log.Println("server listen to port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
