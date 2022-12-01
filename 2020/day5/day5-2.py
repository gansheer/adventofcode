#! /usr/bin/env python3
# coding: utf-8

def getLower(min, max):
    return ((max - min) // 2) + min

def getUpper(min, max):
    return ((max - min) // 2) + min + 1

def findEmptySeat(board, maxRow, maxColumn):

    counter = 0
    for row in board:
        print("{}:{}".format(counter,row))
        counter += 1
    
    print(maxRow)
    print(maxColumn)

    found = False
    started = False
    for rowNum in range(maxRow):
        for columnNum in range(maxColumn):
            #print("Seat {}/{}={}".format(rowNum, columnNum, board[rowNum][columnNum]))
            seat = board[rowNum][columnNum]
            
            if started == False and seat == 1:
                started = True

            if seat == 0 and started == True and found == False:
                found = True
                resultRow  = rowNum
                resultColumn = columnNum
                print("Result {}/{} : ID={}".format(resultRow, resultColumn, ((resultRow * maxColumn) + resultColumn)))
                

def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    board = [[0 for i in range(8)] for j in range(128)]


    for line in Lines:
        minRow =  0
        maxRow =  127
        minColumn = 0
        maxColumn = 7
        for letter in line.rstrip("\n"):
            if letter == 'F':
                maxRow = getLower(minRow, maxRow);
            elif letter == 'B':
                minRow = getUpper(minRow, maxRow);
            elif letter == 'L':
                maxColumn = getLower(minColumn, maxColumn);
            elif letter == 'R':
                minColumn = getUpper(minColumn, maxColumn);
            
        print("Line {}: row/column={}/{}".format(line.rstrip("\n"), minRow, minColumn))
        board[minRow][minColumn] = 1
    
    findEmptySeat(board, 128, 8)


                
    #print("Result {}/{} : ID={}".format(resultRow, resultColumn, ((resultRow * 8) + resultColumn)))
    

if __name__ == '__main__':
    main()