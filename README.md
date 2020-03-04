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
            "http://example.tangelo/api/cropped/part0",
            "http://example.tangelo/api/cropped/part1",
            "http://example.tangelo/api/cropped/part2",
            "http://example.tangelo/api/cropped/part3",
        ]
    ```

# Example with curl
