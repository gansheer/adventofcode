#! /usr/bin/env python3
# coding: utf-8
import itertools

def findInvalidNumber(clean_lines, step_size):
  numbers = list()
  for index, line in enumerate(clean_lines):
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
        return number, index
      else:
        numbers.pop(0)
        numbers.append(number)

  

def findInvalidNumberCombination(limited_clean_lines, invalid_number, step_size):
  print("findInvalidNumberCombination {}".format(invalid_number))
  print(limited_clean_lines)
  numbers = map(lambda x: int(x),limited_clean_lines)
  found = False
  combination_size = 3
  while not found and combination_size < len(limited_clean_lines):
    found = findNumberInvalid(step_size, numbers,invalid_number, combination_size)
    if not found:
      combination_size += 1
    else:
      print("found at combination_size {}".format(combination_size))


def findNumberInvalid(step_size, numbers, number, combination_size):
  combinations = ([numbers[i:i+combination_size] for i in range(0, len(numbers), 1)])
  #combinations = list(itertools.combinations(range(step_size), combination_size))
  for combination in combinations:
    expected_number = sum(map(lambda value: value, list(combination)))
    if number == expected_number:
      print("found at combination {}".format(combination))
      print("*\o/* value found is {}+{}={}".format(min(combination), max(combination), max(combination) + min(combination)))
      return True
  return False


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
    
    
    invalid_number, invalid_index = findInvalidNumber(clean_lines, step_size)
    findInvalidNumberCombination(clean_lines[0:invalid_index], invalid_number, step_size)
    
    print("Result: wrong number {} on line {}".format(invalid_number, invalid_index))

if __name__ == '__main__':
    main()