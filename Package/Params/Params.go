package Params
 
import(
       types "../Types/"
)

var UsersList = make(map[int ] types.User)   // Список пользователей
var ActiveClients = make(map[uint64 ] types.ClientConnType) //Активный соккет клиента
var CounterConnections types.NumberConnectionType//Количество сокет подключений
//var UsersOnline = make(map[uint64 ] User)//  тестировать 
var SessionList= make(map[string ] types.Session) //Список сессий
//var db *gorm.DB





