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

"
multiline string
goes here
"

# booleans
true
false

# Math
3   +  3
2.0 -  1
48  * 2
23  /  0.3

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

## Embedded Library

Boop has a number of functions that come with the language by default. They're written in go. These are functions that you can't overwrite and don't need to import. 

```ruby
# OS info
:exit
:in     # stdin
:out    # stdout (TODO: might not be included)
:flags  # passed flags to the script (not the function)
:args   # args passed to the script 

# Math functions
:pow 2 3  # 8
:sqrt 4   # 2 
:mod 12 7 # 5
```

## Standard Library

The standard library is similar, though it's written in boop, not in go.