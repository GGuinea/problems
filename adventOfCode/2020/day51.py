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

lines = read_lines('input')
lines = remove_end_of_line(lines)
answer = 0
row_first = 0
row_last = 127
col_first = 0
col_last = 7
col = 0
row = 0
cols = []
rows = []

for line in lines:
    for character in line:
        if character == 'F':
            if row_last - row_first == 1:
                row = row_first
                continue
            row_last = row_last - round(((row_last - row_first) / 2))
        elif character == 'B':
            if row_last - row_first == 1:
                row = row_last
                continue
            row_first = row_first + round(((row_last - row_first) / 2))
        elif character == 'R':
            col_first = col_first + round(((col_last- col_first) / 2))
            if col_last - col_first == 1:
                col = col_last 
        elif character == 'L':
            if col_last - col_first == 1:
                col = col_first
            col_last = col_last - round(((col_last- col_first) / 2))
        else:
            print("problem")
    if (row * 8) + col > answer:
        answer = (row * 8) + col
    cols.append(col)
    rows.append(row)


    row_first = 0
    row_last = 127
    col_first = 0
    col_last = 7

previousX = -1
previousCounter = 127
myX = 127
myY = 127
for x in cols:
    if x != previousX and x != max(cols) and x != min(cols):
        if cols.count(x) < previousCounter:
            previousCounter = cols.count(x)
            myX = x
        previousX = x

previousCounter = 127
previousX = -1
for x in rows :
    if x != previousX and x != max(rows) and x != min(rows):
        if rows.count(x) < previousCounter:
            previousCounter = rows.count(x)
            myY = x
        previousX = x

print(answer)
print((myY*8)+myX)


