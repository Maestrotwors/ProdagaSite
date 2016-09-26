package main
 
import(
    	"net/http"
    	"fmt"
        "time"
        "crypto/sha1"
        "github.com/jeffail/gabs" 
        //params "./Package/Params/"
        //types "./Package/Types/"
        //"path/filepath"
        //"github.com/jinzhu/gorm"
        _ "github.com/lib/pq"
        "log"
        //"github.com/gorilla/mux"
        //php "github.com/deuill/go-php"
)

//Authorize Авторизация
func Authorize(w http.ResponseWriter, r *http.Request) {
        var Login= r.PostFormValue("login")
        var Password= r.PostFormValue("password")
        for _, UserInList:= range UsersList {
            if (UserInList.Login== Login && UserInList.Password== Password) {
                Make_Session(w,r,&UserInList)                             
            }
        }
} 

//Ping Проверка наличия сервера
func Ping(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "1") 
} 

//Logout Выход
func Logout(w http.ResponseWriter, r *http.Request) {
    var sess Session
    sess=getSession(r)    
    for _, session:= range SessionList {
        if session.UserId==sess.UserId {
            delete(SessionList,session.SessionId) 
        }
    } 
    fmt.Fprintln(w, "<script>var date = new Date(0);document.cookie = 'SessionId=; path=/; expires=' + date.toUTCString();window.location.replace('/login');</script>")
} 

func Upload(w http.ResponseWriter, r *http.Request) {
    NewSale(w,r)
}

// FileSave
func FileSave(){
    /*
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    f, err := os.OpenFile("Files\\Upload\\Auto\\original\\id1_1_U1.jpg", os.O_WRONLY|os.O_CREATE, 0666) //handler.Filename
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)

    fileNew, err := os.Open("Files\\Upload\\Auto\\original\\id1_1_U1.jpg")
    if err != nil {
        log.Fatal(err)
    }

    img, err := jpeg.Decode(fileNew)
    if err != nil {
        log.Fatal(err)
    }
    fileNew.Close()

    m := resize.Resize(200, 140, img, resize.Lanczos3)

    out, err := os.Create("Files\\Upload\\Auto\\min\\id1_1_U1_min.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    jpeg.Encode(out, m, nil)*/
}

// GormInsert
func GormInsert(Data interface{}){

}

// GormUpdate
func GormUpdate(Data interface{}){

}

// GormDelete
func GormDelete(Data interface{}){

}

//Make_Session Создать сессию
func Make_Session(w http.ResponseWriter, r *http.Request,USER *User){
  w.Header().Set("Content-type", "application/json")  
  t := time.Now()
  hash_string := fmt.Sprintf("hash %s %s", t,USER.Login)
  h := sha1.New()
  h.Write([]byte(hash_string))
  bs := h.Sum(nil)      
  SessionId:=fmt.Sprintf("%x", bs)

  //var time string
  //time= fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d",t.Year(),t.Month(),t.Day(),t.Hour(), t.Minute() ,t.Second())
  //Create := string(time)
  log.Println("Сессия создана. юзер", USER.Id) 
  sess := Session{SessionId:SessionId,UserId:USER.Id, Created:t.Format("2006-01-02 15:04:05.0000")}
  //sess := Session{0,SessionId, USER.Id,Create}
  SessionList[SessionId] = sess
  jsonObj := gabs.New()
  jsonObj.Array()
  jsonObj.ArrayAppend(SessionId)
  jsonObj.ArrayAppend(USER.FIO)
  //jsonObj.ArrayAppend("Par1")
  fmt.Fprintln(w, jsonObj.String())
  //fmt.Fprintln(w, "[\"",SessionId,"\",\"",USER.FIO,"\"]") //Отправляем пользователю сессию и его имя 
  

  db:=*dbfunc()
  db.Save(&sess)
  db.Close()
  
  //db.NewRecord(sessX);db.Create(&sessX);db.NewRecord(sessX);db.Save(&sessX);db.Close()

  //query:=fmt.Sprintf("insert into session (Session,UserId,Fio,BeginDate) values('%s',%d,'%s','%s')  ",SessionId,User.Id,User.FIO,time)
  //if new_query_exec(query){}else{} 
  //jsonObj := gabs.New()
  //jsonObj.Array()
  //jsonObj.ArrayAppend(SessionId)
  //jsonObj.ArrayAppend(USER.FIO)
  //fmt.Fprintln(w, jsonObj.String())
}
