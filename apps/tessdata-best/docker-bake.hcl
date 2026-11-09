target "docker-metadata-action" {}

variable "APP" {
  default = "tessdata-best"
}

variable "VERSION" {
  // renovate: datasource=github-tags depName=tesseract-ocr/tessdata_best
  default = "4.1.0"
}

variable "SOURCE" {
  default = "https://github.com/tesseract-ocr/tessdata_best"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
  labels = {
    "org.opencontainers.image.source" = "${SOURCE}"
  }
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
