def get_data() -> str:
    with open("input.txt") as f:
        return f.read().strip()


def detect_packet(data: str, length: int) -> int:
    for i in range(length, len(data)):
        if len(set(data[i - length : i])) == length:
            return i
    return -1


def p1() -> int:
    return detect_packet(get_data(), 4)


def p2() -> int:
    return detect_packet(get_data(), 14)


if __name__ == "__main__":
    print(p1())
    print(p2())
