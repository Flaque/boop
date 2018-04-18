# Boop

This is a programming language experiment. Take everything with a grain of salt, nothing's working right now and I have no plans at this time to support this language.

## Overview

Boop is a sane scripting language for... scripting.

It's main claim to fame is every command in your $PATH is a function in Boop by default.

In other words, all standard unix commands like `cd`, `grep`, `ls` all work by default. All commands you've `brew install`'d or `apt-get`'d work too.

As such, hello world in Boop is just:

```
echo Hello World
```

You can pass any number of arguments (seperated by a space) into a function:

```
echo My a e s t h e t i c is emoji punk
```

Pipes are also a thing in Boop and they work like they do in Unix.

```
ls | grep *.boop | echo
```

## stdout in Boop

Unlike bash, stdout is not the default for commands. If you want the results of a command, you'll need to `echo` it.

For example, the following will print nothing to stdout.

```
grep hello file.txt
```

However, add an `echo` and we're all good.

```
grep hello file.txt | echo
```

In Boop, it's preferable to have commands speak only when allowed.

## Variables

Like bash, variables are marked with a `$` such as:

```
echo $dogs
```

Unlike bash, you need to use a `$` even when assigning the variable:

```
$dogs = "puppers"
```

## Functions

A function that doesn't take any arguments and doesn't return anything looks like this:

```
func stuff do
  # something in here
end
```

A function with arguments looks like this:

```
func stuff $a $b $c do
  echo $a
  echo $b
  echo $c
end
```

You can return something like this:

```
func stuff do
  return "hello"
end
```

You call your functions like you do bash commands:

```
stuff 1 2 3
```

## If statements

If statements look like this:

```
if $thing do
    # stuff
end
```

You can also just ignore the `do` block and return something:

```
if $blah return "cats"
```

Else statements don't exist are discouraged to avoid large, nested if else blocks.

Rather, prefer a "guard" pattern.

```
func gradeTest $score do
  if $score < 50 return "fail"

  return "pass"
end
```
