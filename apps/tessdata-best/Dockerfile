# syntax=docker/dockerfile:1

# Data-only image: Tesseract's most accurate ("best") trained models, meant to
# be mounted as a Kubernetes image volume into an app's tessdata directory
# (subPath tessdata/<lang>.traineddata). It has no shell and runs nothing.
# The checksum pins each file; a VERSION bump needs the new sha256 too
FROM scratch
ARG VERSION

ADD --chmod=444 --checksum=sha256:92e7a1ad4bf8082e268de57c7823316ec024935702c6ed2a1e473b3a071aa733 \
    https://raw.githubusercontent.com/tesseract-ocr/tessdata_best/${VERSION}/nld.traineddata /tessdata/nld.traineddata
