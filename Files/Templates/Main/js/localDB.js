<% define "localDB.js" %>

if (typeof(Storage) !== "undefined") {} else {
    alert("Сколько лет Вашему браузеру? Он очень старый и не может нормально работать с сайтом! Обновите браузер и пользуйтесь наздоровье");
}

window.indexedDB = window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB;

window.IDBTransaction = window.IDBTransaction || window.webkitIDBTransaction || window.msIDBTransaction;
window.IDBKeyRange = window.IDBKeyRange || window.webkitIDBKeyRange || window.msIDBKeyRange

if (!window.indexedDB) {
    window.alert("Ваш браузер не поддерживает IndexedDB. Установите современный браузер")
}

const customerData = [
    { Id: 1, Question: "Question 1" },
    { Id: 2, Question: "Question 2" },
    { Id: 3, Question: "Question 3" }
];


var db;
var request = window.indexedDB.open("UkrtestLocalDB", 3);

request.onerror = function(event) {
    console.log("DB error: ");
};

request.onsuccess = function(event) {
    db = request.result;
    console.log("DB success: " + db);
};

request.onupgradeneeded = function(event) {
    var db = event.target.result;
    var objectStore = db.createObjectStore("Questions", { keyPath: "Id" });
    for (var i in customerData) {
        objectStore.add(customerData[i]);
    }
}

function read() {
    var transaction = db.transaction(["Questions"]);
    var objectStore = transaction.objectStore("Questions");
    var request = objectStore.get(1);
    request.onerror = function(event) {
        alert("DB Error");
    };
    request.onsuccess = function(event) {
        if (request.result) {
            alert(request.result.Question);
        } else {
            alert("Вопрос 1 не найден");
        }
    };
}

function readAll() {
    var objectStore = db.transaction("Questions").objectStore("Questions");

    objectStore.openCursor().onsuccess = function(event) {
        var cursor = event.target.result;
        if (cursor) {
            alert("Name for id ");
            cursor.continue();
        } else {
            alert("No more entries!");
        }
    };
}

function add() {
    var request = db.transaction(["Questions"], "readwrite")
        .objectStore("Questions")
        .add({ Id: 4, Question: "Vopros 4 added" });

    request.onsuccess = function(event) {
        alert("Vopros 4 added to your database.");
    };

    request.onerror = function(event) {
        alert("Unable to add data\r\nKenny is aready exist in your database! ");
    }

}

function remove() {

    var request = db.transaction(["Questions"], "readwrite")
        .objectStore("Questions")
        .delete(4);
    request.onsuccess = function(event) {
        alert("Kenny's entry has been removed from your database.");
    };
}


<% end %>