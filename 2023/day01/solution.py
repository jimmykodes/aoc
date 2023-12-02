from dataclasses import dataclass
import re
from typing import List


def p1(lines: List[str]):
    total = 0
    for line in lines:
        digits = re.findall(r"\d", line)
        num = digits[0] + digits[-1]
        total += int("".join(num))
    return total


lookup = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


@dataclass
class Result:
    index: int
    value: str


def p2(lines: List[str]):
    total = 0
    for line in lines:
        res = []
        for v in [*lookup.keys(), *lookup.values()]:
            try:
                for match in re.finditer(v, line):
                    v = lookup[v] if v in lookup else v
                    res.append(Result(match.start(), v))
            except ValueError:
                continue
        res.sort(key=lambda x: x.index)
        fd = res[0]
        ld = res[-1]
        total += int(f"{fd.value}{ld.value}")
    return total


if __name__ == '__main__':
    with open('assets/input.txt') as f:
        # with open('assets/test2.txt') as f:
        lines = f.readlines()
    print(p2(lines))
