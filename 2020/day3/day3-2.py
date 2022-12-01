#! /usr/bin/env python3
# coding: utf-8

def num_trees(lines,right,down):
    number_trees = 0
    current_x = 0
    # Strips the newline character 
    for line_number in xrange(0,len(lines),down):
        line = lines[line_number] 
        if current_x < len(line) and line[current_x] == '#':
            number_trees += 1
        print("Line : number_trees={}, current_x={}, char={}, line={}, lineNum={}".format(number_trees,current_x, line[current_x], line, line_number)) 
        current_x = (current_x + right) % (len(line)-1)
    return number_trees


def main():
    file1 = open('input', 'r') 
    lines = file1.readlines()
    
    result = 1
    slopes = [[1,1], [3,1], [5,1], [7,1], [1,2]]
    for slope in slopes:
        result *= num_trees(lines, slope[0], slope[1])
    print("Result: {}".format(result))

if __name__ == '__main__':
    main()