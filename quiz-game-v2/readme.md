This is the version 2 of quiz game, in this part we have converted our blocking code to non-blocking by using go routines and channels, also we have used labels to break out of the for loop. Compare it with the part 1 and you will see the difference both run perfectly fine. 
In order to run:  
`go run quizgame.go -h`  
`go run quizgame.go -limit=<time in seconds>`  
Example:  
`go run quizgame.go -limit=3`