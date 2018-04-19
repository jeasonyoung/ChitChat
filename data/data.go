package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"crypto/rand"
	"fmt"
	"crypto/sha1"
)

var Db *sql.DB

//init
func init(){
	var err error
	if Db,err = sql.Open("mysql", "root:root@tcp(120.78.209.83:3306)/chitchat?charset=utf8"); err != nil {
		log.Fatal(err)
	}
	return
}

//create a random UUID with from RFC 4122
//adapted from http://github.com/nu7hatch/gouuid
func createUUID()(uuid string){
	u := new([16]byte)
	if _, err := rand.Read(u[:]); err != nil {
		log.Fatalln("Cannot generate UUID ", err)
	}
	//0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	//Set the four most significant bits(bits 12 through 15) of the
	//time_hi_and_version field to the 4-bit version number
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string)(cryptext string){
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
