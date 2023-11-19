package repositories

import (
	"database/sql"
	"log"
	"os"

	"github.com/billzayy/elastic-golang/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() ([]models.User, error) {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DBPath"))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var userList []models.User

	for rows.Next() {
		var list models.User

		err := rows.Scan(&list.Id, &list.UserName, &list.Age)

		if err != nil {
			return nil, err
		}

		userList = append(userList, list)
	}
	return userList, nil
}

func CreateIndex() {

}
