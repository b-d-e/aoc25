OCaml solution for Advent of Code Day 1 (Secret Entrance)

Files
- `day_01.ml` - standalone OCaml implementation (reads `input.txt` by default or file passed as first argument).

Build & run

Compile with OCaml compiler:

    ocamlc -o day_01 day_01.ml
    ./day_01 input.txt

Or run with the toplevel (slower):

    ocaml day_01.ml input.txt

If you use dune, create a small dune project and add `day_01.ml` as the executable source.
