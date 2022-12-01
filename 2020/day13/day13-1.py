#! /usr/bin/env python3
# coding: utf-8
import itertools

def find_next_bus(current_timestamp, buses):
  waiting_timestamp = current_timestamp
  while True:
    for bus in buses:
      if waiting_timestamp % bus == 0:
        return waiting_timestamp, bus
    waiting_timestamp += 1
  return 0,0

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)
    print(clean_lines)

    current_timestamp = int(clean_lines[0])
    buses = sorted(map(lambda x: int(x),list(filter(lambda a: a != 'x', clean_lines[1].split(',')))))
    print(current_timestamp)
    print(buses)
    result = find_next_bus(current_timestamp, buses)
    print(result)

    print("Result: {}".format((result[0]-current_timestamp)*result[1]))


if __name__ == '__main__':
    main()