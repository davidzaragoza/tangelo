This is a test for Tangelo Games

# Service Render API
The api have the following endpoints:

- Health Check:

    *GET /api/v1/health*

    Response:
    ```json
    {
        "status": "OK"
    }
    ```

- Image Cropper:

    *POST /api/v1/crop*

    Response:
    ```json
        [
            "http://example.tangelo/api/cropped/part0_0.jpg",
            "http://example.tangelo/api/cropped/part0_1.jpg",
            "http://example.tangelo/api/cropped/part1_0.jpg",
            "http://example.tangelo/api/cropped/part1_1.jpg",
        ]
    ```

- Get Cropped Image

    *GET /api/v1/cropped/:image*

    Response:
    
    The image content

# Example with curl

## Health Check
```bash
curl --location --request GET 'http://localhost:8080/api/v1/health'
```

## Crop image
```bash
curl --location --request POST 'http://localhost:8080/api/v1/crop' --form 'image=@tangelo.png'
```

## Get cropped image
```bash
curl --location --request GET 'http://localhost:8080/api/v1/cropped/tangelo.png_0_0.jpg' --out croppedImage0_0.jpg
```

# Execute locally
In order to execute the application locally, docker must be previously installed on the machine. The image was tested in Docker Swarm mode, but probably docker-compose should work too.

First, the docker image must be built using the following command:
```bash
docker build -t tangelo-renderer .
```

Now, enter to the deployments folder.
```bash
cd deployments
```

This example creates a docker service that execute a Postgres Database, if you want to use another Postgres Database, comment the service and volume postgres from services.yml file. Update the configuration.json file according to the database.

Now deploy the stack:
```bash
docker stack deploy -c services.yml tangelo
```

Finally, check that the service is running:
```bash
docker service logs tangelo_renderer
```