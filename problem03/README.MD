# Problem Statement
Given a string containing characters (a-z), implement a function runLengthEncode that
compresses repeated ‘runs’ of the same character by storing the length of that run, and
provide a function runLengthDecode to reverse the compression. The output can be
anything, as long as you can recreate the input with it. <br>
For example:  <br>
```go
runLengthEncode("aaaaaaaaaabbbaxxxxyyyzyx") // returns "a10b3a1x4y3z1y1x1" 
runLengthDecode("a10b3a1x4y3z1y1x1")       // returns "aaaaaaaaaabbbaxxxxyyyzyx" 
```

# brute force solution
- start by iterating through the chars one by one 
    - for each char we start an inner loop to count all similar chars next to our char

so its simply a 2 for loop with time complexity = O(n*n) because we have a lot of repeated work as following : 
```go
    inp := "aaabbb"
    ➜ i = 0 , counter = 1
        ➜  j = 1
            counter = 2
        ➜ j = 2
            counter = 3
        ➜ j = 3 
            [X] stop, and break after adding the counter next to char at index [i] into our output string builder 
    ➜ i = 1 
       and here is the repeated work !!!  
```

# Sliding Window Pattern 
Since in the brute force solution I did a lot of repeated work, what if I can finish a letter with all its repeats and then start from the next letter, and thats what makes me start thinking of Sliding window pattern 