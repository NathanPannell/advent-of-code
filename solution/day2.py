KEY = {
'A': 1, 'B': 2, 'C': 3, 
'X': 1, 'Y': 2, 'Z': 3
}

def main():
    input = '../input/day2.txt'
    f = open(input, 'r')
    total = 0
    for line in f:
        x0, x1 = line.strip().split(' ')
        y0, y1 = KEY[x0], KEY[x1]
        total += (y0+y1)%3+1 + 3*(y1-1)
    print(total)
    f.close()

if __name__ == '__main__':
    main()