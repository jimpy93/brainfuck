package main

import(
  "io/ioutil"
  "fmt"
  "os"
)

func Interpret(prog []byte){
  ptrs := make([]byte, 10)
  cur := 0
  brackets := make([]int, 10)
  for i :=0; i<len(prog); {
    switch prog[i]{
    case '<':
      cur -= 1
      if cur < 0{
        panic("Data pointer is in negative!")
      }
    case '>':
      cur += 1
      if cur >= len(ptrs){
        ptrs = append(ptrs, 0)
      }
    case '.':
      fmt.Printf("%c", ptrs[cur])
    case '+':
      ptrs[cur] += 1
    case '-':
      ptrs[cur] -= 1
    case '[':
      if ptrs[cur] == 0{
        found := false
        bcount := 0
        for j, c := range prog[i:]{
          if c == '[' {
            bcount += 1
          } else if c == ']'{
            if bcount == 0 {
              i = j
              found = true
              break
            }
            bcount -= 1
          }
        }
        if !found {
          panic("Can't find end of loop!")
        }
      } else{
        brackets = append(brackets, i)
      }
    case ']':
      if ptrs[cur] != 0 {
        i = brackets[len(brackets)-1] - 1
      }
      brackets = brackets[:len(brackets)-1]
    }
    i += 1;
  }
}

func main() {
  argsWithoutProg := os.Args[1:]
  if len(argsWithoutProg) != 1{
    fmt.Println("Usage: brainfuck <bf file>")
    return
  }
  prog, err := ioutil.ReadFile(argsWithoutProg[0])
  if err != nil{
    fmt.Printf("Unable to read file.\n")
    panic(err)
  }

  Interpret(prog)
}
