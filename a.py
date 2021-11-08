import random
import math
# sys just for case stack too much
import sys
sys.setrecursionlimit(10**9)


def printBoard(board, src, dst, discovered):
    for col in range(42):
        print('X', end='')
    print()
    for row in range(20):
        print('X', end='')
        for col in range(40):
            if src and src[0] == row:
                if src[1] == col:
                    print('S', end='')
                    continue
            if str([row, col]) in discovered.keys():
                print('-', end='')
                continue

            if dst and dst[0] == row:
                if dst[1] == col:
                    print('D', end='')
                    continue
            if str(col) in board[row]:
                print('X', end='')
            else:
                print(' ', end='')
        print('X')
    for col in range(42):
        print('X', end='')
    print()


def randomSrcAndDst(board):
    src = [random.randint(0, 19)]
    colSrc = random.randint(0, 39)
    while str(colSrc) in board[src[0]]:
        colSrc = random.randint(0, 39)
    src += [colSrc]
    return src


def euclidDistance(curState, dst):  # heuristic
    return math.sqrt(pow(dst[0] - curState[0], 2) + pow(dst[1] - curState[1], 2))


def Up(src):
    newMove = src[0] - 1
    if newMove >= 0 and str(src[1]) not in board[newMove]:
        return [newMove, src[1]]
    return None


def Down(src):
    newMove = src[0] + 1
    if newMove < 20 and str(src[1]) not in board[newMove]:
        return [newMove, src[1]]
    return None


def Left(src):
    newMove = src[1] - 1
    if newMove >= 0 and str(newMove) not in board[src[0]]:
        return [src[0], newMove]
    return None


def Right(src):
    newMove = src[1] + 1
    if newMove < 40 and str(newMove) not in board[src[0]]:
        return [src[0], newMove]
    return None


def secret(e):  # key to sorted
    return e[1]


def solve(board, dst, states, discovered):
    # condition 1 to stop, when stack is EMPTY
    if states == []:
        printBoard(board, [], dst, discovered)
        print("Can't go to Goal")
        exit()
    node = states.pop()
    src = node[0]
    g = node[2]
    # condition 2 to stop, when we go to GOAL
    if src == dst:
        printBoard(board, src, dst, discovered)
        print("Goal !!!!!!!!!!!")
        exit()
    # update g(x)
    g += 1
    newMove = [x for x in [Up(src), Left(src), Down(src), Right(src)]
               if x and str(x) not in discovered.keys()]
    # update discovered
    for ele in newMove:
        discovered[str(ele)] = 1
    gFunc = [g for x in newMove]
    # f(x) = g(x) + h(x) with h(x) is euclid distance
    newDistance = [euclidDistance(ele, dst) + g for ele in newMove]
    Pair = zip(newMove, newDistance, gFunc)
    for x in Pair:
        states += [x]
    # update states
    states = sorted(states, key=secret)
    solve(board, dst, states[::-1], discovered)


if __name__ == "__main__":
    # read map, map should be 2x1 size ratio
    if(len(sys.argv) < 2):
        print(
            "USAGE: python3 pathFinding.py mapX.txt \nWith X = [1,3]")
        exit()
    f = open(str(sys.argv[1]), "r")
    board = []
    for line in f:
        part = line.split()
        board += [part]
    # u can self pick src and des with syntax "src = [10,15]"
    src = randomSrcAndDst(board)
    dst = randomSrcAndDst(board)
    # init discovered, that will save the path we moved
    discovered = {
        str(src): 1
    }
    printBoard(board, src, dst, discovered)
    print("\n")
    # init states, this is a stack will save all path can move ,with top of stack is the best move
    states = [(src, euclidDistance(src, dst), 0)]
    solve(board, dst, states,  discovered)