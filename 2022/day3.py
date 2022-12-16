import math

CAPITAL_OFFSET = 64
LOWERCASE_OFFSET = 96

def main():
    input = '../input/day3.txt'
    f = open(input, 'r')
    lines = [line.strip() for line in f]
    groups = [[lines[3*i], lines[3*i+1], lines[3*i+2]] for i in range(int(len(lines)/3))]
    results = [[char for char in group[0] if char in group[1] and char in group[2]] for group in groups]
    print(sum([value(i[0]) for i in results]))
    f.close()

def value(char):
    return ord(char) - CAPITAL_OFFSET + 26 if char.isupper() else ord(char) - LOWERCASE_OFFSET

if __name__ == '__main__':
    main()