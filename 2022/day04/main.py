from dataclasses import dataclass
from typing import List, Tuple


@dataclass
class Assignement:
    start: int
    end: int

    @classmethod
    def from_text(cls, text: str) -> "Assignement":
        a, b = text.split("-")
        a = int(a)
        b = int(b)
        start = min([a, b])
        end = max([a, b])

        return cls(start=start, end=end)

    def overlap(self, other: "Assignement") -> bool:
        return (
            other.start <= self.start <= other.end
            or self.start <= other.start <= self.end
        )


def _data(filename: str) -> List[Tuple[Assignement, Assignement]]:
    with open(filename) as f:
        lines = f.readlines()
    out = []
    for line in lines:
        first, second = line.split(",")
        out.append((Assignement.from_text(first), Assignement.from_text(second)))
    return out


def get_data():
    return _data("input.txt")
    # return _data("test.txt")


def p1() -> int:
    data = get_data()
    total = 0
    for p1, p2 in data:
        if (p1.start <= p2.start and p1.end >= p2.end) or (
            p2.start <= p1.start and p2.end >= p1.end
        ):
            total += 1
    return total


def p2() -> int:
    data = get_data()
    total = 0
    for p1, p2 in get_data():
        total += 1 if p1.overlap(p2) else 0
    return total


if __name__ == "__main__":
    print(p1())
    print(p2())
