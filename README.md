# Doccopilot

## How to deploy

1. Build the docker image

```
docker build -t doccopilot:test .
```

2. Start a container

```
docker run -dp 8080:8080 \
    -e API_KEY="Your API Key" \
    --name doccopilot \
    doccopilot:test
```
