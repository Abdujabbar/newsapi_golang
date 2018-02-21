# newsapi-golang
A Go client for the [News API](https://newsapi.org/docs/)
## General 

This is a Go client library for News API V2. 


# Usage

First download the package, run in your console: 
```
go get -u github.com/Abdujabbor/newsapi_golang
```
And then  import the package
```
import apiClient "github.com/Abdujabbor/newsapi_golang"
```

And then we can create our instance and calling methods:
```
req := apiClient.InitNewsAPIClient("https://newsapi.org/v2/", "YOUR_API_KEY")
response, err := req.GetSources()
if err != nil {
    panic(err)
}

fmt.Println(response.Sources)
```

in object response you can get response from service for more details lookup documentation:
https://newsapi.org/docs



## Support
For suggestions
Reach out at [abdujabbor1987@gmail.com]('mailto:abdujabbor1987@gmail.com')