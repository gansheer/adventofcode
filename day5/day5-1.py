#! /usr/bin/env python3
# coding: utf-8

def getLower(min, max):
    print("getLower {}/{}".format(min, max))
    return ((max - min) // 2) + min

def getUpper(min, max):
    print("getUpper {}/{}".format(min, max))
    return ((max - min) // 2) + min + 1

def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    max_id = 0

    for line in Lines:
        print("***** line: {} *****".format(line))
        minRow =  0
        maxRow =  127
        minColumn = 0
        maxColumn = 7
        for letter in line.rstrip("\n"):
            print("get {}".format(letter))
            if letter == 'F':
                maxRow = getLower(minRow, maxRow);
            elif letter == 'B':
                minRow = getUpper(minRow, maxRow);
            elif letter == 'L':
                maxColumn = getLower(minColumn, maxColumn);
            elif letter == 'R':
                minColumn = getUpper(minColumn, maxColumn);
            print("letter: {}, minRow/maxRow=[{}/{}], maxColumn/maxColumn=[{}/{}]".format(letter, minRow, maxRow, minColumn, maxColumn))

        if ((minRow * 8) + minColumn) > max_id:
            max_id = (minRow * 8) + minColumn
        
    
    print("Result: {}".format(max_id))

if __name__ == '__main__':
    main()