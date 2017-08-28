const loginusernamec = document.getElementsByClassName('login-username-container')[0];
const loginpasswordc = document.getElementsByClassName('login-password-container')[0];

const notepassword = document.getElementsByClassName('login-password-container')[0].getElementsByTagName('p')[1];
var username;

const authusername = function(){

    username = loginusernamec.getElementsByTagName('input')[0].value;

    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function(){
        if (this.readyState == 4 && this.status == 200){
            let response = JSON.parse(this.responseText);
            
            if (response.username == '') {                
                loginusernamec.getElementsByTagName('p')[0].innerHTML = 'username not yet registered';
            } else {
                loginusernamec.setAttribute('style', 'display: none;');

                loginpasswordc.setAttribute('style', 'display: block');
            }

            notepassword.innerHTML = 'input this password "'+username+'12"';
        }
    };

    xhttp.open('GET', '/loginusername?username='+username);
    xhttp.send();
};

const authpassword = function() {
    let password = loginpasswordc.getElementsByTagName('input')[0].value;
    
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status==200){
            let response = JSON.parse(this.responseText);

            if (response.status) {
                console.log('berhasil masuk');

                document.location.href = '/';
            } else {
                loginpasswordc.getElementsByTagName('p')[0].innerHTML = 'wrong password';
            }
        }
    }

    xhttp.open('POST', '/loginpassword?username='+username);
    xhttp.send(password);
}

loginusernamec.getElementsByTagName('button')[0].addEventListener('click', authusername);
loginpasswordc.getElementsByTagName('button')[0].addEventListener('click', authpassword);