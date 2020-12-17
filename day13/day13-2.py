#! /usr/bin/env python3
# coding: utf-8
# Thanks to imbadatreading, should have through about LCM

def add_step_size(buses):
    buses_with_step = list()
    for num in range(len(buses)):
        if buses[num] != 'x':
            buses_with_step.append((int(buses[num]), num))
    return buses_with_step


def find_awesome_timestamp(buses):
    current_time = 0
    lcm = 1
    for order_bus in range(len(buses) - 1):
        bus = int(buses[order_bus + 1][0])
        step = int(buses[order_bus + 1][1])
        lcm = lcm * buses[order_bus][0]
        print("bus={}/step={}/lcm={}".format(bus, step, lcm))
        while ((current_time + step) % bus) != 0:
          current_time += lcm

    return current_time


def main():
    input_file = open('input', 'r')
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"), lines)
    print(clean_lines)

    buses = add_step_size(list(clean_lines[1].split(',')))
    print(buses)
    result = find_awesome_timestamp(buses)
    print(result)

    print("Result: {}".format(result))


if __name__ == '__main__':
    main()