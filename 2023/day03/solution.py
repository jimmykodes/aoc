from typing import List

from random import random


def lines(filename: str) -> List[str]:
    with open(filename) as f:
        l = f.readlines()
        return l


def p1(lines: List[str]):
    total = 0
    for y, line in enumerate(lines):
        x = 0
        num_start = -1
        num_end = -1
        while x < len(line):
            char = line[x]
            if char in "0123456789":
                if num_start == -1:
                    num_start = x
            elif num_start != -1:
                # not a digit but we've accumulated a number
                num_end = x  # exclusive
                num = line[num_start:num_end]
                recorded = False
                for _x in range(num_start-1, num_end+1):
                    if _x < 0 or _x >= len(line):
                        # trying to look before first or after last col
                        continue
                    for dy in [-1, 0, 1]:
                        if dy == -1 and y == 0:
                            # first line, don't look up
                            continue
                        if dy == 1 and y == len(lines)-1:
                            # last line, don't look down
                            continue
                        _y = y + dy
                        _char = lines[_y][_x]
                        if _char not in "0123456789.\n":
                            if random() < 0.2:
                                print(num, "-", f"({_x}, {_y}):", _char)
                            if not recorded:
                                # we're adjecent to a symbol so capture the number
                                total += int(num)
                                recorded = True

                num_start = -1
                num_end = -1
            x += 1
    return total


def p2(lines: List[str]):
    total = 0
    gears = {}
    for y, line in enumerate(lines):
        x = 0
        num_start = -1
        num_end = -1
        while x < len(line):
            char = line[x]
            if char in "0123456789":
                if num_start == -1:
                    num_start = x
            elif num_start != -1:
                # not a digit but we've accumulated a number
                num_end = x  # exclusive
                num = line[num_start:num_end]
                recorded = False
                for _x in range(num_start-1, num_end+1):
                    if _x < 0 or _x >= len(line):
                        # trying to look before first or after last col
                        continue
                    for dy in [-1, 0, 1]:
                        if dy == -1 and y == 0:
                            # first line, don't look up
                            continue
                        if dy == 1 and y == len(lines)-1:
                            # last line, don't look down
                            continue
                        _y = y + dy
                        _char = lines[_y][_x]
                        if _char == "*":
                            if not recorded:
                                # we're adjecent to a gear symbol so capture the number
                                if _y not in gears:
                                    gears[_y] = {}
                                if _x not in gears[_y]:
                                    gears[_y][_x] = []
                                gears[_y][_x].append(int(num))
                                recorded = True
                num_start = -1
                num_end = -1
            x += 1
    for v1 in gears.values():
        for parts in v1.values():
            if len(parts) == 2:
                total += parts[0] * parts[1]
    return total


print(p2(lines("assets/input.txt")))
