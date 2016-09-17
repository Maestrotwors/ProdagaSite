package main
 
import(
    //"log"
    //"github.com/jinzhu/gorm"
    //_ "github.com/lib/pq"
    //"fmt"
    //"time"
)

func Cache() {

    db:=*dbfunc()
    defer db.Close()

    allUsers := []User{}
    db.Find(&allUsers)
    for _, data := range allUsers {
        UserX := data
        UsersList[data.Id] = UserX 
    }  
  
    allSessions := []Session{}
    db.Find(&allSessions)
    for _, data := range allSessions {
        SessionX := data
        SessionList[data.SessionId] = SessionX 
    } 

}

