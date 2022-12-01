def get_data():
    with open("input.txt") as f:
        return f.readlines()


def p1():
    max = 0
    elf = 0
    for line in get_data():
        if (l := line.strip()) != "":
            elf += int(l)
            continue
        if elf > max:
            max = elf
        elf = 0
    return max


def p2():
    stack = []
    elf = 0
    for line in get_data():
        if (l := line.strip()) != "":
            elf += int(l)
            continue
        stack.append(elf)
        stack.sort()
        elf = 0
        if len(stack) > 3:
            stack = stack[1:]
    return sum(stack)


if __name__ == "__main__":
    print(p1())
    print(p2())
