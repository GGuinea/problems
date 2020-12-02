def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def get_range(string):
    rangeString = string.split('-')
    return int(rangeString[0]), int(rangeString[1])

def get_character(string):
    character = string.split(":")
    return character[0]

lines = read_lines('input')
answer = 0
for line in lines:
    print("{}".format(line.strip()))
    splitted = line.split(' ')
    min, max = get_range(splitted[0])
    character = get_character(splitted[1])
    splitted[2] = splitted[2][:-1]

    if (splitted[2][min-1] == character):
        print(splitted[2], character, min, max)
        if len(splitted[2]) > max-1:
            if splitted[2][max-1] != character:
                answer += 1
                continue

    if (splitted[2][min-1] != character):
        print(splitted[2], character, min, max)
        if len(splitted[2]) > max-1:
            if splitted[2][max-1] == character:
                answer += 1

    print("answer", answer),
print(answer)



