package coinOnTheTable

//You have a rectangular board consisting of n rows, numbered from 1 to n, and columns, numbered from 1 to m.
//The top left is (1, 1) and the bottom right is (n, m). Initially - at time 0 - there is a coin on the top-left cell of your board.
//Each cell of your board contains one of these letters:
//*: Exactly one of your cells has letter '*'.
//U: If at time t the coin is on cell (i, j) and cell (i, j) has letter 'U', the coin will be on cell (i-1, j) at time t+1, if i > 1. Otherwise, there is no coin on your board at time t + 1 .
//L: If at time  the coin is on cell (i, j) and cell (i, j) has letter 'L', the coin will be on cell (i, j - 1) at time t+1 , if j > 1. Otherwise, there is no coin on your board at time t + 1.
//D: If at time  the coin is on cell (i, j) and cell (i, j) has letter 'D', the coin will be on cell (i + 1, j) at time t+1, if i<N. Otherwise, there is no coin on your board at time t+1.
//R: If at time  the coin is on cell (i, j) and cell (i, j) has letter 'R', the coin will be on cell (i, j + 1) at time t+1, if j<M. Otherwise, there is no coin on your board at time t+1.
//When the coin reaches a cell that has letter '*', it will stay there permanently. When you punch on your board, your timer starts and the coin moves between cells.
//Before starting the game, you can make operations to change the board, such that you are sure that at or before time k the coin will reach the cell having letter '*'.
//In each operation you can select a cell with some letter other than '*' and change the letter to 'U', 'L', 'R' or 'D'.
//You need to carry out as few operations as possible in order to achieve your goal. Your task is to find the minimum number of operations.

func coinOnTheTable(m int32, k int32, board []string) int32 {

}
