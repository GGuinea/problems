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
            if line.find(x) is not -1:
                requiredFields[x] = 1

print(answer)


