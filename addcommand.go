package main

import (
    "bytes"
    "fmt"
    "encoding/json"
    "strings"
    "net/http"
    "io/ioutil"
)
type AddCommand struct{}

func (e AddCommand)  Run(input *Input, options map[string]string)  {

    url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", input.ZoneId )

    // TODO: Lets reverse this, around the ips instead.
    for _, v := range input.DnsRecordType {

        data := make(map [string]interface{})
        data["name"] = input.DnsRecord
        data["ttl"] = 3600
        data["type"] = v
        data["comment"] = fmt.Sprintf("[cniPluginCloudFlare] points to v6-nat-vm cniHash: %s",HashValues(options["CNI_CONTAINERID"], options["CNI_IFNAME"]))
        data["content"] = strings.Split(input.PrevResult.Ips[0].Address, "/")[0]
        data["proxied"] = true
        buf,_ := json.MarshalIndent(data, "", "   ")

        req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
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
    }  
}