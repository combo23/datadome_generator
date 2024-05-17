# DataDome Generator

AntibotAIO is an application that generates datadome payload to bypass their security on sites like footlocker.
Last supported version of DataDome is 4.19.0  
Alternatively you can use [CapSolver](https://www.capsolver.com/)

## Dependencies

To successfully start the application you need to create your own `.env` file based on `.env.example`.
In addition you need to download all libraries used in this project, to do so simply run `go get .`.
Also it's worth stating that this project uses `MongoDB` so you will need to replicate this on your own as well.

## Usage

```bash
go build .
./datadome_generator
```

## Endpoints

### Health Check
**Base URL: /status**

Example

```bash
curl -X GET /status
```

Response

```json
{
  "status": "alive",
}
```

### Generate DataDome payload
**Base URL: /datadome**
**Headers:**
- `X-api-key` (required): Your api key.

Example

```bash
curl -X POST /datadome -H "Content-Type: application/json" -H "X-api-key: test" -d '{
"site":"tickets.mancity.com",
    "user-agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome 119.0.0.0 Safari/537.36",
    "cid":"xxxxx",
    "responsePage":"origin",
    "IP":"your_ip",
    "referer":"https://tickets.mancity.com/en-GB/categories/Men'\''s%20Tickets"
}' 
```

Response

```json
{
    "ch": "jsData=...", //ch payload is posted first
    "le": "jsData=...." //then le
}
```

### Add user
**Base URL: /register**
**Headers:**
- `X-api-key` (required): Admin api key.

```bash
curl -X POST /register -H "Content-Type: application/json" -H "X-api-key: test" -d '{"name":"@user_name", "solves":5000}' 
```

Response

```json
{
  "status": "success",
  "api_key": "XXXXXX"
}
```

### Ban user
**Base URL: /ban**
**Parameters:**
- `apikey` (required): API Key which is assigned to user you wanna ban.
**Headers:**
- `X-api-key` (required): Admin api key.

Example

```bash
curl -X GET /ban?apikey=xxx -H "X-api-key: test"
```

Response

```json
{
  "status": "success"
}
```

### Fetch number of left solves
**Base URL: /solves**
**Parameters:**
- `apikey` (required): API Key which you wanna check remaining solves for.

Example

```bash
curl -X GET /solves?apikey=xxx
```

Response

```json
{
    "solves": 1000
}
```
