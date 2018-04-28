var isUserSet;
var sconn;
var username;
var server = {address: "ws://localhost:8081/ws"};

function isUserSet(){
    console.log('isUserSet function');
    if (Cookies.get('username') === undefined){
       return false;
    } else {
       currentUsername = Cookies.get('username')
       return true;
    }
}

function LoginScreen(){
    console.log('LoginScreen function');
    $('#login-form').show();
    $('#login-play').click(function(){
        username = $('#username').val();
        $('#login-form').hide();
        Login(username);
    });
}

function Login(username){
    if (!sconn) {
        return false;
    }
    if (sconn.readyState != 1){

    }
    var newLogin = {message:'login', username : username}
    sconn.send(JSON.stringify(newLogin));
    currentUsername = Cookies.set('username', username);
    // Login process to the server
    $('#game').show();
}

function ServerConnect(){
    // Here we connect to the server
    return new Promise(function(resolve, reject) {
        sconn = new WebSocket(server.address);
        sconn.onopen = function () {
            resolve(sconn)
        };
        sconn.onclose = function (evt) {
            reject(err);
        };
    });
}

function ServerUnavailable(){
    console.log('The server is unavailable');
}


function playIntroAudio(){
    welcomeLoop = self.game.add.audio('welcome');
    welcomeLoop.play();
}

var LoginState = {
    preload: function(){
        console.log('loginState preload');
        playIntroAudio();
    },
    create: function(){
        console.log('loginState create');
        ServerConnect().then(function() {
            //If there is no cookie with a username then ask for a name to login, otherwise use the same username in the cookie
            if (!isUserSet()){
                LoginScreen();
            }
            if (isUserSet()){
                Login(currentUsername);
            }
            game.state.start('PlayState');
        }).catch(function(err) {
            console.error('Can not connect to the server');
            console.error(err)
        });
        console.log(sconn);
    }
}