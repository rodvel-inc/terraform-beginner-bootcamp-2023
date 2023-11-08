// Declare the package "main". It has a special meaning in Go.
// It is where the execution of the program starts.
package main

// imports the fmt package, which contains functions for formatted I/O.
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

type Config struct {
	Endpoint string
	Token string
	UserUuid string
}

func Provider() *schema.Provider {
	var p* schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": Resource(),
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
				ValidateFunc: validateUUID,
			},
			
		},
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
	value := v.(string)
	if _, err := uuid.Parse(value); err != nil {
		errors = append(errors, fmt.Errorf("Invalid UUID format"))
	}
	log.Print("validateUUID:end")
	return
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log.Print("providerConfigure:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token: d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfigure:end")
		return &config, nil
	}
}

func Resource() *schema.Resource {
	log.Print("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
		Schema: map [string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "name of home",
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
				Description: "Description of home",
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				Description: "domain of home, as in *.cloudfront.net",
			},
			"town": {
				Type: schema.TypeString,
				Required: true,
				Description: "The town to which this home will belong",
			},
			"content_version": {
				Type: schema.TypeInt,
				Required: true,
				Description: "The content version of the home",
			},
		},
	}
	log.Print("Resource:end")
	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Print("resourceHouseCreate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
		"description": d.Get("description").(string),
		"town": d.Get("town").(string),
		"domain_name": d.Get("domain_name").(string),
		"content_version": d.Get("content_version").(int),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil{
		return diag.FromErr(err)
	}

	url:= config.Endpoint+"/u/"+config.UserUuid+"/homes"
	log.Print("URL: "+url)

	//construct the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}
	//set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	//parse response json
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}

	//StatusOK = 200 HTTP response code
	if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to create home resouce, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)

	log.Print("resourceHouseCreate:end")

	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Print("resourceHouseRead:start")
	var diags diag.Diagnostics

	config := m.(*Config)
	homeUUID := d.Id()

	url:= config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Print("URL: "+url)

	//construct the HTTP request
	req, err := http.NewRequest("GET", url , nil)
	if err != nil {
		return diag.FromErr(err)
	}
	//set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()
	var responseData map[string]interface{}

	if resp.StatusCode == http.StatusOK{
		//parse response json
		
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", responseData["name"].(string))
	d.Set("description", responseData["description"].(string))
	d.Set("domain_name", responseData["domain_name"].(string))
	d.Set("content_version", responseData["content_version"].(float64))
	} else if resp.StatusCode != http.StatusNotFound {
		d.SetId("")
	} else if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to read home resouce, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	log.Print("resourceHouseRead:end")

	return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Print("resourceHouseUpdate:start")
	var diags diag.Diagnostics

	config := m.(*Config)
	homeUUID := d.Id()

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
		"description": d.Get("description").(string),
		"content_version": d.Get("content_version").(int),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil{
		return diag.FromErr(err)
	}

	url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Print("URL: "+url)

	//construct the HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}
	//set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	//StatusOK = 200 HTTP response code
	if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to update home resouce, status_code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	log.Print("resourceHouseUpdate:end")

	d.Set("name", payload["name"])
	d.Set("description", payload["description"])
	d.Set("content_version", payload["content_version"])

	return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Print("resourceHouseDelete:start")
	var diags diag.Diagnostics

	config := m.(*Config)
	homeUUID := d.Id()

	url:= config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Print("URL: "+url)

	//construct the HTTP request
	req, err := http.NewRequest("DELETE", url , nil)
	if err != nil {
		return diag.FromErr(err)
	}

		//set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	//StatusOK = 200 HTTP response code
	if resp.StatusCode != http.StatusOK{
		return diag.FromErr(fmt.Errorf("failed to Delete home resouce, status_code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	d.SetId("")
	log.Print("resourceHouseDelete:end")
	return diags
}
