#! /usr/bin/env python3
# coding: utf-8
import itertools


def execute_round(ferry_map):
    new_ferry_map = list()
    for r_index in range((len(ferry_map))):
        #print("r_index {}".format(r_index))
        new_row = list()
        for c_index in range(len(ferry_map[r_index])):
            #print("r_index {}, c_index {}".format(r_index, c_index))
            new_seat = execute_round_seat(ferry_map, r_index, c_index)
            new_row.append(new_seat)
        new_ferry_map.append(new_row)
    return new_ferry_map


def execute_round_seat(old_ferry_map, r_index, c_index):
    #print(old_ferry_map)
    seat = old_ferry_map[r_index][c_index]
    if seat == '.':
        return '.'

    adjacent_seats = get_adjacent_seats(r_index, c_index, old_ferry_map)
    adjacent_occupied_seats = sum(
        map(lambda value: 1 if value == '#' else 0, adjacent_seats))
    #print("occupied {} for adjacent_seats {}({},{})".format(adjacent_occupied_seats, adjacent_seats, r_index, c_index))
    #print(adjacent_occupied_seats)
    if seat == 'L' and adjacent_occupied_seats == 0:
        return '#'
    if seat == '#' and adjacent_occupied_seats >= 5:
        return 'L'

    return seat


def get_adjacent_seats(r_index, c_index, ferry_map):
    max_r = len(ferry_map)
    max_c = len(ferry_map[r_index])
    #print("r_index {}, c_index {}, max_r {}, max_c {}".format(r_index,c_index,max_r,max_c))
    adjacent_seats = []
    for dr in (-1, 0, +1):
        for dc in (-1, 0, +1):
            if (dr != 0 or dc != 0):
                new_r = r_index + dr
                new_c = c_index + dc
                if new_r >= 0 and new_r < max_r and new_c >= 0 and new_c < max_c:
                    #print(dr,dc)
                    new_seat = get_first_signification_seat_in_direction(
                        r_index, c_index, max_r, max_c, dr, dc, ferry_map)
                    adjacent_seats.append(new_seat)

    return adjacent_seats


def get_first_signification_seat_in_direction(r_index, c_index, max_r, max_c,
                                              dr, dc, ferry_map):
    #print("get_first_signification_seat_in_direction")
    seats = list()
    new_r = r_index + dr
    new_c = c_index + dc
    while new_r >= 0 and new_r < max_r and new_c >= 0 and new_c < max_c:
        seats.append(ferry_map[new_r][new_c])
        new_r = new_r + dr
        new_c = new_c + dc

    #print("seats in direction {}".format(seats))
    for seat in seats:
        if seat == '#' or seat == 'L':
            return seat

    return '.'


def count_occupied_seats(ferry_map):
    flatten = lambda t: [item for ferry_map in t for item in ferry_map]
    #print(flatten(ferry_map))
    return flatten(ferry_map).count('#')


def print_ferry_map(ferry_map):
    print('')
    for index, row in enumerate(ferry_map):
        separator = ''
        print("{} {}".format(index, separator.join(row)))
    print('')


def main():
    input_file = open('input', 'r')
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"), lines)
    print(clean_lines)

    ferry_map = list()
    for clean_line in clean_lines:
        line = list(clean_line)
        ferry_map.append(line)

    #print(ferry_map)
    print_ferry_map(ferry_map)
    found = False
    count = 1
    while not found:
        print("<<<<<<< ROUND {} >>>>>>".format(count))
        rows_num = len(ferry_map)
        columns_num = len(ferry_map[0])
        print("ferry_map size {}/{}".format(rows_num, columns_num))
        new_ferry_map = execute_round(ferry_map)
        print_ferry_map(new_ferry_map)
        found = (ferry_map == new_ferry_map)
        ferry_map = new_ferry_map
        count += 1
        #if count > 4:
        #  break

    #new_ferry_map = execute_round(ferry_map)
    #print(new_ferry_map)
    #print(ferry_map == new_ferry_map)
    #ferry_map = new_ferry_map
    #new_ferry_map = execute_round(ferry_map)
    #print(new_ferry_map)
    #print(ferry_map == new_ferry_map)

    print("Result: {}".format(count_occupied_seats(ferry_map)))


if __name__ == '__main__':
    main()