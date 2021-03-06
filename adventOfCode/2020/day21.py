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

    charCounter = 0
    for char in splitted[2]:
        if char == character:
            charCounter+=1
            if charCounter > max:
                break
    if charCounter <= max and charCounter >= min:
        answer += 1
    print("min", min, "max", max, "acutal", charCounter)
    print("answer", answer),
print(answer)



