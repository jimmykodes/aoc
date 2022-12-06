from dataclasses import dataclass
from typing import List, Tuple


@dataclass
class Direction:
    n: int
    src: int
    dest: int

    @classmethod
    def from_line(cls, line: str) -> "Direction":
        parts = line.split(" ")
        n = int(parts[1])
        src = int(parts[3])
        dest = int(parts[5])
        return cls(n, src, dest)


@dataclass
class Crate:
    id: str


class Stack:
    def __init__(self, *args):
        self.stack = []
        for a in args:
            self.add(Crate(id=a))

    def add(self, c: Crate):
        self.stack.append(c)

    def pop(self) -> Crate:
        c = self.stack[-1]
        self.stack = self.stack[:-1]
        return c

    def move(self, n: int, to: "Stack", preserve_order=False):
        if preserve_order:
            intermediate = Stack()
            for _ in range(n):
                intermediate.add(self.pop())
            for _ in range(n):
                to.add(intermediate.pop())
        else:
            for _ in range(n):
                to.add(self.pop())


def _data(filename: str) -> Tuple[List[Stack], List[Direction]]:
    with open(filename) as f:
        data = [line.rstrip() for line in f.readlines()]

    seen_break = False
    stack_lines = []
    moves = []
    for line in data:
        if line == "":
            seen_break = True
            continue
        if seen_break:
            moves.append(Direction.from_line(line))
        else:
            stack_lines.append(line)
    stack_lines.reverse()
    stacks: List[Stack] = []
    for i, line in enumerate(stack_lines):
        if i == 0:
            for j in range(1, len(line), 4):
                stacks.append(Stack())
        else:
            k = 0
            for j in range(1, len(line), 4):
                crate_id = line[j]
                if crate_id != " ":
                    stacks[k].add(Crate(crate_id))
                k += 1
    return (stacks, moves)


def get_data(test=False):
    if test:
        return _data("test.txt")
    return _data("input.txt")


def p1() -> str:
    stacks, dirs = get_data()

    for d in dirs:
        stacks[d.src - 1].move(d.n, stacks[d.dest - 1])

    out = ""
    for stack in stacks:
        if stack is not None:
            out += stack.pop().id
    return out


def p2() -> str:
    stacks, dirs = get_data()

    for d in dirs:
        stacks[d.src - 1].move(d.n, stacks[d.dest - 1], preserve_order=True)

    out = ""
    for stack in stacks:
        if stack is not None:
            out += stack.pop().id
    return out


if __name__ == "__main__":
    print(p1())
    print(p2())
