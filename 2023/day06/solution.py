from typing import List, NamedTuple, Tuple


def races(fn: str, combine=False) -> List[Tuple[int, int]]:
    with open(fn) as f:
        lines = f.readlines()
    times = lines[0].split()[1:]
    distances = lines[1].split()[1:]
    if combine:
        times = [int("".join(times))]
        distances = [int("".join(distances))]
    else:
        times = [int(t) for t in times]
        distances = [int(d) for d in distances]
    return list(zip(times, distances))


def p1(races: List[Tuple[int, int]]) -> int:
    total = 1  # init at one since this is a product not a sum
    for race in races:
        total_time = race[0]
        best_dist = race[1]
        hold_time = total_time//2
        ways = 0
        while get_dist(total_time, hold_time) > best_dist:
            hold_time -= 1
            ways += 1
        hold_time = (total_time//2) + 1
        while get_dist(total_time, hold_time) > best_dist:
            hold_time += 1
            ways += 1
        total *= ways
    return total


def find_upper_bound(total_time, target_dist):
    l = 0
    r = total_time
    while l < r:
        m = ((r - l)//2)+l
        if get_dist(total_time, m) > target_dist:
            if get_dist(total_time, m+1) <= target_dist:
                # peek to see if this is the boundary
                return m
            l = m+1
        else:
            if get_dist(total_time, m-1) > target_dist:
                # peek to see if this is the boundary
                return m-1
            r = m-1
    return r


def find_lower_bound(total_time, target_dist):
    l = 0
    r = total_time
    while l < r:
        m = ((r - l)//2)+l
        if get_dist(total_time, m) > target_dist:
            if get_dist(total_time, m-1) <= target_dist:
                # peek to see if this is the boundary
                return m
            r = m-1
        else:
            if get_dist(total_time, m+1) > target_dist:
                # peek to see if this is the boundary
                return m+1
            l = m+1
    return l


def p1_binary(races: List[Tuple[int, int]]) -> int:
    total = 1  # init at one since this is a product not a sum
    for race in races:
        total_time = race[0]
        best_dist = race[1]
        lb = find_lower_bound(total_time, best_dist)
        rb = find_upper_bound(total_time, best_dist)
        ways = (rb+1)-lb
        total *= ways
    return total


def get_dist(total_time, hold_time):
    speed = hold_time
    run_time = total_time - hold_time
    return run_time * speed


def main():
    r = races("assets/input.txt")
    print(p1_binary(r))
    r = races("assets/input.txt", combine=True)
    print(p1(r))
    print(p1_binary(r))


main()
