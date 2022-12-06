import re

def main():
    myfile = open('../input/day6.txt', 'r')
    line = myfile.readline().strip()
    a, b, c, d, e, f, g, h, i, j, k, l, m, n = ['', '', '', '', '', '', '', '', '', '', '', '', '', '']
    
    counter = 0
    for char in line:
        a, b, c, d, e, f, g, h, i, j, k, l, m, n = b, c, d, e, f, g, h, i, j, k, l, m, n, char
        counter += 1
        if sorted(set([a, b, c, d, e, f, g, h, i, j, k, l, m, n])) == sorted([a, b, c, d, e, f, g, h, i, j, k, l, m, n]) and counter >= 14:
            print(counter)
            break

    myfile.close()

if __name__ == '__main__':
    main()