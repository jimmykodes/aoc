from typing import List, Tuple


def test_data() -> List[str]:
    return ["A Y", "B X", "C Z"]


def get_data() -> List[str]:
    with open("input.txt") as f:
        return [line.strip() for line in f.readlines()]


# A - ROCK
# B - PAPER
# C - SCISSORS

# X - ROCK
# Y - PAPER
# Z - SCISSORS

ROCK = 1
PAPER = 2
SCISSORS = 3
WIN = 6
TIE = 3
LOSS = 0


def p1() -> int:
    OPTIONS = {
        ROCK: {
            ROCK: TIE,
            PAPER: WIN,
            SCISSORS: LOSS,
        },
        PAPER: {
            ROCK: LOSS,
            PAPER: TIE,
            SCISSORS: WIN,
        },
        SCISSORS: {
            ROCK: WIN,
            PAPER: LOSS,
            SCISSORS: TIE,
        },
    }

    def round(p1: int, p2: int) -> int:
        return OPTIONS[p1][p2] + p2

    lookup_table = {
        "A X": round(ROCK, ROCK),
        "A Y": round(ROCK, PAPER),
        "A Z": round(ROCK, SCISSORS),
        "B X": round(PAPER, ROCK),
        "B Y": round(PAPER, PAPER),
        "B Z": round(PAPER, SCISSORS),
        "C X": round(SCISSORS, ROCK),
        "C Y": round(SCISSORS, PAPER),
        "C Z": round(SCISSORS, SCISSORS),
    }
    total = 0
    for line in get_data():
        total += lookup_table[line]
    return total


def p2() -> int:
    OPTIONS = {
        ROCK: {
            WIN: PAPER,
            TIE: ROCK,
            LOSS: SCISSORS,
        },
        PAPER: {
            WIN: SCISSORS,
            TIE: PAPER,
            LOSS: ROCK,
        },
        SCISSORS: {
            WIN: ROCK,
            TIE: SCISSORS,
            LOSS: PAPER,
        },
    }

    def round(p1: int, p2: int) -> int:
        return OPTIONS[p1][p2] + p2

    lookup_table = {
        "A X": round(ROCK, LOSS),
        "A Y": round(ROCK, TIE),
        "A Z": round(ROCK, WIN),
        "B X": round(PAPER, LOSS),
        "B Y": round(PAPER, TIE),
        "B Z": round(PAPER, WIN),
        "C X": round(SCISSORS, LOSS),
        "C Y": round(SCISSORS, TIE),
        "C Z": round(SCISSORS, WIN),
    }
    total = 0
    for line in get_data():
        total += lookup_table[line]
    return total


if __name__ == "__main__":
    print(p1())
    print(p2())
