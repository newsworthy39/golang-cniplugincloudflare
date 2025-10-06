# cniPluginCloudflare

A chained plugin, that can create AAAA records in Cloudflare under a zone


    {
        "type": "cniPluginCloudflare",
        "zone_id": "<CF ZONE ID>",
        "dns_record": "<DNS RECORD>",
        "api_key": "Bearer <Account API Key>"
        "dns_record_type": [ "AAAA", "A" ]
    }

A dns_record_type of AAAA, means that only ipv6 addresses are created. A, means ipv4 adresses from the container are created to.

## build
    go build -o bin

## configuration example,
e.g /etc/cni/net.d/default/ctl-ipvlan-v6-gojira.conflist:

    plugins: [
        { 
            "type": "ipvlan",
            ...
        },
        {
            "type": "cniPluginCloudflare",
            "zone_id": "<zone id>",
            "dns_record": "gojira.domain.tld",
            "api_key": "Bearer <TOKEN>,
            "dns_record_type": [ "AAAA" ]
        }
    ]