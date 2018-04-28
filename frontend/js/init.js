var game = new Phaser.Game(600, 600, Phaser.AUTO, 'game');

game.state.add('BootState', BootState);
game.state.add('LoginState', LoginState);
game.state.add('PlayState', PlayState);
game.state.start('BootState');