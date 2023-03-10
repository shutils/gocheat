package common

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
)

func GetAppDirName() (name string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	name = path.Join(home, ".gocheat")
	return
}

func MakeAppDir(name string) error {
	err := os.Mkdir(name, 0777)
	return err
}

func AppDirIsExist(name string) (r bool) {
	if _, err := os.Stat(name); err != nil {
		r = false
	} else {
		r = true
	}
	return
}

func GetFileNames(appDirName string) (fns []string) {
	fs, err := ioutil.ReadDir(appDirName)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range fs {
		fns = append(fns, v.Name())
	}
	return
}

func GetText(filename string) (t string) {
	tb, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	t = string(tb)
	return
}

func EditFile(filename string) error {
	c := exec.Command("nvim", filename)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	err := c.Run()
	return err
}

type Index struct {
  Text string
  Index int
}

func GetIndex(filename string) (indexes []Index) {
  appDirName := GetAppDirName()
  fp, err := os.Open(filepath.Join(appDirName, filename))
  if err != nil {
    log.Fatalln(err)
  }
  defer fp.Close()

  s := bufio.NewScanner(fp)

  var ta []string
  r := regexp.MustCompile(`<.+>`)
  for s.Scan() {
    ta = append(ta, s.Text())
  }
  for i, v := range ta {
    if r.MatchString(v) {
      indexes = append(indexes, Index{Text: v, Index: i})
    }
  }
  return
}
