#! /usr/bin/env python3
# coding: utf-8


def main():
    file1 = open('input_test', 'r') 
    Lines = file1.readlines()

    number_trees = 0
    current_x = 0
    # Strips the newline character 
    for line in Lines:
        if current_x < len(line) and line[current_x] == '#':
            number_trees += 1
        print("Line : number_trees={}, current_x={}, char={}, line={}".format(number_trees,current_x, line[current_x], line)) 
        current_x = (current_x + 3) % (len(line)-1)
        
    
    
    print("Result: {}".format(number_trees))

if __name__ == '__main__':
    main()