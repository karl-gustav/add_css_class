package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"unicode"
	"unicode/utf8"
)

var (
	classToAdd       string
	tag              string
	mustContain      string
	mustContainRegex *regexp.Regexp
	err              error
)

func main() {
	flag.StringVar(&classToAdd, "c", "", "Class you want to add to files")
	flag.StringVar(&tag, "t", "", "The html tag to add the class to.")
	flag.StringVar(&mustContain, "r", "", "Optional regex the tag must match within < and >.")
	flag.Parse()
	files := flag.Args()
	if classToAdd == "" {
		log.Fatalln("-c argument is required, use --help to get more info")
	}
	if tag == "" {
		log.Fatalln("-t argument is required, use --help to get more info")
	}
	if len(files) == 0 {
		log.Fatalln("You forgot to add files to update")
	}
	mustContainRegex, err = regexp.Compile(mustContain) //Empty regex always matches
	if err != nil {
		log.Fatalln("Your regex failed to compile with this error:", err)
	}
	for _, path := range files {
		fmt.Printf("%s: ", path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		content = fixTagClasses(content, tag, classToAdd)
		fmt.Println("OK")

		err = ioutil.WriteFile(path, content, 0)
		if err != nil {
			panic(err)
		}
	}
}

func fixTagClasses(content []byte, tagName, class string) []byte {
	var haveRead bytes.Buffer
	var tagBuf bytes.Buffer
	tagStart := []byte("<" + tagName)

	for _, char := range content {
		key, _ := utf8.DecodeRune([]byte{char})
		if tagBuf.Len() > 0 {
			tag := tagBuf.Bytes()
			lastTagChar := tag[len(tag)-1]
			if char != byte('>') && lastTagChar != byte('>') {
				tagBuf.WriteByte(char)
			} else {
				if mustContainRegex.Match(tag) {
					haveRead.Write(addClassToTag(tag, tagName, class))
				} else {
					haveRead.Write(tag)
				}
				haveRead.WriteByte(char)
				tagBuf.Reset()
			}
		} else if bytes.HasSuffix(haveRead.Bytes(), tagStart) && (unicode.IsSpace(key) || char == byte('>')) {
			haveRead.Truncate(haveRead.Len() - len(tagStart))
			tagBuf.Write(tagStart)
			tagBuf.WriteByte(char)
		} else {
			haveRead.WriteByte(char)
		}
	}
	return haveRead.Bytes()
}

var cssClassStart = regexp.MustCompile("class=([\"'])")

func addClassToTag(buf []byte, tagName, class string) []byte {
	var tagStart = regexp.MustCompile("<(" + tagName + ")")
	if cssClassStart.Match(buf) {
		buf = cssClassStart.ReplaceAll(buf, []byte("class=${1}"+class+" "))
	} else {
		buf = tagStart.ReplaceAll(buf, []byte(fmt.Sprintf("<${1} class=\"%s\"", class)))
	}

	return buf
}
