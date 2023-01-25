def majority(arr: list[int or float]): # inputs can be a list of integers, floats or both
    print(f"Given array: {arr}")
    arr = sorted(arr) # sort arr in ascending order
    count = {} # create an empty dictionary to store the count of each element
    goal = len(arr)//2 + 1 # goal is the number of times an element must appear to be the majority element
    print("Goal: ", goal, " appearances")

    for num in arr: # loop through array
        if num in count: # if the element is already in the dictionary, increment the frequency
            count[num] += 1
        else: # if the element is not in the dictionary, add it and set the frequency to 1
            count[num] = 1
    
    for number, frequency in count.items(): # loop through all pairs in the dictionary
        if frequency >= goal:
            return repr(repr(number)) + " is the majority element, appearing " + str(frequency) + " times"



print(majority([3, 2, 3]))
print(majority([2, 2, 1, 1, 1, 2, 2]))
