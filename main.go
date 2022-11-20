package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "crypto/md5"
  "encoding/hex"
)

func main() {

  if len(os.Args) < 2{
    fmt.Println("Usage: %s <hash_to_crack>", os.Args[0])
    os.Exit(0)
  }

  hash := os.Args[1]
  fmt.Printf("Hash = %s\n",hash)

  f, err := os.Open("password_list.txt")
  if err != nil{
    log.Fatalln(err)
  }

  fileScanner := bufio.NewScanner(f)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan(){
    data := string(fileScanner.Text())
    digest := calculateMD5(data)
    if digest == hash{
      fmt.Printf("Successfully cracked %s\n", hash)
      fmt.Printf("The password is %s\n", data)
    }
  }

  f.Close()
}


func calculateMD5(text string) string {
  hash := md5.Sum([]byte(text))
  return hex.EncodeToString(hash[:])
}
