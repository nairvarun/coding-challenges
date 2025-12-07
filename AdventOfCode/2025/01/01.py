from typing import TextIO

# part 1
def part1(moves: TextIO) -> None:
    n = 50
    res = 0
    for move in moves:
        direction, turns = move[0], int(move[1:])
        if direction == "R":
            n = (n + turns) % 100
        else:
            n = (n - turns) % 100
        if n == 0:
            res += 1
    print(res)

# part 2
def part2(moves: TextIO) -> None:
    n = 50
    res = 0
    for move in moves:
        direction, turns = move[0], int(move[1:])
        res += turns // 100
        turns %= 100
        if direction == "R":
            if n + turns >= 100:
                res += 1
            n = (n + turns) % 100
        else:
            if n > 0 and turns >= n:
                res += 1
            n = (n - turns) % 100
    print(res)

def main():
    with open("./input.txt", "r") as f:
        # part1(f)
        part2(f)

main()