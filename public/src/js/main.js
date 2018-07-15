import {GameBoard} from './GameBoard.js';
import {SocketConnection} from './SocketConnection';
import {IO} from './IO';

var socketConnection = new SocketConnection();
var io = new IO(socketConnection);
var gameBoard = new GameBoard();
gameBoard.renderBoard();