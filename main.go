/* Owned and managed by L702x
* Licensed under AGPLv3
* Original repo:
*
* Note: this program is purley for jokes, it does not mess with your actual filesystem
* ASCII skull from: https://ascii.co.uk/art/skulls
*
* (c) 2026 L702x
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

func getBanner() string {
	// this mess because Golang hates me
	lines := []string{
		"                            ,--.",
		"                           {    }",
		"                           K,   }",
		"                          /  ~Y`",
		"                     ,   /   /",
		"                    {_'-K.__/",
		"                      `/-.__L._",
		"                      /  ' /`\\_}",
		"                     /  ' /",
		"             ____   /  ' /",
		"      ,-'~~~~    ~~/  ' /_",
		"    ,'             ``~~~  ',",
		"   (                        Y",
		"  {                         I",
		" {      -                    `,",
		" |       ',                   )",
		" |        |   ,..__      __. Y",
		" |    .,_./  Y ' / ^Y   J   )|",
		" \\           |' /   |   |   ||",
		"  \\          L_/    . _ (_,.'(",
		"   \\,   ,      ^^\"\"' / |      )",
		"     \\_  \\          /,L]     /",
		"       '-_~-,       ` `   ./`",
		"          `'{_            )",
		"              ^^\\..___,.--`",
	}
	return strings.Join(lines, "\n")
}

// why do we need this :/

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// bro who uses typewriters in 2026
func typewriter(c *color.Color, text string, delay time.Duration) {
	for _, char := range text {
		c.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	clearScreen() // why?

	// Colorz
	red := color.New(color.FgRed, color.Bold)
	white := color.New(color.FgWhite, color.Faint)
	yellow := color.New(color.FgYellow)
	cyan := color.New(color.FgCyan)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)

	blue.Println(getBanner())


	fmt.Println("\n")
	red.Println(" [!] FILES ENCRYPTED [!]")
	fmt.Println(" -----------------------------------------")
	
	typewriter(yellow, " All files have been locked... or were they?", 40*time.Millisecond)
	typewriter(yellow, " Hah you fell for it! No need to worry, this is just a joke... or is it? >:)", 40*time.Millisecond)
	
	fmt.Println("\n -----------------------------------------")


	for i := 5; i > 0; i-- {
		fmt.Printf("\r %s %d seconds...", cyan.Sprint("System wipe starting in:"), i)
		time.Sleep(1 * time.Second)
	}

	// say goodbye to pc >:)
	fmt.Println("\n")
	typewriter(red, " TOO LATE. THE WIPING PROCESS HAS BEGUN... or has it?", 20*time.Millisecond)
	typewriter(green, " HAHA you fell for it again! lmao, nah this is just a joke! have a nice day\n", 40*time.Millisecond)
	fmt.Println("Links n stuff:\n GitHub repo: \n (c) 2026 L702x\n")

	// the program closes after final msg so im adding this
	fmt.Println("\n")
	white.Print(" [ Press Enter to close this terminal ]")
	fmt.Scanln()
}