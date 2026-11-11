# syntax=docker/dockerfile:1

# Data-only image for Paperless-ngx: a complete Tesseract tessdata folder with
# the most accurate ("best") trained models, mounted as a Kubernetes image
# volume over Paperless's tessdata directory (subPath tessdata). It has no
# shell and runs nothing. The checksums pin each file; a VERSION bump needs
# new sha256s. See README.md

# Tesseract reads its configs (hocr, pdf, txt, ...) and the PDF renderer's
# font from the tessdata folder too. tessdata_best only symlinks them, so take
# them from the Tesseract source of the version Paperless-ngx runs (5.5.0)
FROM docker.io/library/alpine:3.24 AS configs
ARG TESSERACT_VERSION=5.5.0
ADD --checksum=sha256:f2fb34ca035b6d087a42875a35a7a5c4155fa9979c6132365b1e5a28ebc3fc11 \
    https://github.com/tesseract-ocr/tesseract/archive/refs/tags/${TESSERACT_VERSION}.tar.gz /tesseract.tar.gz
RUN \
    tar -xzf /tesseract.tar.gz -C /tmp \
    && mkdir /tessdata \
    && cp -R \
        "/tmp/tesseract-${TESSERACT_VERSION}/tessdata/configs" \
        "/tmp/tesseract-${TESSERACT_VERSION}/tessdata/tessconfigs" \
        "/tmp/tesseract-${TESSERACT_VERSION}/tessdata/pdf.ttf" \
        /tessdata/ \
    && rm /tessdata/configs/Makefile.am /tessdata/tessconfigs/Makefile.am \
    && chmod -R a=rX /tessdata

FROM scratch
ARG VERSION

COPY --from=configs /tessdata/ /tessdata/

# Dutch, English, German, French, Spanish and Italian, plus orientation and
# script detection (osd, used to rotate pages)
ADD --chmod=444 --checksum=sha256:8407331d6aa0229dc927685c01a7938fc5a641d1a9524f74838cdac599f0d06e \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/deu.traineddata /tessdata/deu.traineddata
ADD --chmod=444 --checksum=sha256:8280aed0782fe27257a68ea10fe7ef324ca0f8d85bd2fd145d1c2b560bcb66ba \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/eng.traineddata /tessdata/eng.traineddata
ADD --chmod=444 --checksum=sha256:907743d98915c91a3906dfbf6e48b97598346698fe53aaa797e1a064ffcac913 \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/fra.traineddata /tessdata/fra.traineddata
ADD --chmod=444 --checksum=sha256:8df9c89176fb93f56bf4b2d4ede04c01c1f31d4b7697fbd76cc336df700f3f38 \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/ita.traineddata /tessdata/ita.traineddata
ADD --chmod=444 --checksum=sha256:92e7a1ad4bf8082e268de57c7823316ec024935702c6ed2a1e473b3a071aa733 \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/nld.traineddata /tessdata/nld.traineddata
ADD --chmod=444 --checksum=sha256:9cf5d576fcc47564f11265841e5ca839001e7e6f38ff7f7aacf46d15a96b00ff \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/osd.traineddata /tessdata/osd.traineddata
ADD --chmod=444 --checksum=sha256:e2c1ffdad8b30f26c45d4017a9183d3a7f9aa69e59918be4f88b126fac99ab2c \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/spa.traineddata /tessdata/spa.traineddata
