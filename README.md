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