function majority(arr) {
    console.log(`Given array: [${arr}]`)
    var count = {} // create an empty object to store the count of each element
    const goal = (Math.floor(arr.length/2) + 1) // goal is the number of times an element must appear to be the majority element
    console.log("Goal: " + goal + " appearances")

    for (num in arr) { // loop through array
        if (count.hasOwnProperty(arr[num])) { // if the element is already in the object, increment the frequency
            count[arr[num]] += 1
        } else { // if the element is not in the object, add it and set the frequency to 1
            count[arr[num]] = 1
        }
    }

    for (num in count) { // loop through all items in the object
        if (count[num] >= goal) {
            return `'${num}' is the majority element, appearing ${count[num]} times`
        }
    }
}

console.log(majority([3, 2, 3]))
console.log(majority([2,2,1,1,1,2,2]))