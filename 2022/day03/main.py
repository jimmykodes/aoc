from typing import List, Optional, Tuple, Set


def str_to_tup(line: str) -> Tuple[Set[str], Set[str]]:
    mid = len(line) // 2
    return (set(line[:mid]), set(line[mid:]))


def process_lines(lines: List[str]) -> List[Tuple[Set[str], Set[str]]]:
    out: List[Tuple[Set[str], Set[str]]] = []
    for line in lines:
        out.append(str_to_tup(line.strip()))
    return out


def test_data() -> List[Tuple[Set[str], Set[str]]]:
    data = [
        "vJrwpWtwJgWrhcsFMMfFFhFp",
        "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
        "PmmdzqPrVvPwwTWBwg",
        "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
        "ttgJtRGJQctTZtZT",
        "CrZsJsPPZsGzwwsLwLmpwMDw",
    ]
    return process_lines(data)


def get_data() -> List[Tuple[Set[str], Set[str]]]:
    with open("input.txt") as f:
        return process_lines(f.readlines())


def priority(item: int) -> int:
    if item - 64 <= 26:
        # capital letter
        return item - 64 + 26
    return item - 96


def p1() -> int:
    sum = 0
    for c1, c2 in get_data():
        item = ord(c1.intersection(c2).pop())
        sum += priority(item)
    return sum


def p2():
    data = get_data()
    sum = 0
    for i in range(2, len(data), 3):
        all: Optional[Set[str]] = None
        for j in range(3):
            elf = data[i - j]
            elf = elf[0].union(elf[1])
            if all is None:
                all = elf
            else:
                all = all.intersection(elf)
        if all is None:
            raise Exception("no item found")
        sum += priority(ord(all.pop()))
    return sum


if __name__ == "__main__":
    print(p1())
    print(p2())
