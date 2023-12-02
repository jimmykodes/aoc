from typing import List
from dataclasses import dataclass


@dataclass
class Group:
    num: int
    color: str


@dataclass
class Round:
    groups: List[Group]


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
                groups = [Group(int(g[0]), g[1]) for g in groups]
                game.rounds.append(Round(groups))
            games.append(game)
    return games


print(lines("assets/test.txt"))
