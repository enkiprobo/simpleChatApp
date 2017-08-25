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
    
    let username = '';
    cookiekeyvalue.forEach(function(element, index) {
        if (element.key == cn) {
            username = element.value;
        }
    });

    return username;
}

/* GET USEFUL TAG */
const chatfriendc = document.getElementsByClassName('chat-friend')[0];
const chatheaderc = document.getElementsByClassName('chat-header')[0];
const chatcontentc = document.getElementsByClassName('chat-content')[0];
const chatactionc = document.getElementsByClassName('chat-action')[0];

/* MAIN FUNCTION */
const logout = function(){
    // deleting username cookie using past date as expire date
    document.cookie = 'username=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';

    document.location = '/';
};
const getchatfriend = function() {
    
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

    let selected = chatfriendc.selectedIndex;
    let friend = chatfriendc.options[selected].text;

    console.log(friend);

    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200){
            let response = JSON.parse(this.responseText)

            if (response.chat_detail_id != 0) {
                console.log(message);
            }
        }
    }

    xhttp.open('GET', '/insertmessage?friend='+friend+'&message='+message);
    xhttp.send();
}

/* SET EVENT LISTENER */
chatheaderc.getElementsByTagName('button')[0].addEventListener('click', logout);
chatactionc.getElementsByTagName('button')[0].addEventListener('click', sendmessage);

window.onload = function(){
    getchatfriend();
}