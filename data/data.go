package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	// Database
	_ "github.com/lib/pq"
)

// Db : the variable to access the database connection
var Db *sql.DB

 const (
 	user     = "postgres"
 	password = "postgres"
 	dbname   = "labyrinth"
 )

func init() {
	var err error
	psqlInfo := fmt.Sprintf("user=%s "+"password=%s dbname=%s sslmode=disable",user, password, dbname)
	Db, err = sql.Open("postgres", psqlInfo)
	// Db, err = sql.Open("postgres", "dbname=labyrinth sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	dat, err := ioutil.ReadFile("data/setup.sql")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range strings.Split(string(dat), "\n\n") {
		v = strings.Replace(v, "\n", "", -1)
		fmt.Print(v)
		Db.Exec(v)
	}
	return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
