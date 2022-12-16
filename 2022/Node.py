class Node:

    def __init__(self, name='', size=0):
        self.name = name
        self.size = size
        self.parent = None
        self.children = []

    def get_name(self):
        return self.name

    def set_parent(self, parent):
        self.parent = parent

    def get_parent(self):
        return self.parent

    def add_child(self, child):
        self.children.append(child)

    def get_children(self):
        return self.children

    def get_child(self, name):
        return [child for child in self.children if child.name == name][0]

    def get_size(self):
        if not self.has_children():
            return self.size
        else:
            return sum([child.get_size() for child in self.children])

    def has_children(self):
        return len(self.children) > 0