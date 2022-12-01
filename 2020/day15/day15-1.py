#! /usr/bin/env python3
# coding: utf-8

# Thank to https://gist.github.com/joshbduncan/bfb162a038850f72b505a9baa42b3d01, got lost each time I tried in the problem statement

import itertools

def solve(match, starting_numbers):
    pdict = {}
    for i, p in enumerate(starting_numbers):
        pdict[int(p)] = i + 1

    print(pdict)

    turn = len(pdict) + 1
    same = 0

    while turn < match:
        print(pdict)
        if same in pdict:
            diff = turn - pdict[same]
            pdict[same] = turn
        else:
            diff = 0
            pdict[same] = turn

        same = diff
        turn += 1

    return same
  
def main():
    input_file = open('input', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: x.rstrip("\n"),lines)

    starting_numbers = list(clean_lines)[0].split(',')
    mem = solve(2020, starting_numbers)


    print(mem)
    print("Result: {}".format(mem))


if __name__ == '__main__':
    main()