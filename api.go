package main
 
import(
    "fmt"
    "log"
    "net/http"
    "os"
    "io"
    "github.com/nfnt/resize"
    "image/jpeg"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "crypto/sha1"
)

//Upload Загрузка файлов
func NewSale(w http.ResponseWriter, r *http.Request) {
    if r.Method !="POST"{return}
    Session:=getSession(r)

    /*
    UserText:=r.FormValue("UserText")
    UserPrice:=r.FormValue("UserPrice")
    UserYear:=r.FormValue("UserYear")
    UserMotor:=r.FormValue("UserMotor")
    UserOblast:=r.FormValue("UserOblast")*/
    hash_string := fmt.Sprintf("hash %s", Session.SessionId)
    h := sha1.New()
    h.Write([]byte(hash_string))
    bs := h.Sum(nil)      
    Hash:=fmt.Sprintf("%x", bs)

    Auto := AutoSale{UserText : r.FormValue("UserText")}
    db, err := gorm.Open("postgres", fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",
      DB_PORT,DB_USER, DB_PASSWORD, DB_NAME))  
    if err != nil {
        log.Fatalf("error: %v\n", err)
        //return
    } 
    defer db.Close()
    db.Create(&Auto)
    log.Println(Auto.Id)

    err_ := r.ParseMultipartForm(10000000)
	if err_ != nil {
			return
	}

	m := r.MultipartForm

	files := m.File["files"]
    log.Println(files) 
    for i, _ := range files {
        file, err := files[i].Open()
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        f, err := os.OpenFile(fmt.Sprintf("Files\\Upload\\Auto\\original\\id%d_%d_U%d_%s.jpg",Auto.Id,i+1,Session.UserId,Hash), os.O_WRONLY|os.O_CREATE, 0666) //handler.Filename
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)

        fileNew, err := os.Open(fmt.Sprintf("Files\\Upload\\Auto\\original\\id%d_%d_U%d_%s.jpg",Auto.Id,i+1,Session.UserId,Hash))
        if err != nil {
            log.Fatal(err)
        }

        img, err := jpeg.Decode(fileNew)
        if err != nil {
            log.Fatal(err)
        }
        fileNew.Close()

        m := resize.Resize(200, 140, img, resize.Lanczos3)

        out, err := os.Create(fmt.Sprintf("Files\\Upload\\Auto\\min\\id%d_%d_U%d_%s_min.jpg",Auto.Id,i+1,Session.UserId,Hash))
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()

        jpeg.Encode(out, m, nil)
    }


} 