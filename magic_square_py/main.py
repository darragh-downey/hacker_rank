#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'formingMagicSquare' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY s as parameter.
#
# The eight valid combinations that sum to '15'
#
#    9 5 1
#    7 5 3
#    2 5 8
#    4 5 6
#    2 9 4
#    6 1 8
#    6 7 2
#    8 3 4
#
# Key to magic square is that arr[1][1] == 5
# s[0][0] top left
# s[0][2] top right
# s[2][0] bottom left
# s[2][2] bottom right


def formingMagicSquare(s):
    moves = 0
    combos = [
        [9,5,1], [7,5,3], [2,5,8], [4,5,6], 
        [2,9,4], [6,1,8], [6,7,2], [8,3,4]
    ]

    if s[1][1] != 5:
        # big change
        moves += abs(5 - s[1][1])
        s[1][1] = 5
    
    checkDiagonals(s, combos, moves)
    checkRows(s, combos, moves)
    checkCols(s, combos, moves)

    return moves
    

def checkDiagonals(s, combos, moves):
    corners = [[9, 1], [4, 6]]
    if [s[i][i] for i in range(len(s))] in combos and s[1][1] == 5:
        pass
    if [s[i][len(s)-1-i] for i in range(len(s))] in combos and s[1][1] == 5:
        pass
    
    for corner in corners:
        for c in range(len(corner)):
            if s[0][0] == corner[c]:
                moves += abs(s[2][2] - corner[c-len(corner)])
                s[2][2] = corner[c-len(corner)]
            if s[2][2] == corner[c]:
                moves += abs(s[0][0] - corner[c-len(corner)])
                s[0][0] = corner[c-len(corner)]
            
            if s[2][0] == corner[c]:
                moves += abs(s[0][2] - corner[c-len(corner)])
                s[0][2] = corner[c-len(corner)]
            if s[0][2] == corner[c]:
                moves += abs(s[2][0] - corner[c-len(corner)])
                s[2][0] = corner[c-len(corner)]

    pass


def checkRows(s):
    pass


def checkCols(s):
    pass


def find15(s):
    rows = s
    cols = [[s[i][j] for i in range(len(s))] for j in range(len(s))]
    d_a = [s[i][i] for i in range(len(s))]
    d_b = [s[i][len(s)-1-i] for i in range(len(s))]
    diags = [d_a, d_b]

    sq = [*rows, *cols, *diags]
    idx = [sum(i) for i in sq]

    return sq, idx

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = [[4,9,2], [3,5,7], [8,1,5]]

    #for _ in range(3):
    #    s.append(list(map(int, input().rstrip().split())))

    sq, sums = formingMagicSquare(s)
    print("Square", sq, "\nSums", sums)

    #fptr.write(str(result) + '\n')

    #fptr.close()
