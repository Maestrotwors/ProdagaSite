package main

import (
    "github.com/gorilla/websocket"
    "github.com/jeffail/gabs" 
    "net/http"
    //"sync"
    "fmt"
    "log"
    //types "./Package/Types/"
)



func SendMessageToClient(sockId uint64, msg string) {
  	ActiveClients[sockId].websocket.WriteMessage(websocket.TextMessage, []byte(msg))  
}

func SendMessageToAllClients(msg  string) {
  for _, Client := range ActiveClients {
    ActiveClients[Client.Socket_Id].websocket.WriteMessage(websocket.TextMessage, []byte(msg))
  }
}

func (s *NumberConnectionType) GetNew() uint64{
        s.Lock()
        s.counters+=1
        defer s.Unlock()
        return s.counters
}


func ClientSockServer(w http.ResponseWriter, r *http.Request) {
    //vars := mux.Vars(r)  //vars["key"]
    var SocId uint64=CounterConnections.GetNew() 
    conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    if _, ok := err.(websocket.HandshakeError); ok {
      http.Error(w, "Not a websocket handshake", 400)
      return
    }

    var sess Session
    sess=getSession(r)
    Id := int(sess.UserId) 
    FIO := UsersList[int(sess.UserId)].FIO
    
    log.Println("Клиент подкл.Сокет-",SocId,"   ", Id,"   ",FIO) 
    defer conn.Close();

    jsonObj:= gabs.New()
    jsonObj.Set("DDOS_Start", "Command")
    jsonObj.Set("http://localhost:5001/", "Data", "Site")
    jsonObj.Set(25, "Data", "Routine_numb")
    SendMessageToAllClients(jsonObj.String()) 

    for { 
        _, msg, err := conn.ReadMessage()
        msg1:=fmt.Sprintf("%s",msg)
        mesg:=[]byte(msg1)
        //msg_ := []byte(bytes.Index(byteArray, []byte{0}))

        if err != nil {  
                log.Println("Клиент откл .Сокет-",SocId,"   ", Id,"   ",FIO)  
                delete(ActiveClients,SocId) 

                /*
                jsonObj_:= gabs.New()
                jsonObj_.Array("Clients")
                for _, Client := range ActiveClients {
                  jsonObj_.ArrayAppend(Client, "Clients")
                }
                jsonObj_.Set("Clients_List", "Command")

                jsonObj:= gabs.New()
                jsonObj.Set("Clients_List", "Command")
                */
                //SendMessageToAllAdmins(jsonObj_.String())    
          return
        }

        var Command string
        var ok bool
        jsonParsed, err_ := gabs.ParseJSON(mesg)
          if err_ != nil {
            log.Println("Данные от клиента пришли не в json формате")
            return
          }
        Command, ok= jsonParsed.Path("Command").Data().(string)
        if ok == false {
          log.Println("Данные от клиента пришли без команды")
          return
        } 

          switch Command {

          default:
              log.Println("Функция,которую отослал клиент, не распознана- ")
        }

    }      
}


//--------Send Message

/*
func SendMessageToClient(sockId uint64, msg string) {
  	ActiveClients[sockId].websocket.WriteMessage(websocket.TextMessage, []byte(msg))  
}

func SendMessageToAllClients(msg  string) {
  for _, Client := range ActiveClients {
    ActiveClients[Client.Socket_Id].websocket.WriteMessage(websocket.TextMessage, []byte(msg))
  }
}

func SendMessageToAdmin(sockId uint64, msg string) {
   // ActiveAdmins[sockId].websocket.WriteMessage(websocket.TextMessage, []byte(msg))  
}

func SendMessageToAllAdmins(msg  string) {
  for _, Admin := range ActiveAdmins {
    ActiveAdmins[Admin.Socket_Id].websocket.WriteMessage(websocket.TextMessage, []byte(msg))
  }
}*/