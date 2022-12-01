#! /usr/bin/env python3
# coding: utf-8
import itertools

def findAdaptersByDifference(device_joltage, adapters):
  adapters_1 = []
  adapters_3 = []
  current_charging_outlet = 0
  ordered_adapter = sorted(adapters)
  print(ordered_adapter)
  for adapter in ordered_adapter:
    print("Adapter {} for {}".format(adapter, device_joltage))
    if (adapter - 1) == current_charging_outlet:
      print("Found adapters_1 {} for {}".format(adapter, current_charging_outlet))
      adapters_1.append(adapter)
      current_charging_outlet = adapter
    elif adapter - 2 == current_charging_outlet:
      print("Found adapters_2 {} for {}".format(adapter, current_charging_outlet))
      current_charging_outlet = adapter
    elif adapter - 3 == current_charging_outlet:
      print("Found adapters_3 {} for {}".format(adapter, current_charging_outlet))
      adapters_3.append(adapter)
      current_charging_outlet = adapter
  
  adapters_3.append(device_joltage)
  return adapters_1, adapters_3




def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: int(x.rstrip("\n")),lines)
    print(clean_lines)
    
    device_joltage = max(clean_lines) + 3
    result = findAdaptersByDifference(device_joltage, clean_lines)
    print("Result: {} = {} for 1, {} for 3".format(len(result[0])*len(result[1]), result[0], result[1]))
      


if __name__ == '__main__':
    main()