# Website Metatags-go

This lib aims to create a function that returns all meta tags from a website. This lib is a paclage form [microblog felipe mateus](https://microblog.felipemateus.com) replace old lib [metatags](https://github.com/eufelipemateus/metatags) from laravel.

## Functions Exported

### Get from url

```go
func GetMetatagsFromURL(url string) map[string]string{
```

This function aims to get all url meta tags from a remote website.

### Get from HTML

```go

func GetMetatagsFromHTML( html string) map[string]strin
```

This function aims to obtain all meta tags from a string that contains valid html.


## Author

**[Felipe Mateus](https://felipemateus.com)** - [eufelipemateus](https://github.com/eufelipemateus)
