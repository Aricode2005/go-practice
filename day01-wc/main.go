package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Counts struct {
	Lines, Words, Bytes int
}
type Options struct {
	Lines, Words, Bytes bool
}

func Count(r io.Reader) (Counts, error) {
	var c Counts
	br := bufio.NewReader(r)
	inWord := false
	for {
		b, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c, err
		}
		c.Bytes++
		if b == '\n' {
			c.Lines++
		}

		switch b {
		case ' ', '\n', '\t', '\r', '\v', '\f':
			inWord = false
		default:
			if !inWord {
				c.Words++
				inWord = true
			}
		}

	}
	return c, nil
}
func countFile(path string) (Counts, error) {
	f, err := os.Open(path)
	if err != nil {
		return Counts{}, fmt.Errorf("open %s:%w", path, err)

	}
	defer f.Close()
	return Count(f)
}

func printCounts(c Counts, name string, opt Options) {
	if opt.Lines {
		fmt.Printf("%8d", c.Lines)
	}
	if opt.Words {
		fmt.Printf("%8d", c.Words)
	}
	if opt.Bytes {
		fmt.Printf("%8d", c.Bytes)
	}
	if name != "" {
		fmt.Printf(" %s", name)
	}
	fmt.Println()
}

func (c *Counts) Add(o Counts) {
	c.Lines += o.Lines
	c.Words += o.Lines
	c.Bytes += o.Bytes
}

func main() {
	showLines := flag.Bool("l", false, "count lines")
	showWords := flag.Bool("w", false, "count words")
	showBytes := flag.Bool("c", false, "count bytes")
	flag.Parse()
	opt := Options{Lines: *showLines, Words: *showWords, Bytes: *showBytes}
	if !opt.Lines && !opt.Words && !opt.Bytes {
		opt = Options{Lines: true, Words: true, Bytes: true}
	}
	args := flag.Args()
	if len(args) == 0 {
		c, err := Count(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "gowc:", err)
			os.Exit(1)
			return
		}
		printCounts(c, "", opt)
		return
	}
	failed := false
	var total Counts
	for _, path := range args {
		c, err := countFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "gowc:", err)
			failed = true
			continue

		}
		printCounts(c, path, opt)
		total.Add(c)
	}
	if len(args) > 1 {
		printCounts(total, "total", opt)
	}
	if failed {
		os.Exit(1)
	}
}

// func main() {
// 	c, err := Count(strings.NewReader("hello world\nsecond line\n"))
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}
// 	fmt.Printf("%d %d %d\n", c.Lines, c.Words, c.Bytes)
// }
