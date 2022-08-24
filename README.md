# few
fetch-execute-write instruction pipeline using go channels

A program that uses channels to simulate CPU instruction pipeline (3 stage, fetch,execute/write). 
First go routine, generates few random instructions and sends them over a channel to a second routine which interprets and executes the instructions and sends results over a second channel to a third routine which displays results.
