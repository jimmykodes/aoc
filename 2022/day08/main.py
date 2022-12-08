from typing import List, Optional, Tuple


class Tree:
    def __init__(self, height: int):
        self.height = height
        self.left: Optional["Tree"] = None
        self.right: Optional["Tree"] = None
        self.top: Optional["Tree"] = None
        self.bottom: Optional["Tree"] = None

    def add_left(self, left: "Tree"):
        self.left = left
        left.right = self

    def add_top(self, top: "Tree"):
        self.top = top
        top.bottom = self

    def dir_visible(self, direction: str) -> Tuple[bool, int]:
        visible = True
        tree = getattr(self, direction)
        viewing_distance = 0
        while tree is not None:
            viewing_distance += 1
            if tree.height >= self.height:
                visible = False
                break
            tree = getattr(tree, direction)
        return (visible, viewing_distance)

    @property
    def visible(self) -> bool:
        return (
            self.dir_visible("left")[0]
            or self.dir_visible("top")[0]
            or self.dir_visible("right")[0]
            or self.dir_visible("bottom")[0]
        )

    @property
    def view_distance(self) -> int:
        return (
            self.dir_visible("left")[1]
            * self.dir_visible("top")[1]
            * self.dir_visible("right")[1]
            * self.dir_visible("bottom")[1]
        )


def _data(filename: str) -> List[str]:
    with open(filename) as f:
        return [line.strip() for line in f.readlines()]


def get_data() -> List[List[Tree]]:
    data = _data("input.txt")
    trees: List[List[Tree]] = []
    for y in range(len(data)):
        row: List[Tree] = []
        for x in range(len(data[0])):
            t = Tree(int(data[y][x]))
            if x > 0:
                t.add_left(row[-1])
            if y > 0:
                t.add_top(trees[-1][x])
            row.append(t)
        trees.append(row)
    return trees


def p1() -> int:
    data = get_data()
    count = 0
    for row in data:
        for tree in row:
            if tree.visible:
                count += 1
    return count


def p2() -> int:
    data = get_data()
    max = 0
    for row in data:
        for tree in row:
            if (vd := tree.view_distance) > max:
                max = vd
    return max


if __name__ == "__main__":
    print(p1())
    print(p2())
