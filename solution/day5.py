import re

def main():
    input = '../input/day5.txt'
    f = open(input, 'r')
    all_stacks = [[] for i in range(9)]
    for row in range(8):
        line = f.readline()
        [all_stacks[i].insert(0, line[4*i + 1]) for i in range(9) if line[4*i + 1] != ' ']
    # [print(stack) for stack in all_stacks]

    f.readline()
    f.readline()

    # Get movement instructions
    for line in f:
        x0, num, x1, start, x2, end = line.split(' ')
        move(int(num), all_stacks[int(start)-1], all_stacks[int(end)-1])

    [print(stack) for stack in all_stacks]
    for stack in all_stacks:
        print(stack.pop(), end="")
    f.close()

def move(number, start, end):
    temp = []
    for i in range(number):
        temp.append(start.pop())
    for i in range(number):
        end.append(temp.pop())

if __name__ == '__main__':
    main()