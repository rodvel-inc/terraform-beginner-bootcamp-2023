// Declare the package "main". It has a special meaning in Go.
// It is where the execution of the program starts.
package main

// imports the fmt package, which contains functions for formatted I/O.
import (
	//"log"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//Defines the "main" function. This is the entry popint for the application.
// When you run the program, it starts the execution from this point.
func main () {
	//Format.PrintLine and prints to bash standard output.
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
	fmt.Println("Hello World!")
}

func Provider() *schema.Provider {
	var p* schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{

		},
		DataSourcesMap: map[string]*schema.Resource{
			
		},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for the external service",
			},
			"token": {
				Type: schema.TypeString,
				Sensitive: true,
				Required: true,
				Description: "The bearer token for authorization",
			},
			"user_uuid": {
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
			//	ValidateFunc: validateUUID
			},
			
		},
	}
//	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

//func validateUUID(v interface{}, k string) (ws []string, errors []error) {
//	log.Print('validateUUID:start')
//	value := v.(string)
//	if _,err = uuid.Parse(value); err != nil {
//		errors = apppend(error, fmt.Errorf("Invalid UUID format"))
//	}
//	log.Print('validateUUID:end')
//}
