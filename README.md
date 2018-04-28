# parchis

## Idea behind the project

Provide the simplest but effective server to play Parchis on over 100.000 Boards at a time. Why 100.000 ?
Because I want to demonstrate that with a go Server 400.000 People can be Entertained.

## Way of this game server to work

The simplest premises of this Gameserver are:
  * A game server provides a platform to work according the game rules, outputting the results of any action so events
  around them can be programmed and a business can be delivered.
  * A Server has room´s, each room has it´s own characteristics, Inside a Room can be any number of boards,
  a board can hold a Match at a time, the Match has players, an status and a history.

# How to run this project

docker run -dit -p 8001:80 -v /home/diegoseso/go/src/github.com/diegoseso/parchis/frontend:/usr/share/nginx/html/html nginx

## Other dependencies:
- jQuery
- Bootstrap

## Instalation
This project need bower dependences. Go to folder project and run the following command:

```
bower install
```