<% define "vue.js" %>
<% template "localDB.js" %>

window.MainApp = { router: null, MyParam1: "MyParam111111", MyParam2: "MyParam2222" }

var Index = Vue.extend({
    template: "#MainPage"
})

var Forum = Vue.extend({
    template: "#ForumPage"
})

var UserInfo = Vue.extend({
    template: "#UserInfo"
})

var Sale = Vue.extend({
    template: "#SalePage",
    data: function() {
        return {
            image: ''
        }
    },
    methods: {
        SaleSubmit: function() {},
        show: function() {
            alert("showww");
        }
    }
})


var About = Vue.extend({
    template: "#AboutPage"
})

$(document).ready(function() {

    $('input[type="file"]').change(function() {
        var file = this.files; //Files[0] = 1st file
        var numFiles = this.files.length;
        alert(numFiles);
        if (file[0]) {
            var reader = new FileReader();
            reader.readAsDataURL(file[0], 'UTF-8');
            reader.onload = function(event) {
                var result = event.target.result;
                $('.pic_1').attr('src', event.target.result);

            };
        }
    })

    $("form#FormUpload").submit(function(event) {
        event.stopPropagation();
        event.preventDefault();
        var formData = new FormData($(this)[0]);
        $.ajax({
            url: '/api/upload',
            type: 'POST',
            data: formData,
            async: false,
            cache: false,
            contentType: false,
            processData: false,
            success: function(returndata) {
                alert(returndata);
            }
        });

        return false;
    });
});


var Login = Vue.extend({
    template: "#LoginPage",
    data: function() {
        return {
            param1: "parammmmmm1"
        }
    },
    methods: {
        login: function() {
            $.post("/api/authorize", {
                "login": $("#login").val(),
                "password": $("#password").val()
            }).success(function(response) {
                alert(response);
                if (response > "") {
                    localStorage.Auth = "1";
                    arr = JSON.parse(JSON.stringify(response));
                    document.cookie = "SessionId=" + arr[0];
                    localStorage.UserName = arr[1];
                    UserName = arr[1];
                    router.go("/about");
                    //window.location.replace("/");
                } else {
                    alert("Не подходит");
                    localStorage.Auth = "0";
                    localStorage.UserName = "";
                    UserName = "";
                    document.cookie = "SessionId=0";
                }
            });
        },
        UploadFile: function(e) {
            alert('upl');
            //var files = e.target.files || e.dataTransfer.files;
        },
        onFileUpload: function(e) {
            alert('upl');
            //var files = e.target.files || e.dataTransfer.files;
        }
    },
    events: {
        //'App':function
    }
    //props: ['UserName']
})


//var Connection;
var App = Vue.extend({
    el: '#app',
    data: function() {
        return {
            UserName: "Roman", //localStorage.UserName,
            Auth: localStorage.Auth,
            Online: window.navigator.online,
            Connected: false,
            Param: 7,
            WSConnection: null,
            ConnectionValueChange: function(val) { this.Connected = val },
            messages: []
        }
    },
    /*
        components: {
            'Login': Login
        },*/
    events: {
        'Login': function(msg) {
            this.messages.push(msg)
        }
    },
    methods: {
        ServerConnect: function() {
            console.log("sc");
            WSConnection = WebSocketConnect();
        },
        connected: function(val) {
            this.Connected = val;
            //var Connection = WebSocketConnect();
        },
        isOnline: function() {
            this.Online = true;
            this.ServerConnect();
            //var Connection = WebSocketConnect();
        },
        isOffline: function() {
            this.Online = false;
            WSConnection.Close();
        },
        logout: function() {
            alert(11111111111111111);
            UserName = "";
            localStorage.UserName = "";
            localStorage.Auth = "0";
            document.cookie = "SessionId=0";
            //window.location.replace("/");
        }
    },
    created: function() {
        window.addEventListener("online", this.isOnline, false);
        window.addEventListener("offline", this.isOffline, false);
        if (this.Online) { this.isOnline() }
        this.ServerConnect();
    },
})



window.MainApp.router = new VueRouter({
    history: true
})

window.MainApp.router.map({
    '/': {
        component: Index
    },
    '/user/id:id': {
        component: UserInfo
    },
    '/forum': {
        component: Forum
    },
    '/login': {
        component: Login
    },
    '/about': {
        component: About
    },
    '/sale': {
        component: Sale,
        auth: true
    }
})

window.MainApp.router.start(App, '#app')


window.MainApp.router.beforeEach(function(transition) {
        if (transition.to.auth && localStorage.Auth != "1") {
            transition.redirect('/login');
        } else {
            transition.next();
        }
    })
    //var App = Vue.extend({})



//                                                                       Разобрать



login1222 = function() {
    alert(777);
    $.post("/api/authorize", {
        "login": $("#login").val(),
        "password": $("#password").val()
    }).success(function(data) {
        if (data > "") {
            document.cookie = "SessionId=" + data;
            //window.location.replace("/");
        } else {
            alert("Не подходит");
        }
    });
}

logout11111 = function() {
    $.post("/api/logout", {

    }).success(function(data) {
        alert("logout");
    });
}

//$('#form input').keydown(function(e) {
//    if (e.keyCode == 13) {
//        login();
//    }
//});



<% end %>