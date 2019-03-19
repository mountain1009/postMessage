package main

import (
    // "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


func gormConnect() *gorm.DB {
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "ryohei1009"
    DBNAME   := "postMessage"

    var CONNECT = USER+":"+PASS+"@/"+DBNAME
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
      panic(err.Error())
    }
    return db
}

type User struct {
	gorm.Model
	Name string
}

func main() {
    db := gormConnect()
    db.CreateTable(&User{})
    db.AutoMigrate(&User{})
    defer db.Close()
    db.LogMode(true)
}

// func gormConnect() *gorm.DB {
//     DBMS     := "mysql"
//     USER     := "root"
//     PASS     := "ryohei1009"
//     DBNAME   := "postMessage"

//     CONNECT = USER+":"+PASS+"@/"+DBNAME
//     db,err := gorm.Open(DBMS, CONNECT)

//     if err != nil {
//       panic(err.Error())
//     }
//     return db
// }