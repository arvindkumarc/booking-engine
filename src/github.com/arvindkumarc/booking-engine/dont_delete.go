package dummy

import (
"fmt"
"github.com/revel/cmd/revel"
)

func TestMe() {
  var cmdRun = &Command{
	UsageLine: "run [import path] [run mode] [port]",
	Short: "run a Revel application",
	Long: `
	Run the Revel web application named by the given import path.

	For example, to run the chat room sample application:

	revel run github.com/revel/revel/samples/chat dev

	The run mode is used to select which set of app.conf configuration should
	apply and may be used to determine logic in the application itself.

	Run mode defaults to "dev".

	You can set a port as an optional third parameter. For example:

	revel run github.com/revel/revel/samples/chat prod 8080`,
	}
	fmt.Println(cmdRun)
}
