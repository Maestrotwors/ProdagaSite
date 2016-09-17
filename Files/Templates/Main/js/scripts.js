<% define "vue.js" %>
<% template "localDB.js" %>

var Index = Vue.extend({
    template: <% template "MainPage.html" %>
})

var Forum = Vue.extend({
    template: <% template "ForumPage.html" %>
})

var Chat = Vue.extend({
    template: `<p>This is Chat!</p>`
})

var Settings = Vue.extend({
    template: <% template "SettingsPage.html" %>
})

var About = Vue.extend({
    template: <% template "AboutPage.html" %>
})

var Training = Vue.extend({
    template: <% template "TrainingPage.html" %>
})

var Login = Vue.extend({
    template: <% template "LoginPage.html" %>,
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
                //alert(response);
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
            //router.replace("/");
            //alert("logout");
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

var router = new VueRouter({
    history: true
})

router.map({
    '/': {
        component: Index,
    },
    '/forum': {
        component: Forum,
        auth: true
    },
    '/login': {
        component: Login
    },
    '/about': {
        component: About
    },
    '/training': {
        component: Training,
        auth: true,
        activate: function() {
            alert("Training");
        }
    },
    '/chat': {
        component: About
    },
    '/settings': {
        component: Settings,
        auth: true
    }
})

router.start(App, '#app')


router.beforeEach(function(transition) {
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