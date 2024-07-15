package utils

import "os"

func CheckDir(path string)bool{
  info,err := os.Stat(path)
  if err != nil {
    return false
  }
  if info.IsDir() == true{
    return true
  }
  return false
}
