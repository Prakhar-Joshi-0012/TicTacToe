
# Tic Tac Toe Game
Simple Tic Tac Toe game created in Go using sdl2 and openGL. 
* You are player O and Computer is player X. 
* You play first
  
![Alt StartScreen](/impls/start.png)
![Alt PlayScreen](/impls/play.png)
![Alt FinishScreen](/impls/finish.png)

## About Game
Tic Tac Toe is a puzzle game for two players, called "X" and "O", who take turns marking the spaces in a 3×3 grid. The player who succeeded in placing three respective marks in a horizontal, vertical, or diagonal row wins the game. There is only single player option present in the game, in which you play against computer.
## How To play
To play the game press space and click the tile in which you want to place your move.
## Prerequites
Go 1.21.2 or later required to build
## Usage
Simplest way to install would be using `git clone`.
```sh
$ git clone git@github.com:Prakhar-Joshi-0012/TicTacToe.git
```
Then run the following command to resolve sdl2 module dependencies
```sh
$ go mod tidy
```
use `build` command to create executable.
```sh
$ go build
$ ./TicTacToe
```


