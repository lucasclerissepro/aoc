# Advent of code

Storing all of advent of code answers in this repository. 

> The code stored in this repository does not reflect the real submissions 
> i'm doing everyday and is refactored to be faster and more readable.

## Getting started

This repository uses [Earthly](https://earthly.dev) to build and run exercises.

You can run the solutions for all of 2022 questions by doing:

```
cd 2022
earthly +run
```

Or run a specific day solution by doing:

```
cd 2022/1
earthly +run
```


## Trying out

There's docker images built for every day so you can run it:

```shell

# Day 1 part one
docker run -v $(pwd)/data/input.txt:/aoc/data/input.txt ghcr.io/lucasclerissepro/aoc-2022-1-one:latest 

# Day 1 part two
docker run -v $(pwd)/data/input.txt:/aoc/data/input.txt ghcr.io/lucasclerissepro/aoc-2022-1-two:latest 

...

```

You can mount any input in the container at the path `/aoc/data/input.txt`
