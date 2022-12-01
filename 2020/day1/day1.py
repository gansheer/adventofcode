#! /usr/bin/env python3
# coding: utf-8


def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    numbers_length = 0
    numbers = []
    # Strips the newline character 
    for line in Lines: 
        print("Line{}: {}".format(numbers_length, line.strip())) 
        numbers_length += 1
        numbers.append(int(line.strip()))
    
    
    #numbers = sorted(numbers)

    for number in numbers:
        print("Number: {}".format(number)) 

    result = 0
    firstCounter = 0
    while ((result == 0) and (firstCounter < numbers_length)):
        firstNumber = numbers[firstCounter]
        print("firstNumber/firstCounter: {}/{}".format(firstNumber, firstCounter))
        secondCounter = 0
        while result == 0 and secondCounter < numbers_length:
            secondNumber = numbers[secondCounter]
            print("secondNumber/secondCounter: {}/{}".format(secondNumber, secondCounter))
            if secondCounter != firstCounter:
                thirdCounter = 0
                while result == 0 and thirdCounter < numbers_length:
                    if thirdCounter != firstCounter and thirdCounter != secondCounter :
                        thirdNumber = numbers[thirdCounter]
                        if firstNumber + secondNumber + thirdNumber == 2020:
                            print("Them: {}/{}/{}".format(firstNumber,secondNumber,thirdNumber))
                            result = firstNumber * secondNumber * thirdNumber
                        else:
                            print("Not them: {}/{}/{}".format(firstNumber,secondNumber,thirdNumber))
                    thirdCounter += 1
            secondCounter += 1
        firstCounter += 1
    
    print("Result: {}".format(result))

if __name__ == '__main__':
    main()
