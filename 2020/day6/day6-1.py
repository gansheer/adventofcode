#! /usr/bin/env python3
# coding: utf-8


def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    result = 0
    group = set()

    for line in Lines:
        print("***** line: {} *****".format(line.rstrip("\n")))
        if line.strip() == '':
            result += len(group)
            group = set()
            print("Group size {}".format(result))
        else:
            group = group.union(line.rstrip("\n"))
    
    result += len(group)
        
    
    print("Result: {}".format(result))

if __name__ == '__main__':
    main()