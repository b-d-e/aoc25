let read_lines filename =
  let ic = open_in filename in
  let rec loop acc =
    try
      let line = input_line ic in
      loop (line :: acc)
    with End_of_file -> close_in ic; List.rev acc
  in
  loop []

let mod100 x = ((x mod 100) + 100) mod 100

let part1 lines =
  let pos = ref 50 in
  let cnt = ref 0 in
  List.iter (fun ln ->
      let s = String.trim ln in
      if String.length s > 0 then
        let dir = s.[0] in
        let dist_str = String.sub s 1 (String.length s - 1) in
        try
          let dist = int_of_string dist_str in
          (match dir with
           | 'L' -> pos := mod100 (!pos - dist)
           | 'R' -> pos := mod100 (!pos + dist)
           | _ -> ());
          if !pos = 0 then incr cnt
        with _ -> ()
    ) lines;
  !cnt

let hits_in_rotation pos dir dist =
  match dir with
  | 'R' ->
      let offset = (100 - (pos mod 100)) mod 100 in
      let offset = if offset = 0 then 100 else offset in
      if dist >= offset then 1 + (dist - offset) / 100 else 0
  | 'L' ->
      let offset = pos mod 100 in
      let offset = if offset = 0 then 100 else offset in
      if dist >= offset then 1 + (dist - offset) / 100 else 0
  | _ -> 0


let part2 lines =
  let pos = ref 50 in
  let cnt = ref 0 in
  List.iter (fun ln ->
      let s = String.trim ln in
      if String.length s > 0 then
        let dir = s.[0] in
        let dist_str = String.sub s 1 (String.length s - 1) in
        try
          let dist = int_of_string dist_str in
          cnt := !cnt + hits_in_rotation !pos dir dist;
          (match dir with
           | 'L' -> pos := mod100 (!pos - dist)
           | 'R' -> pos := mod100 (!pos + dist)
           | _ -> ());
        with _ -> ()
    ) lines;
  !cnt

  
let () =
  let filename = if Array.length Sys.argv > 1 then Sys.argv.(1) else "input.txt" in
  let lines = try read_lines filename with _ -> [] in
  Printf.printf "Part 1: %d\n" (part1 lines);
  Printf.printf "Part 2: %d\n" (part2 lines)
