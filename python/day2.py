def collatz(num: int):
    arr = []
    arr.append(num)

    while num != 1:
        if num % 2 == 0:
            num //= 2
        else:
            num = (num * 3) + 1
        arr.append(num)
    
    return arr

print(collatz(11))