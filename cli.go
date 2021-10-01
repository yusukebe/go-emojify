package emojify

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kyokomi/emoji"
	"golang.org/x/crypto/ssh/terminal"
)

type CLI struct {
}

func (cli *CLI) Run() {
	var text string
	if terminal.IsTerminal(0) {
		if len(os.Args) <= 1 {
			cli.showUsage()
			os.Exit(0)
		}
		text = os.Args[1]
		text = text + "\n"
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		text = string(b)
	}
	emoji.Print(text)
}

func (cli *CLI) showUsage() {
	fmt.Println("Usage: $ go-emojify \"I love :ramen:\"")
}
