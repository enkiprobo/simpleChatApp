/* HELPER FUNCTION */
// get cookies and make it as array of object cookie key value pair
function getcookiekeyvalue(){
    let cookies = document.cookie;
    let arrcook = cookies.split(';');
    let cookiekeyvalue = arrcook.map(function(cookie){
        let tempcookie = cookie.split('=');
        return {
            key: tempcookie[0],
            value: tempcookie[1]
        };       
    });

    return cookiekeyvalue;
}
// get cookie with name as cn
function getuserfromcookie(cn) {
    let cookiekeyvalue = getcookiekeyvalue();
    
    console.log(cookiekeyvalue);

    let username = '';
    cookiekeyvalue.forEach(function(element, index) {
        if (element.key == cn) {
            username = element.value;
        }
    });

    console.log(username);
    return username;
}
// get selected friend
function getselectedfriend(){
    let selected = chatfriendc.selectedIndex;
    let friend = chatfriendc.options[selected].text;

    return friend;
}
// get only hour and minute from timestamp
function gettimeonlyhourandminute(datetime){
    let dateandtime = datetime.split('T');

    return dateandtime[1].substring(0, dateandtime[1].length - 4);
}
// get time now only hour and minute from date.now
function gettimeonlyhourandminutenow(){
    let now = new Date();

    return now.getHours() + ':' + now.getMinutes();
}
// insert new message html
function insertmessageHTML(messagecontentstr, message_author, messagedatestr){
    let friend = getselectedfriend();

    let message = document.createElement('div');
    let messagecontent = document.createElement('p');
    let messagedate = document.createElement('p');
    
    messagecontent.setAttribute('class', 'message');
    messagecontent.innerHTML = messagecontentstr;
    messagedate.setAttribute('class', 'date');
    messagedate.innerHTML = messagedatestr;

    if (message_author == friend) {
        message.setAttribute('class', 'other-chat');

        message.appendChild(messagecontent);
        message.appendChild(messagedate);
    } else {
        message.setAttribute('class', 'own-chat');
    
        message.appendChild(messagedate);
        message.appendChild(messagecontent);
    }

    return message;
}

/* GET USEFUL TAG */
const chatfriendc = document.getElementsByClassName('chat-friend')[0];
const chatheaderc = document.getElementsByClassName('chat-header')[0];
const chatcontentc = document.getElementsByClassName('chat-content')[0];
const chatactionc = document.getElementsByClassName('chat-action')[0];
const userloginnow = getuserfromcookie('username');

console.log("userloginnow");
/* WEBSOCKET */
var ws = new WebSocket('ws://10.20.33.100:8080/livechat');
/*
    {
        message_owner:,
        message:,
        message_date:
    }
*/
ws.onmessage = function(event) {    
    let md = JSON.parse(event.data);
    let friend = getselectedfriend();

    console.log(md);

    if (md.message_owner == friend) {
        let messagedate = md.message_date;
        let messageHTML = insertmessageHTML(md.message, md.message_owner, messagedate);

        chatcontentc.appendChild(messageHTML);
    }
};

/* MAIN FUNCTION */
const logout = function(){
    // deleting username cookie using past date as expire date
    document.cookie = 'username=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';

    ws.close();
    document.location = '/';
};
const getchatfriend = function(resolve, reject) {
    
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            let response = JSON.parse(this.responseText);

            response.forEach(function(friend){
                let chatfriend = document.createElement('option');
                chatfriend.setAttribute('value', friend.username);
                chatfriend.innerHTML = friend.username;

                chatfriendc.appendChild(chatfriend);
            });

            resolve(); // for promise
        }
    };

    xhttp.open('GET', '/getchatfriends');
    xhttp.send();
};
const sendmessage = function(){
    
    chatactionc.getElementsByTagName('p')[0].innerHTML = '';

    let message = chatactionc.getElementsByTagName('input')[0].value;
    if (message == ''){
        chatactionc.getElementsByTagName('p')[0].innerHTML = 'message can\'t be empty'; 
        return
    }

    let friend = getselectedfriend();

    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200){
            let response = JSON.parse(this.responseText)

            if (response.chat_detail_id != 0) {
                chatactionc.getElementsByTagName('input')[0].value = '';

                let hourminute = gettimeonlyhourandminutenow();
                let messageHTML = insertmessageHTML(message, userloginnow, hourminute);

                chatcontentc.appendChild(messageHTML);

                let messagesocket = {
                    message_owner: userloginnow,
                    message: message,
                    message_date: hourminute 
                };
                ws.send(JSON.stringify(messagesocket));
            }
        }
    };

    xhttp.open('GET', '/insertmessage?friend='+friend+'&message='+message);
    xhttp.send();
};
const getsetchatdetail = function() {

    let friend = getselectedfriend();

    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200){
            let response = JSON.parse(this.responseText);

            response.forEach(function(m) {
                let message = insertmessageHTML(m.message, m.message_author, gettimeonlyhourandminute(m.create_date));
                chatcontentc.appendChild(message);
            });
        }
    };

    xhttp.open('GET', '/getchatdetail?friend='+friend);
    xhttp.send();
};
const selectchatfriend = function() {

    chatcontentc.innerHTML = '';

    getsetchatdetail();
};

/* SET EVENT LISTENER */
chatheaderc.getElementsByTagName('button')[0].addEventListener('click', logout);
chatactionc.getElementsByTagName('button')[0].addEventListener('click', sendmessage);
chatfriendc.addEventListener('change', selectchatfriend);

/* MAIN */
window.onload = function(){
    let order = new Promise(getchatfriend)
        .then(() => selectchatfriend());
}