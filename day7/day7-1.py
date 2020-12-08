#! /usr/bin/env python3
# coding: utf-8

def find_roots(bag, clean_lines, root_bags):
  for line in clean_lines:
    lineName = line.split(" bags contain")[0]
    lineSubNames = line.split(" bags contain")[1]
    if bag in lineSubNames:
      if lineName not in root_bags:
        root_bags.append(lineName)
        find_roots(lineName, clean_lines, root_bags)
    

def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)

    #bags = []
    #bags['bright white'] = ['shiny gold']
    #bags['muted yellow'] = ['shiny gold']
    #bags['dark orange'] = ['bright white', 'muted yellow']
    #bags['light red'] = ['bright white', 'muted yellow']

    result = 0
    bags = list()
    bagNames = set()
    root_bags = []
    find_roots('shiny gold', clean_lines, root_bags)

    print("Result: {}".format(len(root_bags)))

if __name__ == '__main__':
    main()