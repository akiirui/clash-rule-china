# clash-rule-china

Clash rules based on [felixonmars/dnsmasq-china-list][links-felix-list].

_**Updated every Saturday.**_

## Links

- Apple
  - https://raw.githubusercontent.com/akiirui/clash-rule-china/main/release/apple.yaml
  - https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/apple.yaml
- China
  - https://raw.githubusercontent.com/akiirui/clash-rule-china/main/release/china.yaml
  - https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/china.yaml
- Google
  - https://raw.githubusercontent.com/akiirui/clash-rule-china/main/release/google.yaml
  - https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/google.yaml

## Usage

### Rule Providers Settings

```yaml
rule-providers:
  apple:
    behavior: domain
    type: http
    url: "https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/apple.yaml"
    interval: 172800
    path: ./apple.yaml
  china:
    behavior: domain
    type: http
    url: "https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/china.yaml"
    interval: 172800
    path: ./china.yaml
  google:
    behavior: domain
    type: http
    url: "https://cdn.jsdelivr.net/gh/akiirui/clash-rule-china@main/release/google.yaml"
    interval: 172800
    path: ./google.yaml
```

### Rules Settings

```yaml
rules:
  # DIRECT
  - RULE-SET,apple,DIRECT
  - RULE-SET,china,DIRECT
  - RULE-SET,google,DIRECT # Not recommend
  # FINAL
  - GEOIP,CN,DIRECT,no-resolve
  - MATCH,PROXY
```

## Thanks

- [felixonmars/dnsmasq-china-list][links-felix-list]

[links-felix-list]: https://github.com/felixonmars/dnsmasq-china-list
