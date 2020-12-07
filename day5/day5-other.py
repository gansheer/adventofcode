#! /usr/bin/env python3
# coding: utf-8

# get file contents
file = open("input", 'r')
contents = file.readlines()
file.close()

# converts the string portion to a decimal integer based on binary
def getPosition(str, one):
    """ 
    needs to know which digit is considered 1 because based in the
    samples provided by the prompt, the first digit isn't always 1.
    """
    return int(''.join(('1' if c == one else '0') for c in str), 2)

# create a list of boarding passes as dicts
boardingPasses = []
for string in contents:
    passDict = {
        'row': getPosition(string.strip()[:-3], 'B'),
        'col': getPosition(string.strip()[-3:], 'R')
    }
    passDict['id'] = passDict['row'] * 8 + passDict['col']
    boardingPasses.append(passDict)

# get highest ID in the boarding passes
highest = max([p['id'] for p in boardingPasses])

# print the result
print(highest)

"""
PART 2
Find a boarding pass with an ID that falls +1/-1 between two other passes. 
"""

# order the list of boarding pass dicts by their id
boardingPasses.sort(key=lambda p : p['id'])

i = 1
while i < len(boardingPasses):
    # check this plus the one below it for a 2 difference in id
    priorID = boardingPasses[i - 1]['id']
    thisID = boardingPasses[i]['id']
    if thisID - priorID == 2:
       # print the answer
       print(thisID - 1)
    
    i += 1