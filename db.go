package main

import(
     //"github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
      //"database/sql"
      "log"
      "fmt"
)




//DB_Migration Миграция БД
func DB_Migration() {
    db, err := gorm.Open("postgres", fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",
      DB_PORT,DB_USER, DB_PASSWORD, DB_NAME))  
    if err != nil {
        log.Fatalf("error: %v\n", err)
        //return
    } 
    defer db.Close()

    //db.LogMode(true)
    db.AutoMigrate(
        User{},
        Region{},
        //Question{},
        //Variant{},
        //QuestionComment{},
        AutoSale{},
        Session{})

    //db.Model(&Variant{}).AddForeignKey("Question_id", "Questions (id)", "CASCADE", "CASCADE")
    //db.Model(&Question{}).AddForeignKey("User_id", "Users (id)", "CASCADE", "CASCADE")
    //db.Model(&QuestionComment{}).AddForeignKey("Question_id", "Questions (id)", "CASCADE", "CASCADE")
    //db.Model(&QuestionComment{}).AddForeignKey("User_id", "Users (id)", "CASCADE", "CASCADE")
    //db.Model(&DataBase{}).AddForeignKey("User_id", "Users (id)", "CASCADE", "CASCADE")
    //db.Model(&Session{}).AddForeignKey("User_id", "Users (id)", "CASCADE", "CASCADE")

    //db.Model(&OCENKA{}).AddForeignKey("disc_id", "discs (id)", "CASCADE", "CASCADE")

    //db, checkErr(err) := sql.Open("postgres", dbinfo)


    //checkErr(err)

    //var lastInsertId int
    //err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
    //err = db.QueryRow("select count(*) from ")
        
    //checkErr(err)
    //fmt.Println("", lastInsertId)

}

func dbfunc() *gorm.DB{
    db_, err_db := gorm.Open("postgres", fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",
      DB_PORT,DB_USER, DB_PASSWORD, DB_NAME))  
    if err_db != nil {
        log.Fatalf("error: %v\n", err_db)
        //return
    } 
    //defer db_.Close()
    return db_
}


func checkErr(err error) {
    if err != nil {
        log.Println(err)
    }
}