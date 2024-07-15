package utils

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

var w *zip.Writer

func ZipDir(path string)error{
  if CheckDir(path) == false {
    return errors.New("The path is not a directory")
  }
  f,err := os.Create("./out/output.zip")

  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  w = zip.NewWriter(f)

  defer w.Close()
  
  err = filepath.Walk(path,walker)
  if err != nil {
    log.Fatal(err)
  }
  return nil
  
}


func walker(path string,info os.FileInfo,err error)error{
  if err != nil {
    return err
  }
  file,err := os.Open(path)
  if info.IsDir() {
    return nil
  }
  if err != nil {
    return err
  }

  defer file.Close()

  f,err := w.Create(path)
  if err != nil {
    return err
  }
  io.Copy(f,file)
  return nil
}

