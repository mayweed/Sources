import java.util.*
import java.io.*
import java.math.*

/*
You get a sudoku grid from a player and you have to check if it has been correctly
filled.

A sudoku grid consists of 9×9 = 81 cells split in 9 sub-grids of 3×3 = 9 cells.
For the grid to be correct, each row must contain one occurrence of each digit (1 to
9), each column must contain one occurrence of each digit (1 to 9) and each sub-grid
must contain one occurrence of each digit (1 to 9).

You shall answer true if the grid is correct or false if it is not.
*/

fun main(args : Array<String>) {
   val input = Scanner(System.`in`)
   var m = Array(9, {i -> Array(9, {j -> 0})}) 
    for (i in 0 until 9) {
        for (j in 0 until 9) {
            val n = input.nextInt()
            m[i][j]=n
        }
    }

    // Write an answer using println()
System.err.println(m[1][3]);

    println("true or false")
}
