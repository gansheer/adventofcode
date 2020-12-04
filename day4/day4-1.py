#! /usr/bin/env python3
# coding: utf-8


def main():
    file1 = open('input_test', 'r') 
    Lines = file1.readlines()

    valid_passports = 0
    passport = dict()

    for line in Lines:
        print(line)
        if line.strip() == '':
            print("empty line")
            valid = ('byr' in passport) and \
                ('iyr' in passport) and \
                ('eyr' in passport) and \
                ('hgt' in passport) and \
                ('hcl' in passport) and \
                ('ecl' in passport) and \
                ('pid' in passport)
            if valid:
                valid_passports += 1
            passport = dict()
        else:
            datas = line.replace('\n','').split(' ');
            for data in datas:
                print("data: {}, key=[{}], value=[{}]".format(data, data.split(':')[0], data.split(':')[1]))
                passport[data.split(':')[0]] = data.split(':')[1]

    
    print("Result: {}".format(valid_passports))

if __name__ == '__main__':
    main()