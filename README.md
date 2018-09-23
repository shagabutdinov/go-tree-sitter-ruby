[![Build Status](https://travis-ci.com/shagabutdinov/go-tree-sitter-ruby.svg?branch=master)](https://travis-ci.com/shagabutdinov/go-tree-sitter-ruby)

Go Tree Sitter Ruby
===================

Go bindings for [tree-sitter](https://github.com/tree-sitter/tree-sitter)
[ruby](https://github.com/tree-sitter/tree-sitter-ruby) on top of [go-tree-sitter](https://github.com/smacker/go-tree-sitter).


Installation
============

```
$ go get github.com/shagabutdinov/go-tree-sitter-ruby
$ cd $GOPATH/src/github.com/shagabutdinov/go-tree-sitter-ruby
$ ./scripts/build.sh
# build may take a significant time because ruby grammar has
# 800+k lines of C code
```


Usage
=====

```
package main

import (
  "github.com/shagabutdinov/go-tree-sitter-ruby"
  "github.com/smacker/go-tree-sitter"
  "log"
)

func main() {
  parser := sitter.NewParser()
  parser.SetLanguage(ruby.GetLanguage())

  sourceCode := []byte("variable = 'value'")
  tree := parser.Parse(sourceCode)
  defer tree.Delete()

  log.Println(tree.RootNode().String())
}
```
