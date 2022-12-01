#! /usr/bin/env python3
# coding: utf-8
import itertools
from math import log2

# should not have tried on test values XD

def apply_mask_to_position(position, mask):
    print(position)
    new_positions = []
    binary = list(convert_to_binary(position))
    floating_indexes = []
    for i, m in enumerate(mask):
        if m == '1':
            binary[i] = '1'
        elif m == 'X':
            binary[i] = 'X'
            floating_indexes.append(i)
    num_addresses = 2 ** len(floating_indexes)
    bits_used = int(log2(num_addresses))
    for i in range(num_addresses):
        print(i)
        alternate_bits = bin(i)[2:].zfill(bits_used)
        for j, b in enumerate(alternate_bits):
            binary_index = floating_indexes[j]
            binary[binary_index] = b
        new_positions.append(int("".join(binary), 2))
    return new_positions

def convert_to_binary(number_str):
  return "{0:b}".format(number_str).zfill(36)

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)

    mems = dict()
    mask = []
    for clean_line in clean_lines:
      print(clean_line)
      if 'mask = ' in clean_line:
        mask = list(clean_line.split(' = ')[1].replace(' ',''))
        print(mask)
      else:
        position = int(clean_line.split(' = ')[0].replace('mem[','').replace(']',''))
        value = int(clean_line.split(' = ')[1].replace(' ',''))
        positions_to_change = apply_mask_to_position(position, mask)
        for new_position in positions_to_change:
                mems[new_position] = value
        
    print(mems)
    print("Result: {}".format(sum(mems.values())))


if __name__ == '__main__':
    main()