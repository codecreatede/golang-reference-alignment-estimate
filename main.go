package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

/*

 Author Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-28

 This package deals with the genome or the gene alignment and you can define a reference sequence for the same.
 By defining the reference sequence, you are using that sequence against the other sequence for the calculation
 of the alignment estimates.

*/

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	alignment string
	sequence  string
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This is used to estimate the alignment according to the specified sequence",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&alignment, "alignment", "a", "alignment file to be used for the estimate", "alignment file")
	rootCmd.Flags().
		StringVarP(&sequence, "refsequence", "s", "sequence to be used as reference", "reference sequence")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	type alignmentIDStore struct {
		id string
	}

	type alignmentSeqStore struct {
		seq string
	}

	type withoutRef struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignmentID := []alignmentIDStore{}
	alignmentSeq := []alignmentSeqStore{}
	sequenceSpec := []string{}
	referenceSeq := sequence
	refID := []string{}
	refSeq := []string{}
	outRef := []withoutRef{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentID = append(alignmentID, alignmentIDStore{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeq = append(alignmentSeq, alignmentSeqStore{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceSpec = append(sequenceSpec, string(line))
		}
	}

	for i := range alignmentID {
		if alignmentID[i].id == sequence {
			continue
		} else {
			outRef = append(outRef, withoutRef{
				id:  string(alignmentID[i].id),
				seq: string(alignmentSeq[i].seq),
			})
		}
	}

	for i := range alignmentID {
		if alignmentID[i].id == sequence {
			refID = append(refID, alignmentID[i].id)
			refSeq = append(refSeq, alignmentSeq[i].seq)
		}
	}

	refMinus := []string{}

	for i := range outRef {
		refMinus = append(refMinus, outRef[i].seq)
	}

	counterAT := 0
	counterAG := 0
	counterAC := 0

	for i := 0; i < len(refMinus); i++ {
		for j := 0; j < len(refMinus[0]); j++ {
			for k := 0; k < len(referenceSeq); k++ {
				if string(refMinus[i][j]) == "A" && string(referenceSeq[k]) == "T" {
					counterAT++
				}
				if string(refMinus[i][j]) == "A" && string(referenceSeq[k]) == "C" {
					counterAG++
				}
				if string(refMinus[i][j]) == "A" && string(referenceSeq[k]) == "G" {
					counterAC++
				}
			}
		}
	}

	counterTG := 0
	counterTC := 0
	counterTA := 0

	for i := 0; i < len(refMinus)-1; i++ {
		for j := 0; j < len(refMinus[0]); j++ {
			for k := 0; k < len(referenceSeq); k++ {
				if string(refMinus[i][j]) == "T" && string(referenceSeq[k]) == "G" {
					counterTA++
				}
				if string(refMinus[i][j]) == "T" && string(referenceSeq[k]) == "C" {
					counterTC++
				}
				if string(refMinus[i][j]) == "T" && string(referenceSeq[k]) == "A" {
					counterTA++
				}
			}
		}
	}
	counterGC := 0
	counterGA := 0
	counterGT := 0

	for i := 0; i < len(refMinus); i++ {
		for j := 0; j < len(refMinus[0]); j++ {
			for k := 0; k < len(referenceSeq); k++ {
				if string(refMinus[i][j]) == "G" && string(referenceSeq[k]) == "C" {
					counterGC++
				}
				if string(refMinus[i][j]) == "G" && string(referenceSeq[k]) == "A" {
					counterGA++
				}
				if string(refMinus[i][j]) == "G" && string(referenceSeq[k]) == "T" {
					counterGT++
				}
			}
		}
	}
	counterCA := 0
	counterCT := 0
	counterCG := 0

	for i := 0; i < len(refMinus)-1; i++ {
		for j := 0; j < len(refMinus[0]); j++ {
			for k := 0; k < len(referenceSeq); k++ {
				if string(refMinus[i][j]) == "C" && string(referenceSeq[k]) == "A" {
					counterCA++
				}
				if string(refMinus[i][j]) == "C" && string(referenceSeq[k]) == "T" {
					counterCT++
				}
				if string(refMinus[i][j]) == "C" && string(referenceSeq[k]) == "G" {
					counterCG++
				}
			}
		}
	}
	fmt.Println(
		"The collinearity block for A as a base pattern and T as a mismatch is %d",
		counterAT,
	)
	fmt.Println("The collinearity block for A as a base pattern G as a mismatch is %d", counterAG)
	fmt.Println(
		"The collinearity block for A as a base pattern and C as a mismatch is %d",
		counterAC,
	)
	fmt.Println(
		"The collinearity block for T as a base pattern and G as a mismatch is %d",
		counterTG,
	)
	fmt.Println("The collinearity block for T as a base pattern C as a mismatch is %d", counterTC)
	fmt.Println(
		"The collinearity block for T as a base pattern and A as a mismatch is %d",
		counterTA,
	)
	fmt.Println(
		"The collinearity block for G as a base pattern and C as a mismatch is %d",
		counterGC,
	)
	fmt.Println("The collinearity block for G as a base pattern A as a mismatch is %d", counterGA)
	fmt.Println(
		"The collinearity block for G as a base pattern and T as a mismatch is %d",
		counterGT,
	)
	fmt.Println(
		"The collinearity block for C as a base pattern and A as a mismatch is %d",
		counterCA,
	)
	fmt.Println("The collinearity block for C as a base pattern T as a mismatch is %d", counterCT)
	fmt.Println(
		"The collinearity block for C as a base pattern and G as a mismatch is %d",
		counterCG,
	)
}
