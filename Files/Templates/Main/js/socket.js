<% define "socket.js"  %>
WebSocketConnect = function() {
    var Sock = new WebSocket("ws://localhost/websocket");
    try {
        Sock.onopen = function(m) {
            console.log("Соединение открыто...");
            //console.log(Params);
            //Params.ConnectionValueChange(true)
            //App.data.al();
            //App.set(data, 'Connected', true)
            //App.$set('Connected', true)
            //Connection.ChatRoomConnect("5555");
            /*var Send= {
              "Command":"DDOS_Start",
              "Data":{"Site":"http://localhost:5001/","Routine_numb":27}
            };
            sock.send(JSON.stringify(Send));*/
        }
        Sock.onmessage = function(m) {
            //console.log(m);
            data_ = m['data'];
            data = JSON.parse(data_);
            switch (data["Command"]) {
                case 'Clients_List':
                    //var data_=data["Clients_List"]
                    //console.log(data["Clients"]);
                    MyAppl.clients = data["Clients"];
                    break
                case 'Client_Data_Changed':
                    alert("Client_Data_Changed");
                    break
                    //case 'Ddos_Connections':
                    // var x = _.find(MyAppl.clients, { "Socket_Id": data["Soc_Id"] });
                    // x["DdosConnections"] = data["Count"];
                    //MyAppl.clients=data["Clients"]; 
                    //break

                default:
                    alertify.error("Функция не распознана");
            }

        };
        Sock.onerror = function(m) {
            alertify.error('Ошибка подключения');
        };
        Sock.onclose = function(m) {
            console.log("Соединение разорвано");
            var intervalID = setInterval(function() {
                $.get("/api/ping", function() {
                        clearInterval(intervalID);
                        var Connection = WebSocketConnect();
                    })
                    .error(function() {
                        console.log("Не удалось подключиться к серверу");
                        App.Online = null;
                        console.log(App);
                    })
            }, 3000);
        };

    } catch (exception) {
        console.log("Не удалось подключиться к серверу");
    }
    return {
        Close: function() {
            Sock.close();
        },
        Password_Change: function(AltPassword, NewPassword) {
            var Send = {
                "Command": "Password_Change",
                "Data": { "AltPassword": AltPassword, "NewPassword": NewPassword }
            };
            Sock.send(JSON.stringify(Send));
        },
        ChatRoomConnect: function(RoomId) {
            var SendData = {
                "Command": "ChatRoomConnect",
                "Data": { "RoomId": RoomId }
            };
            Sock.send(JSON.stringify(SendData));
        }
    }
}

<% end %>