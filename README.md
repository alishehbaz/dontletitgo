### Rate Limiting library in Go

#### I plan to implement all the main rate limiting algorithms in Go

##### Starting with Token Bucket

To Run, make sure you have go installed and then type this in the terminal

> go run main.go

### TODOS

TODO: Implement multithreading \
TODO: Fix the issue with consuming the left over token, i.e when numOfTokens requested > current tokens left, so it ends up not using any token at all \
TODO: Allow multiple time modes like seconds, minuts, hours, milliseconds
