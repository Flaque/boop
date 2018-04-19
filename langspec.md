```ruby
# ints
3
4
5

# Doubles
2.0
4.0

# strings
hello
there
./foo.txt        # one string
cat of the hat   # 4 single strings
"cat of the hat" # one single string

` # this line is ignored
multiline string
goes here
` # So is this line

# booleans
true
false

# Math
3   plus  3
2.0 sub  1
48  mult 2
23  div  0.3
2   pow  2
10  mod  2

# Variables
:a = hello
:b = world

# Conditionals
2 == 2 # true
1 < 2
1 > 2
1 <= 2
1 >= 2
:a or :b
:a and :b

# If statements
if :jazz < 2 do

end

# For loops
for :isWorking do

end

# of statement maybe?
for :pup of :puppies do
    echo :pup
end

# Functions
func :get do
  return 3
end

func :add :a :b do
  return :a plus :b # Return statements go to a virtual stdout
end

# Implicit arguments with reserved variable names passed into every function
:in
:flags
:args  # global os arguments

# Function calls
:sum = :get # If it's a funciton, we'll evaluate it
:sum = :add 2 3

# Embedding function calls
:add (:get) 3 # Must use () parens

# :echo is boop's print
:echo hello world

# Unix Pipes
:sum | :echo
:add 4 5 | :echo

# All tools in $PATH are treated as functions
:echo hello there
:grep go -r ./
:ls

# Lists
:puppies = [snuggle fluffer "mr cuddles" pupperino] # Spaces = commas
:products = [
    [price name count]
    [23.0 "hair dryer" 4]
    [2 snickers 80]
] # Embedded lists

:puppies = :append :puppies spot # adding to a list

# Accessing lists
:puppies 1    # returns "fluffer"
:products 0 0 # returns "price"

# Structures
:ledger = {
    :employees = [george bob]
    :subledger = {
        :name = foo
        :employees = [cat dog]
    }
}

:ledger employees 0 # returns george
:ledger subledger name

# Exporting
export :ledger # Exports it as a cmd line tool or as an importable function somwehere else

# Importing
import ./foo/jazz # Auto imports :jazz
import ./foo/jazz as :j
```
