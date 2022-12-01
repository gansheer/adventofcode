#! /usr/bin/env python3
# coding: utf-8
import itertools

def apply_mask_to_value(value, mask):
  print(value)
  retval = ''
  for v, m in zip(value, mask):
      if m == '0':
          retval += '0'
      elif m == '1':
          retval += '1'
      else:
          retval += v
  print(int(retval, 2))
  return int(retval, 2)

def convert_to_bytes(number_str):
  return "{0:b}".format(number_str).zfill(36)

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)

    mems = dict()
    mask = []
    for clean_line in clean_lines:
      if 'mask = ' in clean_line:
        mask = list(clean_line.split(' = ')[1].replace(' ',''))
        print(mask)
      else:
        position = int(clean_line.split(' = ')[0].replace('mem[','').replace(']',''))
        value = int(clean_line.split(' = ')[1].replace(' ',''))
        mems[position] = apply_mask_to_value(convert_to_bytes(value), mask)
        
    print(mems)
    print("Result: {}".format(sum(mems.values())))


if __name__ == '__main__':
    main()