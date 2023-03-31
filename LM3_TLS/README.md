# LM3: Assignment Client GoLang TLS Programming

Using GoLang perform TLS encryption and decryption on a user string input.

You can work together with your pair on the project.  You need to upload your project to github and provide a screen capture png of it running.

Step 1: Make sure to create the following files in the project working for example C:/Users/joeoa/GolandProjects/TLS_Server
```bash
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key

# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
```
```bash
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
![Output](https://github.com/dharmik529/IST402/blob/main/Assets/buildGoLangTLSCert.PNG)

Step 2: Create a Go Project here is the server code 
[main-1.go](https://github.com/dharmik529/IST402/blob/main/LM3_TLS/main-1.go)

Step 3 In the Go project add the client code
[client.go](https://github.com/dharmik529/IST402/blob/main/LM3_TLS/client.go)

Step 4: Make sure you have the following files
Step 5: Run the server then run the client
