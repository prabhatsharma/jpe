#!/bin/bash

# Generate a CA key file
openssl genrsa -aes256 -passout pass:somepassword -out server.pass.key 4096

# Generate the key file for the cert
openssl rsa -passin pass:somepassword -in server.pass.key -out server.key 

# Delete the CA key file. We don't need it anymore
rm server.pass.key

# Generate the CSR
echo '\n Generating CSR \n'
openssl req -new -key server.key -out server.csr -config csr.conf

# Generate the cert
printf "\n Generating certificate \n"
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt -extfile csr.conf -extensions v3_req

# Generate base64 encoded cert that needs to be placed in ValidatingWebhookConfiguration manifest
cat server.crt | base64 --wrap=0
# cat server.crt | base64


printf "\nChecking SAN \n"
# Validate SAN of the certificate
openssl x509  -noout -text -in server.crt | grep DNS:

printf "\n"
