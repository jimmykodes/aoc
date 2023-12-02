from typing import Dict, List
from dataclasses import dataclass


@dataclass
class Group:
    num: int
    color: str


@dataclass
class Round:
    groups: Dict[str, int]


@dataclass
class Game:
    id: int
    rounds: List[Round]


def lines(file: str) -> List[Game]:
    games = []
    with open(file) as f:
        for line in f.readlines():
            line = line.strip()
            game, rounds = line.split(": ")
            id = int(game.split(" ")[1])
            game = Game(id, list())
            rounds = rounds.split("; ")
            for round in rounds:
                groups = round.split(", ")
                groups = [group.split(" ") for group in groups]
                groups = {g[1]: int(g[0]) for g in groups}
                game.rounds.append(Round(groups))
            games.append(game)
    return games


def p1(games: List[Game]):
    total = 0
    max_red = 12
    max_green = 13
    max_blue = 14
    for game in games:
        possible = True
        for round in game.rounds:
            if round.groups.get("red", 0) > max_red:
                possible = False
                break
            if round.groups.get("blue", 0) > max_blue:
                possible = False
                break
            if round.groups.get("green", 0) > max_green:
                possible = False
                break
        if possible:
            total += game.id
    return total


l = lines("assets/input.txt")
print(p1(l))
