package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
)
type DelCommand struct{}

func (e DelCommand) deleteDnsRecord(dnsrecordid string, input *Input) {

    //println("Delete zone_id: ", zone_id, " with ", dns_record_id)

    url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", input.ZoneId, dnsrecordid)

    req, err := http.NewRequest("DELETE", url, nil)
    req.Header.Add("Authorization", input.ApiKey)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error while reading the response bytes:", err)
    }
    //fmt.Println(string([]byte(body)))
}

func (e DelCommand) scanDnsRecords(input *Input) CFGetDnsResponseResult {
    url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", input.ZoneId )

    req, err := http.NewRequest("GET", url, nil)
     req.Header.Add("Authorization", input.ApiKey)
    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error while reading the response bytes:", err)
    }
    //fmt.Println(string([]byte(body))) .

    response := CFGetDnsResponseResult{}
    json.Unmarshal([]byte(body), &response)
    return response
}

func (e DelCommand) Run(input *Input, options map[string]string) {

    CFGetDnsResponseResult := e.scanDnsRecords(input)

    comment := fmt.Sprintf("[cniPluginCloudFlare] points to v6-nat-vm cniHash: %s",HashValues(options["CNI_CONTAINERID"], options["CNI_IFNAME"]))
    
    // Find, the state info, if possible by 
    for _, v := range CFGetDnsResponseResult.Result {
        if v.Comment == comment {
            e.deleteDnsRecord(v.Id, input)
        }
    }        
}