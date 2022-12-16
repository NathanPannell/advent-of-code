input = '../input/day1.txt'
arr = [0]
input_stream = open(input, 'r')
for line in [line.strip() for line in input_stream]:
    if line:
        arr[0] += int(line)
    else:
        arr.insert(0, 0)
print(sorted(arr))
input_stream.close()