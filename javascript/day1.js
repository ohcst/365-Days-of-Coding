function majority(arr) {
    var count = {}
    const goal = (Math.floor(arr.length/2) + 1)
    console.log("Goal: " + goal + " appearances")

    for (num in arr) {
        if (count.hasOwnProperty(arr[num])) {
            count[arr[num]] += 1
        } else {
            count[arr[num]] = 1
        }
    }

    for (const num in count) {
        if (count[num] >= goal) {
            return `'${num}' is the majority element, appearing ${count[num]} times`
        }
    }
}

console.log(majority([3, 2, 3]))
console.log(majority([2,2,1,1,1,2,2]))