package magicapp

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"github.com/spf13/cobra"
	"github.com/365admin/koksmat-toolbelt/endpoints"
	"github.com/365admin/koksmat-toolbelt/cmds"

)

func RegisterCmds(){
	magicCmd := &cobra.Command{
	   Use:   "magic",
	   Short: "Magic Buttons",
	   Long:  `Describe the main purpose of this kitchen`,
   }
 
RootCmd.AddCommand(magicCmd)
	setupCmd := &cobra.Command{
	   Use:   "setup",
	   Short: "Setup",
	   Long:  `Describe the main purpose of this kitchen`,
   }
 
RootCmd.AddCommand(setupCmd)
	tasksCmd := &cobra.Command{
	   Use:   "tasks",
	   Short: "Tasks",
	   Long:  `Describe the main purpose of this kitchen`,
   }
 
RootCmd.AddCommand(tasksCmd)
	provisionCmd := &cobra.Command{
	   Use:   "provision",
	   Short: "Provision",
	   Long:  `Describe the main purpose of this kitchen`,
   }
 
RootCmd.AddCommand(provisionCmd)
	decommissionCmd := &cobra.Command{
	   Use:   "decommission",
	   Short: "Decommision",
	   Long:  `Describe the main purpose of this kitchen`,
   }
 
RootCmd.AddCommand(decommissionCmd)
}
