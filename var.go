package main
 
import(
    //"github.com/jinzhu/gorm"
    //_ "github.com/lib/pq"
    //"log"
)

var UsersList = make(map[int ] User)   // Список пользователей
var CounterConnections = NumberConnectionType{counters: 0} //Количество сокет подключений
var ActiveClients = make(map[uint64 ] ClientConnType) //Активный соккет клиента
//var UsersOnline = make(map[uint64 ] User)//  тестировать 
var SessionList= make(map[string ] Session) //Список сессий
//var db *gorm.DB





