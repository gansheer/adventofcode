#! /usr/bin/env python3
# coding: utf-8
import itertools

def execute_command(current_x, current_y, current_direction, command):
  if command[0] == 'F':
    current_x, current_y = move_direction(current_x, current_y, current_direction, int(command[1:]))
  if command[0] == 'R' or command[0] == 'L':
    current_direction = change_direction(current_direction, command[0], int(command[1:]))
  if command[0] == 'E' or command[0] == 'W' or command[0] == 'S' or command[0] == 'N' :
    current_x, current_y = move_direction(current_x, current_y, command[0], int(command[1:]))

  return current_x, current_y, current_direction

def move_direction(current_x, current_y, direction, length):
  if direction == 'E':
    current_x += length
  if direction == 'W':
    current_x -= length
  if direction == 'S':
    current_y += length
  if direction == 'N':
    current_y -= length
  return current_x, current_y

def change_direction(current_direction, turn_direction, length):
  if turn_direction == 'R':
    while length > 0:
      current_direction = go_right(current_direction)
      length -= 90
  if turn_direction == 'L':
    while length > 0:
      current_direction = go_left(current_direction)
      length -= 90
  return current_direction

def go_right(current_direction):
  if current_direction == 'N':
    return 'E'
  if current_direction == 'E':
    return 'S'
  if current_direction == 'S':
    return 'W'
  if current_direction == 'W':
    return 'N'
  
def go_left(current_direction):
  if current_direction == 'N':
    return 'W'
  if current_direction == 'W':
    return 'S'
  if current_direction == 'S':
    return 'E'
  if current_direction == 'E':
    return 'N'

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)
    print(clean_lines)

    current_x = 0
    current_y = 0
    current_direction = 'E'
    for clean_line in clean_lines:
      current_x, current_y, current_direction = execute_command(current_x, current_y, current_direction, clean_line)
      print("Ship is : W - {} - E /N - {} - S, going to {}".format(current_x, current_y, current_direction))

    print("Result: {}".format(abs(current_x)+abs(current_y)))


if __name__ == '__main__':
    main()