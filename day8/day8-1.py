#! /usr/bin/env python3
# coding: utf-8

def execute_command(clean_lines, index_line, accumulator):
  command = clean_lines[index_line].split()[0]
  operation = clean_lines[index_line].split()[1]
  #print("command {} / operation {} / index_line {} / accumulator ".format(command,operation, index_line, accumulator))
  if command == 'nop':
    return (index_line+1), (accumulator)
  elif command == 'jmp':
    return eval(str(index_line) + operation), accumulator
  else:
    return (index_line+1), eval(str(accumulator) + operation)
    
def execute_program(modified_lines):
  has_loop = False
  accumulator = 0
  new_index = 0
  visited_indexes = []
  
  while len(modified_lines) > new_index :
    if new_index in visited_indexes:
      print("already visited {}".format(new_index))
      has_loop = True
      #print("{} indexes alreadu visited :{}".format(len(visited_indexes), visited_indexes))
      break
    else:
      #print("execute_command {} with accumulator {}".format(new_index, accumulator))
      visited_indexes.append(new_index)
      new_index, accumulator = execute_command(modified_lines, new_index, accumulator)
  
  return accumulator, has_loop

def change_program(clean_lines, change_try_number):
  modified_lines = []
  counter = 0
  
  for line in clean_lines:

    if line.split()[0] in ['nop','jmp']:
      if counter == change_try_number:
        if line.split()[0] == 'nop':
          print("Modifying {} to jmp".format(line))
          modified_lines.append(line.replace('nop','jmp'))
        else:
          print("Modifying {} to nop".format(line))
          modified_lines.append(line.replace('jmp','nop'))
      else:
        modified_lines.append(line)
      counter += 1
    else:
      modified_lines.append(line)
  
  return modified_lines

    
def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)
    
    change_try_number = 0
    accumulator = 0
    has_loop = True
    while has_loop:
      modified_lines = change_program(clean_lines, change_try_number)
      #print(modified_lines)
      accumulator, has_loop = execute_program(modified_lines)
      change_try_number += 1

      
    

    print("Result: accumulator={}/has_loop={}".format(accumulator, has_loop))

if __name__ == '__main__':
    main()