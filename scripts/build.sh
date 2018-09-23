cp -R \
  tree-sitter-ruby/src/parser.c \
  tree-sitter-ruby/src/scanner.cc \
  tree-sitter-ruby/src/tree_sitter \
  .

go build
