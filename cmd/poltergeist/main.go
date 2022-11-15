package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	inputFile        = flag.String("file", "", "Input file")
	flagStdin        = flag.Bool("stdin", false, "Read from stdin")
	flagEncode       = flag.Bool("encode", false, "Encode whitespace")
	flagDecode       = flag.Bool("decode", false, "Decode whitespace")
	flagTest         = flag.Bool("test", false, "print whitespace table")
	fileBytes        []byte
	characterMapping = map[rune]rune{
		'a': '\u0009', // CHARACTER TABULATION
		'b': '\u000A', // LINE FEED (LF)
		'c': '\u000B', // LINE TABULATION
		'd': '\u000C', // FORM FEED (FF)
		'e': '\u000D', // CARRIAGE RETURN (CR)
		'f': '\u0020', // SPACE
		'g': '\u0085', // NEXT LINE (NEL)
		'h': '\u00A0', // NO-BREAK SPACE
		'i': '\u1680', // OGHAM SPACE MARK
		'j': '\u2000', // EN QUAD
		'k': '\u2001', // EM QUAD
		'l': '\u2002', // EN SPACE
		'm': '\u2003', // EM SPACE
		'n': '\u2004', // THREE-PER-EM SPACE
		'o': '\u2005', // FOUR-PER-EM SPACE
		'p': '\u2006', // SIX-PER-EM SPACE
		'q': '\u2007', // FIGURE SPACE
		'r': '\u2008', // PUNCTUATION SPACE
		's': '\u2009', // THIN SPACE
		't': '\u200A', // HAIR SPACE
		'u': '\u2028', // LINE SEPARATOR
		'v': '\u2029', // PARAGRAPH SEPARATOR
		'w': '\u202F', // NARROW NO-BREAK SPACE
		'x': '\u205F', // MEDIUM MATHEMATICAL SPACE
		'y': '\u3000', // IDEOGRAPHIC SPACE
		'z': '\uFEFF', // ZERO WIDTH NO-BREAK SPACE
	}
	characterMappingHex = map[string]rune{
		"a": '\u0009', // CHARACTER TABULATION
		"b": '\u000A', // LINE FEED (LF)
		"c": '\u000B', // LINE TABULATION
		"d": '\u000C', // FORM FEED (FF)
		"e": '\u000D', // CARRIAGE RETURN (CR)
		"f": '\u0020', // SPACE
		"0": '\u0085', // NEXT LINE (NEL)
		"1": '\u00A0', // NO-BREAK SPACE
		// "2": '\u1680', // OGHAM SPACE MARK
		"2": '\u2000', // EN QUAD
		"3": '\u2001', // EM QUAD
		"4": '\u2002', // EN SPACE
		"5": '\u2003', // EM SPACE
		"6": '\u2004', // THREE-PER-EM SPACE
		"7": '\u2005', // FOUR-PER-EM SPACE
		"8": '\u2006', // SIX-PER-EM SPACE
		"9": '\u2007', // FIGURE SPACE
		// "a":'\u2008', // PUNCTUATION SPACE
		// "a":'\u2009', // THIN SPACE
		// "a":'\u200A', // HAIR SPACE
		// "a":'\u2028', // LINE SEPARATOR
		// "a":'\u2029', // PARAGRAPH SEPARATOR
		// "a":'\u202F', // NARROW NO-BREAK SPACE
		// "a":'\u205F', // MEDIUM MATHEMATICAL SPACE
		// "a":'\u3000', // IDEOGRAPHIC SPACE
		// "a":'\uFEFF', // ZERO WIDTH NO-BREAK SPACE
	}
)

func init() {
	flag.Parse()
	// check that only one of encode or decode is set
	if *flagEncode && *flagDecode {
		fmt.Println("Only one of encode or decode can be set")
		os.Exit(1)
	}
	if !*flagEncode && !*flagDecode {
		fmt.Println("One of encode or decode must be set")
		os.Exit(1)
	}

	// check that input file is set or stdin is set
	if *inputFile == "" && !*flagStdin {
		fmt.Println("One of input file or stdin must be set")
		os.Exit(1)
	}

	if *flagStdin {
		// read stdin
		var err error
		fileBytes, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if *inputFile == "" {
			fmt.Println("Input file path must be set")
			os.Exit(1)
		} else {
			//  read inputFile
			openFile, err := os.Open(*inputFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer openFile.Close()
			// process bytes
			fileBytes, err = io.ReadAll((openFile))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func main() {
	if *flagTest {
		printAllWhiteSpace()
		os.Exit(0)
	}
	if *flagEncode {
		encodeWhiteSpaceHex()
	}
	if *flagDecode {
		decodeWhiteSpaceHex()
	}

}
func encodeWhiteSpaceHex() {
	// convert fileBytes to hex string
	hexFile := hex.EncodeToString(fileBytes)
	// fmt.Println(hexFile)
	// convert hexfile to whitespace
	for _, char := range hexFile {
		fmt.Printf("%c", characterMappingHex[string(char)])
	}
}

func decodeWhiteSpaceHex() {
	// convert fileBytes to hex string
	fileString := string(fileBytes)
	// read unicode values as hex
	var hexString string
	for _, char := range fileString {
		// use lookup table to convert to hex
		for key, value := range characterMappingHex {
			if char == value {
				hexString += key
			}
		}
	}

	// fmt.Println(hexString)
	// convert hex to bytes
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(decodedBytes))

}

func printAllWhiteSpace() {
	for key, value := range characterMapping {
		fmt.Printf("%c: %c\n", key, value)
	}
}
