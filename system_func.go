package main

import(
	"net/http"
    "log"
)

func ParseBool(str string) (value bool) {
    switch str {
    	case "1","true", "TRUE", "True":
    			return true
    	case "0","false", "FALSE", "False":
    			return false
    }
    return false
}

func getSession(r *http.Request) Session{
    cookie, err := r.Cookie("SessionId")
    if err != nil {
      session:=Session{}
      return session
    }else{
        
        session:=SessionList[cookie.Value]
        log.Println("Session считана", session)
        return session
    }
}


 

//-----------------------SECURE

func (wr SecureAuthorize) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var sess Session
    sess=getSession(r)
    if (sess.SessionId>""){                        
      wr.Handler.ServeHTTP(w, r) 
      return       
    }else{
      MainHandler(w, r,"About")  
    }  
    //fmt.Fprintln(w, "Вы не зарегистрированы! Сюда лезть нельзя! Доступ закрыт!")
}

/*
func (wr SecureGold) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var sess Session
    sess=getSession(r)
    if (sess.UserRole >= "1" &&   sess.Authorized== true)  {                        
      wr.Handler.ServeHTTP(w, r) 
      return       
    }  
    fmt.Fprintln(w, "Вы не Gold ! Доступ закрыт!")
    return 
}

func (wr SecureAdmin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var sess Session
    sess=getSession(r)
    //if (sess.UserRole >= "5" &&   sess.Authorized == true)  { 
    if (sess.UserId==1)  {                        
      wr.Handler.ServeHTTP(w, r) 
      return       
    }  
    fmt.Fprintln(w, "Вы не администратор! Доступ закрыт!")
    return 
} */

