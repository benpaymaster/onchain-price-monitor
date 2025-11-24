provider "docker" {}

resource "docker_image" "backend" {
  name         = "backend:latest"
  keep_locally = true
}

resource "docker_container" "backend" {
  image = docker_image.backend.latest
  name  = "backend"
  ports {
    internal = 5000
    external = 5000
  }
}
