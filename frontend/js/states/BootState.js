var BootState = {
    preload: function(){
        console.log('Enter Boot state preload');
        game.load.image('board', 'img/board.png');
        game.load.image('blue', 'img/blue.png');
        game.load.image('green', 'img/green.png');
        game.load.image('red', 'img/red.png');
        game.load.image('yellow', 'img/yellow.png');
        game.load.spritesheet('dice-init', 'img/dice-init.png', 157, 158, 4);
        game.load.image('dice1', 'img/dice1.png');
        game.load.image('dice2', 'img/dice2.png');
        game.load.image('dice3', 'img/dice3.png');
        game.load.image('dice4', 'img/dice4.png');
        game.load.image('dice5', 'img/dice5.png');
        game.load.image('dice6', 'img/dice6.png');
        this.load.audio('welcome', ['audio/welcome.ogg', 'audio/welcome.mp3']);

    },
    create: function(){
        this.state.start('LoginState');
    }
};