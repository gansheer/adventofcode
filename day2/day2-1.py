#! /usr/bin/env python3
# coding: utf-8


def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    valid_passwords = 0
    numbers = []
    # Strips the newline character 
    for line in Lines:
        print(line) 
        minOccurences = int(line.split('-')[0])
        maxOccurences = int(line.split('-')[1].split(' ')[0])
        letter = line.split('-')[1].split(' ')[1].split(':')[0]
        password = line.split(': ')[1]
        occurences = password.count(letter)
        print("Line : min={}, max={}, letter={}, password={}, occurences={}".format(minOccurences,maxOccurences,letter,password, occurences)) 
        if occurences >= minOccurences and occurences <= maxOccurences:
            valid_passwords += 1
    
    
    print("Result: {}".format(valid_passwords))

if __name__ == '__main__':
    main()