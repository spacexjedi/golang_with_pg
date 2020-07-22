import (

	pg "github.com/go-pg/pg",
	"log",
	"db"
)

func connect() {

	opts:= &db.Options {

		User: "",
		Password: "",
		Address: "",
	}

	var db = db:= pg.Connect(opts)

	if db == nil {

		log.Printf("Failed to connect to database\n")
		os.Exit(100)
	}

		log.Printf("Connection to database successful\n")
		closeErr := db.Close()

	if closeErr != nil {

		log.Printf("Error while closing the connection, Reason\n", closeErr)
		os.Exit(100)
	}

		log.Printf("Connection closed successful\n")

		return

}