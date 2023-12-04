from dataclasses import dataclass
from typing import List


@dataclass
class Card:
    id: int
    wins: int


def cards(filename: str) -> List[Card]:
    out = []
    with open(filename) as f:
        for line in f.readlines():
            line = line.strip()
            c, line = line.split(": ")
            winning, have = line.split(" | ")
            num = int(c.split()[1])
            winning = {int(x) for x in winning.split()}
            have = {int(x) for x in have.split()}
            out.append(Card(num, len(have.intersection(winning))))
    return out


def p1(cards: List[Card]) -> int:
    return sum([2**(c.wins-1) for c in cards if c.wins > 0])


def p2(cards: List[Card]) -> int:
    total = 0
    for c in cards[::-1]:
        total += wins(c, cards) + 1
    return total


def cached(func):
    cache = {}

    def wrapper(card: Card, cards: List[Card]):
        if card.id in cache:
            print(f"cache hit: {card.id}")
            return cache[card.id]
        print(f"cache miss: {card.id}")
        res = func(card, cards)
        cache[card.id] = res
        return res

    return wrapper


@cached
def wins(card: Card, cards: List[Card]):
    total = card.wins
    for i in range(card.wins):
        if card.id+i >= len(cards):
            break
        total += wins(cards[card.id+i], cards)
    return total


print(p2(cards("assets/input.txt")))
