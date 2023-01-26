function collatz(num) {
    arr = []
    arr.push(num)

    while (num != 1) {
        if (num % 2 == 0) {
            num = Math.floor(num / 2)
        } else {
            num = (num * 3) + 1
        }
        arr.push(num)
    }

    return arr
}

console.log(`[${collatz(11).toString()}]`)