package main
 
import(
     "sync"
     "github.com/gorilla/websocket"
     //"github.com/jinzhu/gorm"
     "time"
     "net/http"
)

var counterConnections NumberConnectionType 

func return_CounterConnections() uint64{  
    counterConnections.Lock()
    defer counterConnections.Unlock()
    return counterConnections.counters
}

type User struct {
    Id  int `gorm:"primary_key"`
    Login string  
    Password string       
    FIO string 
    RegisetDate time.Time `sql:"DEFAULT:current_timestamp"`
    //Role string
    //RoditelId int
}

type CategoryType struct {
    Id  int `gorm:"primary_key"`
    CategoryName string  //Auto, Nedvigimost
}

type Region struct {
    Id  int `gorm:"primary_key"`
    Region string
}

type AutoSale struct {
    Id  int `gorm:"primary_key"`
    RegUserId int
    UserText string
    UserDescription string
    DateTime time.Time `sql:"DEFAULT:current_timestamp"`
    RegionId int
    Price float64   
    Currency int
    Year int
    Motor int 
}

type Session struct{
    Id int `gorm:"primary_key"`
    SessionId string
    UserId int `sql:"not null"`
    Created string `sql:"DEFAULT:current_timestamp"`
}

/*
type Question struct {
    Id  int `gorm:"primary_key"`
    Question string  
    UserId int
    CreatedTime time.Time `sql:"DEFAULT:current_timestamp"`
}

type Variant struct {
    Id  int `gorm:"primary_key"`
    QuestionId int `sql:"not null"` 
    Variant string
    IsTrue bool
}

type QuestionComment struct {
    Id  int `gorm:"primary_key"`
    QuestionId int
    UserId int `sql:"not null"`
    CreatedTime time.Time `sql:"DEFAULT:current_timestamp"`
    DeletedTime time.Time  
    Comment string  
}

type DataBase struct {
    Id  int `gorm:"primary_key"`
    Name string  
    Description string
    UserId int `sql:"not null"`
    CreatedTime time.Time `sql:"DEFAULT:current_timestamp"`
    DeletedTime time.Time  
}



type Forum struct{
    Id int `gorm:"primary_key"`
    Forum string
    UserId int `sql:"not null"`
    Path int //Parent
    Created string `sql:"DEFAULT:current_timestamp"`
}

type ForumDirectories struct{
    Id int `gorm:"primary_key"`
    Directory string
    ParentId int `sql:"not null"`
    Created string `sql:"DEFAULT:current_timestamp"`
}

type DBDirectories struct{
    Id int `gorm:"primary_key"`
    Directory string
    ParentId int `sql:"not null"`
    Created string `sql:"DEFAULT:current_timestamp"`
}*/

//------------------------

type ClientConnType struct {
    websocket *websocket.Conn
    Socket_Id uint64
    USER User
}

type NumberConnectionType struct {
    counters uint64 
    sync.Mutex
}

type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}

type SecureAuthorize struct {
    http.Handler
} 








//----------------- ������� ����� ���������


//getSession Functions
/*func getSession(r *http.Request) Session{
    cookie, err := r.Cookie("SessionId")
    if err != nil {
      session:=Session{}
      return session
    }else{
        //log.Println("���� �������", cookie.Value)
        session:=SessionList[cookie.Value]
        return session
    }
}

type DISC struct {
    ID uint `gorm:"primary_key"`
    Name string
}
 

type OCENKA struct {
    ID uint `gorm:"primary_key"`
    USERID uint
    DISCID uint
    Value int
}

type USER struct {
    ID uint `gorm:"primary_key"`
    Name string
    Login string
    Password string
    Folders []FOLDER
}

type FOLDER struct{
    ID uint `gorm:"primary_key"`
    UserId uint
    HashId uint
    Name string
}

*/


