# Problem 02
Write the Linux command needed to change a file name from original.txt to
changed.txt

# Solution 
To test the command run the following 
```shell
make test 
```
I created a Makefile for you so it can automate the task by deleting the old renamed file, creating a new file with the old name, renaming it to the new name 
```Makefile
test: 
# i used the -f so incase you don't have a file with this name before running the makefile for the first time 
	rm -f changed.txt 
	touch original.txt
	mv original.txt changed.txt
```