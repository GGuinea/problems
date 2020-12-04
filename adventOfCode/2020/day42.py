import re

def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

def is_not_blank(s):
    return bool(s and s.strip())

def set_defailt(requiredFields):
    requiredFields['byr'] = 0
    requiredFields['iyr'] = 0
    requiredFields['eyr'] = 0
    requiredFields['hgt'] = 0
    requiredFields['hcl'] = 0
    requiredFields['ecl'] = 0
    requiredFields['pid'] = 0
    return requiredFields

def check_if_present(requiredFields):
    for x in requiredFields.values():
        if x == 0:
            return 0
    return 1

def check_byr(line, startPosition):
    byr_start= 1920
    byr_end= 2002
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    if int(splitted[1]) >= byr_start and int(splitted[1]) <= byr_end:
        return 1
    else:
        return 0
    
def check_iyr(line, startPosition):
    iyr_start= 2010
    iyr_end= 2020 
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    if int(splitted[1]) >= iyr_start and int(splitted[1]) <= iyr_end:
        return 1
    else:
        return 0

def check_eyr(line, startPosition):
    eyr_start = 2020
    eyr_end = 2030
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    if int(splitted[1]) >= eyr_start and int(splitted[1]) <= eyr_end:
        return 1
    else:
        return 0

def check_hgt(line, startPosition):
    hgt_start_cm = 150
    hgt_end_cm = 193
    hgt_start_in = 59
    hgt_end_in = 76
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    stop = 0
    start = 0
    if splitted[1].find('in') != -1:
        start = hgt_start_in
        stop = hgt_end_in
    else:
        start = hgt_start_cm
        stop = hgt_end_cm
    hgt = int(re.search(r'\d+', splitted[1]).group())
    if hgt >= start and hgt <= stop:
        return 1
    else:
        return 0

def check_hcl(line, startPosition):
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    x = re.search("#[0-9a-f]{6}", splitted[1])
    if len(splitted[1]) == 7 and x != None:
        return 1
    else:
        return 0

def check_ecl(line, startPosition):
    eyeColors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    for col in eyeColors:
        if splitted[1] == col:
            return 1

    return 0

def check_pid(line, startPosition):
    nubmerOfDigits = 9
    space = line.find(" ", startPosition, len(line))
    if space == -1:
        space = len(line)
    splitted = line[startPosition:space].split(':')
    x = re.search("[0-9]{9}", splitted[1])
    if len(splitted[1]) == nubmerOfDigits and x != None:
        return 1
    else:
        return 0

lines = read_lines('input')
line = remove_end_of_line(lines)
requiredFields = {'byr':0, 'iyr':0, 'eyr':0, 'hgt':0, 'hcl':0, 'ecl':0, 'pid':0}
answer = 0
for line in lines:
    if len(line) == 0:
        answer += check_if_present(requiredFields) 
        requiredFields = set_defailt(requiredFields)
    else:
        for x in requiredFields:
            startPosition = line.find(x)
            if startPosition != -1:
                if x == 'byr':
                    requiredFields[x] = check_byr(line, startPosition)
                if x == 'iyr':
                    requiredFields[x] = check_iyr(line, startPosition)
                if x == 'eyr':
                    requiredFields[x] = check_eyr(line, startPosition)
                if x == 'hgt':
                    requiredFields[x] = check_hgt(line, startPosition)
                if x == 'hcl':
                    requiredFields[x] = check_hcl(line, startPosition)
                if x == 'ecl':
                    requiredFields[x] = check_ecl(line, startPosition)
                if x == 'pid':
                    requiredFields[x] = check_pid(line, startPosition)

print(answer)


