# Advent of code

Storing all of advent of code answers in this repository.

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

# Day 1
docker run ghcr.io/lucasclerissepro/aoc-2022-1:latest 

# Day 2
docker run ghcr.io/lucasclerissepro/aoc-2022-2:latest

# Day 3
docker run ghcr.io/lucasclerissepro/aoc-2022-3:latest

...

```
