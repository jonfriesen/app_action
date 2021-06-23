# DigitalOcean App Platform Image and DigitalOcean Registry publish
This action uses [doctl](https://github.com/digitalocean/action-doctl) to upload docker images to Digital Ocean's Container Registry and then deploy it to App Platform

PRs are welcome.

# Usage
Add this step to a job to automatically build an image from the Dockerfile and publish it with a unique tag (based on the commit SHA). The image will also be tagged with `latest`.

### Example:

```yaml
    - name: Publish Image to Container Registry
      uses: ripplr-io/docr-docker-publish@v1
      with:
        image_path: registry-path/image-path
```

# Inputs
- `image_path` - (**Required**) Image path in the registry. In the format `registry-name/image-name`.
- `sha_size` - (Optional) Number of characters from the commit SHA. Default is `'8'`
- `dockerfile` - (Optional) The path + name of the Dockerfile you want to build (-f flag). Default is `'Dockerfile'` Example: `./docker/dev.Dockerfile`.
- `docker_build_context` - (Optional) The `docker build` context. Default is `'.'`.
- `app_image_name` - Name of the app on App Platform.

# Outputs
The urls of the uploaded images will be accessible with the variables:
- `image_url` - Url of the uploaded image with the SHA tag.
- `image_latest_url` - Url of the uploaded image with the latest tag.

## License

This GitHub Action and associated scripts and documentation in this project are released under the [MIT License](LICENSE).
