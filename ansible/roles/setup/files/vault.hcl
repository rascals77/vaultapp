api_addr          = "https://127.0.0.1:8200"
disable_mlock     = false
default_lease_ttl = "168h"
max_lease_ttl     = "720h"

# This path is what Vault sees in the docker container
storage "file" {
  path = "/vault/file"
}

listener "tcp" {
  address         = "0.0.0.0:8200"

  # To configure the listener to use a CA certificate, concatenate the
  # primary certificate and the CA certificate together. The primary
  # certificate should appear first in the combined file
  tls_cert_file   = "/vault/config/tls/certs/server.pem"
  tls_key_file    = "/vault/config/tls/certs/server-key.pem"
  tls_min_version = "tls12"
}
