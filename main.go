package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// blast format
//qseqid sseqid staxids bitscore score length qlen' -query $f -num_threads 40 > $f.VsNR.out"
// contig00187     gi|1005469232|ref|XP_015771527.1|       70779   872     2252    507     2206
//kreport format needed
// readID	seqID	taxID	score	2ndBestScore	hitLength	queryLength	numMatches

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("readID	seqID	taxID	score	2ndBestScore	hitLength	queryLength	numMatches")

	for scanner.Scan() {
		//values := []string{}
		//fmt.Println(scanner.Text())
		s := strings.Split(scanner.Text(), "\t")
		//readID, seqID, taxID, bitscore, score, hitLength, queryLen  := s[0], s[1], s[2] ,s[3], s[4], s[5], s[6]


		//fmt.Println(readID, seqID, taxID, bitscore,score, hitLength, queryLen, 1)
		s = append(s,"1")
		result:= strings.Join(s, "\t")
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


