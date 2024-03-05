# aws-session-go

## A simple tool to generate an AWS MFA session

> **NOTES:** <br>
> - Expects to use  config file and not environment variables for your local aws config (A custom file can be specified)
> - Expects MFA device to be named the same as your IAM username (This can be adjusted if needed. Please just submit a PR or issue)
> - Currently only supports for IAM users (SSO and other users can be included later)
> - ==This tool will modify your aws credentials file if not alternate credentials file is specified==

### **Usage**
```
Usage of aws-session-go:
  -creds-file string
    	Path to AWS credentials file
  -duration int
    	Duration in seconds (default 3600)
  -help
    	Show usage
  -mfa-token string
    	MFA token (required)
  -profile string
    	AWS profile (default "default")
```

### Examples
```
Usage: go run main.go -mfa-token <mfa-token> -profile <profile> -duration <duration> -creds-file <path/to/credentials/file>
Example: go run main.go -mfa-token 123456 -profile default -duration 3600 -creds-file /path/to/credentials/file
```

#### *Installation*
```
git clone https://github.com/emilio507/aws-session-go.git
cd aws-session-go
go build . 
sudo mv aws-session-go /usr/local/bin
aws-session-go -h
```