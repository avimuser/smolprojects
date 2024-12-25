package cache

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	Header map[string][]string
	Body   []byte
}

func New() Cache {
	return Cache{
		Header: make(map[string][]string),
	}
}

func expect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./cache.db")
	expect(err)
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS cache(url TEXT, header BLOB, body BLOB)", nil)
	expect(err)
}

func Store(URL string, cache Cache) {
	header, err := json.Marshal(cache.Header)
	expect(err)
	_, err = DB.Exec("INSERT INTO cache VALUES (?, ?, ?)", URL, header, cache.Body)
	expect(err)
}

func Get(URL string) (cache Cache, cached bool) {
	headerRow := DB.QueryRow("SELECT header, body FROM cache WHERE url = ?", URL)
	var header []byte

	err := headerRow.Scan(&header, &cache.Body)

	if err == sql.ErrNoRows {
		return New(), false
	} else if err != nil {
		expect(err)
	}

	err = json.Unmarshal(header, &cache.Header)
	expect(err)

	return cache, true
}

func Clear() {
	Init()
	DB.Exec("DROP TABLE IF EXISTS cache", nil)
	DB.Close()
}
