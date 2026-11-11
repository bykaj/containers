# paperless-tessdata

A data-only image with a complete [Tesseract](https://github.com/tesseract-ocr/tesseract) `tessdata` folder for [Paperless-ngx](https://github.com/paperless-ngx/paperless-ngx). The image has no shell and runs nothing. It exists to be mounted as a Kubernetes [image volume](https://kubernetes.io/docs/concepts/storage/volumes/#image).

## Why

The Paperless image only ships Debian's `fast` models, and has no Dutch model at all. Paperless can add languages with `PAPERLESS_OCR_LANGUAGES`, but that runs `apt install` on every start. Under a restrictive egress policy, that download can race the DNS-based rules and fail the start. With this image the kubelet pulls the models, so the pod needs no network access for them. The image also carries the more accurate `best` models.

## Contents

`/tessdata/` holds:

- the [`tessdata_best`](https://github.com/tesseract-ocr/tessdata_best) models (the image tag is their version) for Dutch (`nld`), English (`eng`), German (`deu`), French (`fra`), Spanish (`spa`) and Italian (`ita`)
- `osd`, for orientation and script detection (page rotation)
- `configs/`, `tessconfigs/` and `pdf.ttf`, taken from the Tesseract source of the version Paperless runs (5.5.0), because Tesseract reads its output configs and PDF font from the same folder

Every downloaded file is pinned by sha256 in the `Dockerfile`.

## Usage

Mount the folder over Paperless's own tessdata directory. Containerd only mounts directories from an image volume, so mount the `tessdata` subPath, not single files:

```yaml
persistence:
  tessdata:
    type: image
    image:
      repository: ghcr.io/bykaj/paperless-tessdata
      tag: 4.1.0@sha256:<digest>
    globalMounts:
      - path: /usr/share/tesseract-ocr/5/tessdata
        readOnly: true
        subPath: tessdata
```

Then pick languages with `PAPERLESS_OCR_LANGUAGE` (for example `nld+eng`), and leave `PAPERLESS_OCR_LANGUAGES` unset.

## Updating

- **Models:** a new `tessdata_best` version needs a new sha256 for every model in the `Dockerfile`.
- **Tesseract:** when Paperless ships a new Tesseract version, bump `TESSERACT_VERSION` and its tarball checksum. For a new major version, also move the mount path (`5` is the major version).
- **Note:** `best` models are slower than `fast` ones, and every extra language in `PAPERLESS_OCR_LANGUAGE` slows OCR further.
