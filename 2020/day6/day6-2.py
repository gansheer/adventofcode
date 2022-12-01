#! /usr/bin/env python3
# coding: utf-8
from itertools import groupby


def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),Lines)
    print(clean_lines)

    result = 0
    groups = [list(g) for k,g in groupby(clean_lines,lambda line:line=='') if not k]
    for group in groups:
        lst = [set(x) for x in group]
        print("Group: {}, intersection: {}".format(group, len(lst[0].intersection(*lst))))
        result += len(lst[0].intersection(*lst))
    
    print("Result: {}".format(result))

if __name__ == '__main__':
    main()