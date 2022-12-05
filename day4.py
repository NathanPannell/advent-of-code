import re

def main():
    input = 'day4.txt'
    f = open(input, 'r')
    total = 0
    for line in [line.strip() for line in f]:
        x0, x1, x2, x3 = [int(i) for i in re.split(',|-', line)]
        total += [i for i in range(x0, x1 + 1) if i in range(x2, x3 + 1)] != [] 
    print(total)
    f.close()

if __name__ == '__main__':
    main()