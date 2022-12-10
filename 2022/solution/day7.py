import Node

root = Node.Node('/')
total = 0
min_space = 8381165

def main():
    myfile = open('../input/day7.txt', 'r')
    lines = [line.strip().split(' ') for line in myfile][1:]
    
    current = root
    for token in lines:
        if token[0] == '$' and token[1] == 'cd':
            current = change_directory(current, token[2])
        elif token[0] == 'dir':
            print('adding dir', token[1], 'as child of', current.name)
            new_node = Node.Node(token[1])
            new_node.parent = current
            current.add_child(new_node)
        elif token[0].isnumeric():
            print('adding file', token[1], 'as child of', current.name)
            new_node = Node.Node(token[1], int(token[0]))
            new_node.parent = current
            current.add_child(new_node)

    total = get_sum_of_small_directories(root)
    print(total)
    print(find_smallest_space(root))


    myfile.close()
    

def change_directory(current, target):
    if target == '/':
        print('returning to root', 'from', current.name)
        return root
    elif target == '..':
        print('returning parent of', current.name)
        return current.parent if current.parent != None else root
    elif target in [child.name for child in current.children]:
        print('changing to dir', target, 'from', current.name)
        return current.get_child(target)
    else:
        print('creating new dir', target, 'from', current.name)
        new_node = Node.Node(target)
        new_node.set_parent(current)
        return new_node

def get_sum_of_small_directories(current):
    # Calculate size of current directory
    # this_dir = current.get_size() if current.get_size() <= 100000 else 0

    # # Calculate sum of children
    # this_children = sum([get_sum_of_small_directories(child) for child in current.children if current.has_children()])

    # return this_dir + this_children
    addition = 0
    # print(current.name, [child.name for child in current.children])
    if len(current.children) > 0 and 0 < current.get_size() <= 100000:
        # print(current.name, [child.name for child in current.children], current.get_size())
        addition = current.get_size()
    
    return sum([get_sum_of_small_directories(child) for child in current.children]) + addition

def find_smallest_space(current):
    space = [find_smallest_space(child) for child in current.children]
    space.append(current.get_size())
    possible_values = [size for size in space if size >= min_space]

    return min(possible_values) if len(possible_values) > 0 else 0
    
def print_text(input: text):
    [print(char) for char in input]
    


if __name__ == '__main__':
    main()