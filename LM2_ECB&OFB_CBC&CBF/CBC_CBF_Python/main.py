from Cryptodome.Cipher import AES
import os

# Initialization Vector (IV) for CBC mode
IV = os.urandom(16)

# Key for AES encryption
key = os.urandom(16)

# Create AES cipher objects for CBC and CFB modes
cbc_cipher = AES.new(key, AES.MODE_CBC, IV)
cfb_cipher = AES.new(key, AES.MODE_CFB, IV)

# Get input plaintext from user
plaintext = input("Enter the plaintext: ")

# Pad the plaintext to a multiple of 16 bytes (the AES block size)
if len(plaintext) % 16 != 0:
    plaintext = plaintext.ljust(len(plaintext) + 16 - len(plaintext) % 16)

# Encrypt the plaintext using CBC mode
cbc_ciphertext = cbc_cipher.encrypt(plaintext.encode('utf-8'))

# Encrypt the plaintext using CFB mode
cfb_ciphertext = cfb_cipher.encrypt(plaintext.encode('utf-8'))

# Print the results
print("Plaintext:", plaintext)
print("IV (for CBC mode):", IV.hex())
print("Key:", key.hex())
print("Ciphertext (CBC mode):", cbc_ciphertext.hex())
print("Ciphertext (CFB mode):", cfb_ciphertext.hex())

