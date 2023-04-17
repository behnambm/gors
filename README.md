# GORS: A simple way of handling CORS in Traefik

This plugin intercepts requests and processes the OPTIONS requests and returns the desired response.

### Features:

- Supports for disabling or enabling the plugin 
- Ability to config allowed `Origin`, `Headers`, `Methods` and `Max-Age`  


## Examples

#### Allow all origins 

```yaml
AllowedOrigins:
    - "*"
```

#### Limit allowed origins

```yaml
AllowedOrigins:
    - "https://foo.example.com"
    - "https://bar.example.com"
```


## Defaults:

```yaml
AllowedOrigins: empty list 
AllowedHeaders: empty list
Disabled: true
PreflightMaxAge: 3600
AllowedMethods: empty list
```

# TODO:

-  Ability to use Regex for allowed `Origin`




