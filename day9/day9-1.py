#! /usr/bin/env python3
# coding: utf-8
import itertools

def findNumber(step_size, numbers, number):
  combinations = list(itertools.combinations(range(step_size), 2))
  for left, right in combinations:
    expected_number = numbers[left] + numbers[right]
    if number == expected_number:
      return True
  return False

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)
    step_size = 25
    
    numbers = list()
    for line in clean_lines:
      print("Numbers : {}".format(numbers))
      if len(numbers) < step_size:
        numbers.append(int(line))
      else:
        print(numbers)
        number = int(line)
        found = findNumber(step_size, numbers,number)
        print("found {} ? {}".format(number, found))
        if not found:
          print("Result: wrong number {}".format(number))
          return
        else:
          numbers.pop(0)
          numbers.append(number)
      


if __name__ == '__main__':
    main()