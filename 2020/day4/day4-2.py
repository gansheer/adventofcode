#! /usr/bin/env python3
# coding: utf-8
import re

def isPassportValid(passport):
    print('byr:{}'.format(isPassportYearValid(passport,'byr',1920, 2020)))
    print('iyr:{}'.format(isPassportYearValid(passport,'iyr',2010, 2020)))
    print('eyr:{}'.format(isPassportYearValid(passport,'eyr',2020, 2030)))
    print('hgt:{}'.format(isPassportHeightValid(passport)))
    print('hcl:{}'.format(isPassportHairColorValid(passport)))
    print('ecl:{}'.format(isPassportEyeColorValid(passport)))
    print('pid:{}'.format(isPassportPidValid(passport)))

    return isPassportYearValid(passport,'byr',1920, 2020) and \
                isPassportYearValid(passport,'iyr',2010, 2020) and \
                isPassportYearValid(passport,'eyr',2020, 2030) and \
                isPassportHeightValid(passport) and \
                isPassportHairColorValid(passport) and \
                isPassportEyeColorValid(passport) and \
                isPassportPidValid(passport)


def isPassportYearValid(passport, key, min, max):
    return (key in passport) and len(passport[key]) == 4 and int(passport[key]) >= min and  int(passport[key]) <= max

def isPassportHeightValid(passport):
    if 'hgt' not in passport:
        return False
    if 'cm' in passport['hgt']:
        return (int(passport['hgt'].split('cm')[0]) >= 150) and  (int(passport['hgt'].split('cm')[0]) <= 193)
    elif 'in' in passport['hgt']:
        return (int(passport['hgt'].split('in')[0]) >= 59) and  (int(passport['hgt'].split('in')[0]) <= 76)
    return False


def isPassportHairColorValid(passport):
    if ('hcl' not in passport):
        return False
    return len(passport['hcl']) == 7 and bool(re.match('#[0-9a-f]{6}', passport['hcl'])) == True

def isPassportEyeColorValid(passport):
    if ('ecl' not in passport):
        return False
    for color in ['amb','blu','brn','gry','grn','hzl','oth']:
        if passport['ecl'].strip() == color:
            return True
    return False

def isPassportPidValid(passport):
    if ('pid' not in passport):
        return False
    return len(passport['pid']) == 9 and bool(re.match('[0-9]{9}', passport['pid'])) == True

def main():
    file1 = open('input', 'r') 
    Lines = file1.readlines()

    valid_passports = 0
    passport = dict()

    for line in Lines:
        print(line)
        if line.strip() == '':
            print("empty line")
            if isPassportValid(passport):
                print("valid password")
                valid_passports += 1
            passport = dict()
        else:
            datas = line.replace('\n','').split(' ');
            for data in datas:
                print("data: {}, key=[{}], value=[{}]".format(data, data.split(':')[0], data.split(':')[1]))
                passport[data.split(':')[0]] = data.split(':')[1]

    
    if passport != {}:
        if isPassportValid(passport):
                print("valid password")
                valid_passports += 1

    print("Result: {}".format(valid_passports))

if __name__ == '__main__':
    main()