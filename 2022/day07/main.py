from dataclasses import dataclass
from typing import Dict, List, Optional


@dataclass
class File:
    name: str
    size: int

    def __repr__(self):
        return f"- {self.name} (file, size={self.size})"


class Dir:
    def __init__(self, name: str, level: int = 0, up: Optional["Dir"] = None):
        self.name = name
        self.files: List[File] = list()
        self.dirs: Dict[str, "Dir"] = dict()
        self.up: Optional["Dir"] = up
        self.level: int = level

    def add_file(self, name: str, size: int) -> None:
        self.files.append(File(name=name, size=size))
        self.files.sort(key=lambda x: x.size, reverse=True)

    def add_dir(self, name: str) -> None:
        self.dirs[name] = Dir(name, level=self.level + 1, up=self)

    @property
    def size(self) -> int:
        return sum(f.size for f in self.files) + sum(d.size for d in self.dirs.values())

    def __repr__(self) -> str:
        prefix = " " * self.level * 2
        out = f"{prefix}- {self.name} (dir, size={self.size})\n"
        for dir in self.dirs.values():
            out += dir.__repr__()
        for file in self.files:
            out += prefix + "  " + file.__repr__() + "\n"
        return out


class FS:
    def __init__(self):
        self.dir = Dir("/")
        self.cursor = self.dir

    def cd(self, to: str):
        if to == "..":
            if self.cursor.up is None:
                raise Exception("no dir to change to")
            self.cursor = self.cursor.up
        elif to[0] == "/":
            tos = [t for t in to.split("/") if t]
            for to in tos:
                self.cd(to)
        else:
            try:
                self.cursor = self.cursor.dirs[to]
            except KeyError:
                raise Exception(f"no directory named {to}")

    def ls(self, contents: List[str]):
        for line in contents:
            info, ident = line.split(" ")
            if info == "dir":
                self.cursor.add_dir(ident)
                continue
            self.cursor.add_file(ident, int(info))

    def __repr__(self) -> str:
        return self.dir.__repr__()


def _data(filename: str) -> List[str]:
    with open(filename) as f:
        return [line.strip() for line in f.readlines()]


def init_fs() -> FS:
    # lines = _data("test.txt")
    lines = _data("input.txt")
    fs = FS()
    i = 0
    while i < len(lines):
        line = lines[i]
        if line[0] == "$":
            line = line.split(" ")
            if line[1] == "cd":
                fs.cd(line[2])
            elif line[1] == "ls":
                next = lines[i + 1]
                contents = []
                while True:
                    contents.append(next)
                    i += 1
                    try:
                        next = lines[i + 1]
                    except IndexError:
                        break
                    if next[0] == "$":
                        break
                fs.ls(contents)
        i += 1
    return fs


def p1():
    fs = init_fs()
    stack = [fs.dir]
    dirs = []
    while len(stack) > 0:
        item = stack.pop(0)
        if item.size < 100000:
            dirs.append(item)
        for d in item.dirs.values():
            stack.append(d)

    return sum(d.size for d in dirs)


def p2():
    fs = init_fs()
    total_disk_size = 70000000
    disk_utilization = fs.dir.size
    required_space = 30000000
    free_space = total_disk_size - disk_utilization
    missing_space = required_space - free_space
    stack = [fs.dir]
    dirs = []
    while len(stack) > 0:
        item = stack.pop()
        if item.size > missing_space:
            dirs.append(item)
        for d in item.dirs.values():
            stack.append(d)
    dirs.sort(key=lambda x: x.size)
    return dirs[0].size


if __name__ == "__main__":
    print(p1())
    print(p2())
