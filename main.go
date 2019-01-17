package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gregdel/pushover"
	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
)

var (
	appToken string
	userKey  string
	stdin    *os.File
	debug    bool
)

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-pushover-handler",
		Short: "The Sensu Go handler plugin for pushover",
		RunE:  run,
	}

	/*
	   Security Sensitive flags
	     - default to using envvar value
	     - do not mark as required
	     - manually test for empty value
	*/

	cmd.Flags().StringVarP(&appToken,
		"app.token",
		"a",
		os.Getenv("PUSHOVER_APP_TOKEN"),
		"Pushover v1 API app token, use default from PUSHOVER_APP_TOKEN env var")

	cmd.Flags().StringVarP(&userKey,
		"user.key",
		"u",
		os.Getenv("PUSHOVER_USER_KEY"),
		"Pushover v1 API user key, use default from PUSHOVER_USER_KEY env var")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}

	if appToken == "" {
		_ = cmd.Help()
		return fmt.Errorf("app token is empty")
	}

	if userKey == "" {
		_ = cmd.Help()
		return fmt.Errorf("user key is empty")
	}

	eventJSON, err := ioutil.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %s", err)
	}

	event := &types.Event{}
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal stdin data: %s", err)
	}

	if err = event.Validate(); err != nil {
		return fmt.Errorf("failed to validate event: %s", err)
	}

	if !event.HasCheck() {
		return fmt.Errorf("event does not contain check")
	}

	return notifyPushover(event)
}

func notifyPushover(event *types.Event) error {
	app := pushover.New(appToken)
	recipient := pushover.NewRecipient(userKey)
	message := &pushover.Message{
		Title:   fmt.Sprintf("%s/%s", event.Entity.Name, event.Check.Name),
		Message: event.Check.Output,
	}
	response, err := app.SendMessage(message, recipient)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Pushover message sent for event %s/%s\n", event.Entity.Name, event.Check.Name)

	if debug == true {
		log.Println(response)
	}

	return nil
}
