#! /usr/bin/env python3
# coding: utf-8

# why does it work with -1...

def find_children(bag, clean_lines):
  for line in clean_lines:
    lineBag = line.replace(".","").replace(" ","").replace("bags","").replace("bag","").split("contain")[0]
    print(lineBag)
    lineSubBags =  line.replace(".","").replace(" ","").replace("bags","").replace("bag","").split("contain")[1]
    print(lineSubBags)
    if bag.replace(" ", "") in lineBag:
      if "noother" in lineSubBags:
        return 1
      total = 0
      for child in lineSubBags.split(","):
        total += int(child[0]) * find_children(child[1:], clean_lines)
      return total + 1

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)

    result = find_children('shiny gold', clean_lines)

    print("Result: {}".format(result - 1 ))

if __name__ == '__main__':
    main()