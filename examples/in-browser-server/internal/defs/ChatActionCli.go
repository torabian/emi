//go:build !wasm

package defs

import (
	"bufio"
	"context"
	"github.com/urfave/cli/v3"
	"os"
)

// ChatActionCliFlags returns the query-parameter flags the
// ChatAction action can bind from urfave v3.
//
func ChatActionCliFlags() []cli.Flag {
	flags := []cli.Flag{}
	return flags
}

// ChatActionCliReactiveHandler builds a full *cli.Command for the
// ChatAction reactive action: stdin becomes the read side (one frame per
// line), and whatever the factory's returned channel produces is written straight to
// stdout - so the generated command composes as one leg of a Linux pipe
// (`producer | app the-action | consumer`). Piping ends the command the same way an
// EOF on a socket would: stdin closing (or scanner error) ends the read loop and the
// command returns.
func ChatActionCliReactiveHandler(
	factory func(session ChatActionSession) (chan []byte, error),
) *cli.Command {
	meta := ChatActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ChatActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		done := make(chan bool)
		read := make(chan ChatActionReadChan)
		session := ChatActionSession{
			Ctx:  c,
			Done: done,
			Read: read,
		}
		out, err := factory(session)
		if err != nil {
			return err
		}
		go func() {
			for {
				select {
				case <-done:
					return
				case data, more := <-out:
					os.Stdout.Write(data)
					if !more {
						return
					}
				}
			}
		}()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := make([]byte, len(scanner.Bytes()))
			copy(line, scanner.Bytes())
			read <- ChatActionReadChan{Data: line}
		}
		return scanner.Err()
	}
	return cmd
}

// ChatActionCliReactive is a high-level convenience wrapper around
// ChatActionCliReactiveHandler. It registers the generated command as a
// subcommand of an existing urfave v3 *cli.Command, the same way ChatActionGin
// registers a route on a Gin engine.
func ChatActionCliReactive(
	app *cli.Command,
	factory func(session ChatActionSession) (chan []byte, error),
) {
	app.Commands = append(app.Commands, ChatActionCliReactiveHandler(factory))
}
