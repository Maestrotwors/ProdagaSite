package main
 
import(
    	"net/http"
    	"fmt"
        "bytes"
        "html/template"
	    "github.com/tdewolff/minify"
	    "github.com/tdewolff/minify/css"
	    "github.com/tdewolff/minify/html"
	    "github.com/tdewolff/minify/js"
        "github.com/gorilla/mux"
        "io/ioutil"
        //"log"
)


//MainHandler Хандлер для загрузки шаблонов
func MainHandler(w http.ResponseWriter, r *http.Request,Path string) {
    var sess Session
    sess=getSession(r)  
    if (sess.SessionId>""){
    //    MainHandler(w, r,"Login")
    //    return
    }
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    templ  := template.New("templ")
    templ .Delims("<%", "%>")
    var templates = template.Must(templ.ParseGlob("Files/Templates/"+Path+"/*/*"))  
    var doc bytes.Buffer 
    var docWrite bytes.Buffer
    m := minify.New()
    m.AddFunc("text/css", css.Minify)
    m.AddFunc("text/html", html.Minify)
    m.AddFunc("text/javascript", js.Minify)
    err := templates.ExecuteTemplate(&doc, "CorePage", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }  
    if err := m.Minify("text/html", &docWrite, &doc); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //FilesCache.Set("IndexPage", &docWrite, cache.DefaultExpiration)
    fmt.Fprintln(w, &docWrite)
}

//ApiHandler Хандлер для API
func ApiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    
    vars:=mux.Vars(r)
    method:=vars["method"]
    switch method {
        case "authorize":
            Authorize(w,r)
        case "ping":
            Ping(w,r)
        case "upload":
            Upload(w,r)
            fmt.Fprintln(w, "Upload") 
        case "Logout":
            Logout(w,r)

        default:  
            fmt.Fprintln(w, "Метод не распознан") 
    }
} 

//StartHandler Хандлер роутинга основных маршрутов 
func StartHandler(w http.ResponseWriter, r *http.Request,Handler string) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    switch Handler {      
        case "Logout":
            Logout(w,r)
        case "About":
            GoAuth(w,r,"About")
        case "Main":
            GoAuth(w,r,"Main")
        case "Index":
            MainHandler(w, r,"Index")

        default: 
            fmt.Fprintln(w, "Ошибка. Хандлер не найден") 
    }
}

//GoAuth Авторизация
func GoAuth(w http.ResponseWriter, r *http.Request, Path string) {
    var sess Session
    sess=getSession(r)  
    if (sess.SessionId>""){
        MainHandler(w, r,Path)
    }else{
        MainHandler(w, r,Path)
        //http.Redirect(w, r, "/login", http.StatusSeeOther) 
        //MainHandler(w, r,"AboutNA")
    }
} 


var manifest=""
//ManifestHandler Манифест
func ManifestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/cache-manifest .cache")
	manifestb, err := ioutil.ReadFile("Manifest/manifest.mf")
    if err != nil {
        return
    }
    manifest= string(manifestb) 
	fmt.Fprintln(w, manifest)
}

//serveSingle Загрузка определённого файла
func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}


