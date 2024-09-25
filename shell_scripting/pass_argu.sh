#! /bin/bash


# simple way to take argument 
# $0 is for shell script name and rest is for arguments passed to it
#echo $1 $2 $3 ' > Enter the values'

# Take arguments and store into array
args=($@)

# Simple way to print is ${args[1]},${args[2]}

# good way to print all at once 
echo $@  # in this way args[0] is first argument and not the script file name

# to print number of args passed 
echo $#