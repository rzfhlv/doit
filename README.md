# Just DOIT

## Requirement

- go version go1.19 or higher
- docker and docker-compose

## How to use

- clone from repo: 

    ``` git clone https://github.com/rzfhlv/doit.git ```

- in your root project directory copy environment file:

    ``` cp env.example .env ```

- golang initialize:

    ``` go mod tidy ```

- run the app:

    ``` make up ```

- check the app with curl:

    ``` curl -X GET http://localhost:8090/v1/health-check ```

- if all running properly, will show response like this:
    
    ``` {"status":"ok", "message":"I'm health"} ```