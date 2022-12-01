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
        firstPosition = int(line.split('-')[0])
        secondPosition = int(line.split('-')[1].split(' ')[0])
        letter = line.split('-')[1].split(' ')[1].split(':')[0]
        password = line.split(': ')[1]
        firstValid = letter == password[firstPosition-1]
        secondValid = letter == password[secondPosition-1]
        print("Line : firstPosition={}, secondPosition={}, letter={}, password={}, firstValid={}, secondValid={}".format(firstPosition,secondPosition,letter,password, firstValid, secondValid)) 
        if firstValid ^ secondValid:
            valid_passwords += 1
    
    
    print("Result: {}".format(valid_passwords))

if __name__ == '__main__':
    main()