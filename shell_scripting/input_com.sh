#! /bin/bash

# this is about taking user's input using read command

#echo "Enter the name: "
#read name
#echo hi  $name , nice meeting you

# input on the same line
#read -p 'Enter username :' username  # -p for same prompt


# to take password as input (don't show while typing)
#read -sp 'Enter sudo password :' password  # -s stands for silent
#echo
#echo 'hi bhadve' $username
#echo "pass is" $password

# take multiple input and store in array
#echo "Enter values : "
#read -a names    # -a for array
#echo "values :" ${names[0]},${names[1]}


# when you don't pass any value to read command , it is
# automatically stored inside SYSTEM var called REPLY
echo "Enter values : "
read     # -a for array
echo "values : $REPLY"
