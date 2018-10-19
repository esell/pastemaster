package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	hashids "github.com/speps/go-hashids"
)

type conf struct {
	DatabaseHost string
	DatabaseUser string
	DatabasePass string
}

var (
	configFile   = flag.String("c", "conf.json", "config file location")
	parsedconfig = conf{}
	useSSL       = flag.Bool("s", false, "enable SSL")
	salt         = "blah"
)

func main() {

	flag.Parse()
	readConfig()

	db, err := sql.Open("mysql", getConnString())
	if err != nil {
		log.Fatal("error opening db: ", err)
	}
	defer db.Close()

	// verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	// create DB if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS Pastes")
	if err != nil {
		log.Fatal("unable to create db: ", err)
	}

	_, err = db.Exec("USE Pastes")
	if err != nil {
		log.Fatal("unable to switch to db: ", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS p ( link varchar(30), content longtext )")
	if err != nil {
		log.Fatal("unable to create table: ", err)
	}

	http.Handle("/newPaste", newPasteHandler(db))
	http.Handle("/show", showPasteHandler(db))
	http.Handle("/", http.FileServer(assetFS()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readConfig() {
	if os.Getenv("PMDBHOST") == "" || os.Getenv("PMDBUSER") == "" || os.Getenv("PMDBPASS") == "" {
		fmt.Println("env vars not set, looking for config file...")

		file, err := ioutil.ReadFile(*configFile)
		if err != nil {
			log.Fatal("unable to read config file, exiting...")
		}
		if err := json.Unmarshal(file, &parsedconfig); err != nil {
			log.Fatal("unable to marshal config file, exiting...")
		}
	} else {
		parsedconfig.DatabaseHost = os.Getenv("PMDBHOST")
		parsedconfig.DatabaseUser = os.Getenv("PMDBUSER")
		parsedconfig.DatabasePass = os.Getenv("PMDBPASS")
	}
}

func newPasteHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			log.Println("error parsing form: ", err)
		}
		content := r.PostFormValue("message")

		// generate URL
		ref, err := createLink()
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}

		// insert into DB
		stmt, err := db.Prepare("INSERT INTO p VALUES(?, ?)")
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}
		res, err := stmt.Exec(ref, string(content))
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}
		rowCnt, err := res.RowsAffected()
		if err != nil || rowCnt < 1 {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}
		// return link
		w.Write([]byte(ref))
	})
}

func showPasteHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		recID := r.URL.Query().Get("id")

		stmt, err := db.Prepare("SELECT content FROM p WHERE link = ?")
		var pasteContent string
		err = stmt.QueryRow(recID).Scan(&pasteContent)
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}

		afsFile, err := assetFS().Open("show.html")
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}

		readafsFile, err := ioutil.ReadAll(afsFile)
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}

		t, err := template.New("foo").Parse(string(readafsFile))
		if err != nil {
			http.Error(w, "Yikes, something went wrong!", http.StatusInternalServerError)
			return
		}

		t.Execute(w, pasteContent)

	})
}

func getConnString() string {
	connString := parsedconfig.DatabaseUser + ":" + parsedconfig.DatabasePass + "@tcp(" + parsedconfig.DatabaseHost + ":3306)/"
	if *useSSL {

		return connString + "?tls=true"
	}

	return connString
}

func getRandomInt() int {
	var n *big.Int
	max := *big.NewInt(99999999999)
	n, err := rand.Int(rand.Reader, &max)

	if err != nil {
		log.Println("error creating randomInt: ", err)
	}
	return int(n.Int64())
}

func createLink() (string, error) {
	hid := hashids.NewData()
	hid.Salt = salt
	hid.MinLength = 30
	hashEncoder, err := hashids.NewWithData(hid)
	currentTime := time.Now().Nanosecond()
	e, err := hashEncoder.Encode([]int{currentTime, getRandomInt()})
	if err != nil {
		return "", err
	}
	return e, nil
}
