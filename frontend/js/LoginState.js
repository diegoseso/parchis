var LoginState = {
    preload: function(){
    },
    create: function(){

        //If there is no cookie with a username then ask for a name to login, otherwise use the same username in the cookie
        if (!isUserSet()){
            LoginScreen();
        }
        if (isUserSet()){
            this.state.start('PlayState');
        }
    }
};