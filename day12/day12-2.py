#! /usr/bin/env python3
# coding: utf-8
import itertools

def execute_command(current_x, current_y, waypoint_x, waypoint_y, command):
  if command[0] == 'F':
    current_x, current_y = move_forward(current_x, current_y, waypoint_x, waypoint_y, int(command[1:]))
  if command[0] == 'R' or command[0] == 'L':
    waypoint_x, waypoint_y = change_direction(waypoint_x, waypoint_y, command[0], int(command[1:]))
  if command[0] == 'E' or command[0] == 'W' or command[0] == 'S' or command[0] == 'N' :
    waypoint_x, waypoint_y = move_direction(waypoint_x, waypoint_y, command[0], int(command[1:]))

  print("Waypoint W - |{}| - E /N - |{}| - S".format(waypoint_x, waypoint_y))
  return current_x, current_y, waypoint_x, waypoint_y

def move_forward(current_x, current_y, waypoint_x, waypoint_y, length):
  current_x += (waypoint_x * length)
  current_y += (waypoint_y * length)
  return current_x, current_y

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

def change_direction(waypoint_x, waypoint_y, turn_direction, length):
  if turn_direction == 'R':
    while length > 0:
      waypoint_x, waypoint_y = go_right(waypoint_x, waypoint_y)
      length -= 90
  if turn_direction == 'L':
    while length > 0:
      waypoint_x, waypoint_y = go_left(waypoint_x, waypoint_y)
      length -= 90
  return waypoint_x, waypoint_y

def go_right(waypoint_x, waypoint_y):
    new_waypoint_x = 0 - waypoint_y
    new_waypoint_y = waypoint_x
    return new_waypoint_x, new_waypoint_y
  
def go_left(waypoint_x, waypoint_y):
    new_waypoint_x = waypoint_y
    new_waypoint_y = 0 - waypoint_x
    return new_waypoint_x, new_waypoint_y

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)
    print(clean_lines)

    current_x = 0
    current_y = 0
    waypoint_x = 10
    waypoint_y = -1
    for clean_line in clean_lines:
      current_x, current_y, waypoint_x, waypoint_y = execute_command(current_x, current_y, waypoint_x, waypoint_y, clean_line)
      print("Ship is : W - |{}| - E /N - |{}| - S, going to W - |{}| - E /N - |{}| - S".format(current_x, current_y, waypoint_x, waypoint_y))

    print("Result: {}".format(abs(current_x)+abs(current_y)))


if __name__ == '__main__':
    main()