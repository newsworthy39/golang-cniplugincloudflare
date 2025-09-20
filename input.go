package main

import (
    "encoding/json"
    "os"
    "io"
     "crypto/sha1"
    "encoding/hex"
)

type CFGetDnsResponseResult struct {
    Result []CFGetDnsResponse `json:"result"`
}

type CFGetDnsResponse struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Content string `json:"content"`
    Comment string `json:"comment"`
}

type Ips struct {
    Address string `json:"address"` 
    Gateway string `json:"gateway"` 
    Ifindex int `json:"interface"` 
}

type PreviousPresult struct {
    Ips []Ips `json:"ips"`
}

type Input struct {
    raw               []byte
    CNIVersion        string          `json:"cniVersion"`
    Name              string          `json:"name"`
    Type              string          `json:"type"`
    ApiKey            string          `json:"api_key"`
    DnsRecordType     []string        `json:"dns_record_type"`
    DnsRecord         string          `json:"dns_record"`
    ZoneId            string          `json:"zone_id"`
    PrevResult        PreviousPresult `json:"prevResult"`
}

func ReadJSONInput() *Input {

    raw, _:= io.ReadAll(os.Stdin)

    var input Input
    json.Unmarshal(raw, &input)

    input.raw = raw
    
    return &input
}

func  HashValues( arg1 string, arg2 string) string {
    combined := arg1 + ":" + arg2
    hash := sha1.Sum([]byte(combined))

    // Convert to hex string
    hashString := hex.EncodeToString(hash[:])

    return hashString
}