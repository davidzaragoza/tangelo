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
