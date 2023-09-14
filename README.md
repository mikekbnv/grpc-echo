# grpc-echo

## Usage 

```bash
grpcurl -plaintext -d '{"text": "Your message"}' localhost:50051 echo.Echo/Echo