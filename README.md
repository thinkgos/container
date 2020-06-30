# container

container implements some containers, currently the containers are not thread-safe. `safe` path support thread-safe.

[![GoDoc](https://godoc.org/github.com/thinkgos/container?status.svg)](https://godoc.org/github.com/thinkgos/container)
[![Build Status](https://travis-ci.org/thinkgos/container.svg?branch=master)](https://travis-ci.org/thinkgos/container)
[![codecov](https://codecov.io/gh/thinkgos/container/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/container)
![Action Status](https://github.com/thinkgos/container/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/container)](https://goreportcard.com/report/github.com/thinkgos/container)
[![Licence](https://img.shields.io/github/license/thinkgos/container)](https://raw.githubusercontent.com/thinkgos/container/master/LICENSE)  

- **[How to use this repo](#how-to-use-this-package)**
- **[Containers](#Containers-Interface)**
  - [Sets](#sets) support`interface{}` which implement Comparator interface and builtin type.
  - [Stack](#stack) 
    - stack use container/list.
    - quick stack use builtin slice.
  - [Queue](#queue) use container/list
  - [PriorityQueue](#priorityqueue) use builtin slice with container/heap
  - [LinkedList](#linkedlist) use container/list
  - [ArrayList](#arraylist) use builtin slice.
  - [LinkedMap](#linkedMap) use container/list and builtin map.
- **[safe container](#safe-container)**
  - [fifo](#fifo) 
    > FIFO solves this use case:
    > * You want to process every object (exactly) once.
    > * You want to process the most recent version of the object when you process it.
    > * You do not want to process deleted objects, they should be removed from the queue.
    > * You do not want to periodically reprocess objects.

  - [heap](#heap) Heap is a thread-safe producer/consumer queue that implements a heap data structure.It can be used to implement priority queues and similar data structures.
- **[others](#others)**
  - [Comparator](#Comparator) 
    - [Sort](#sort) sort with Comparator interface
    - [Heap](#heap) heap with Comparator interface