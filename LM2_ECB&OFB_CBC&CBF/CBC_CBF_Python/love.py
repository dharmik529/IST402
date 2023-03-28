from Cryptodome.Cipher import AES

# Define the key and IV for encryption
key = bytes('This is a key123', encoding='utf-8')
iv = bytes('This is an IV456', encoding='utf-8')

# Define the plaintext to be encrypted
plaintext = input("Enter the plaintext to be encrypted: ").encode()
print("Plaintext:", plaintext.decode())

# Pad the plaintext to be a multiple of 16 bytes for AES encryption
length = 16 - (len(plaintext) % 16)
plaintext += length * b' '

# Define the AES cipher for CBC mode
cbc_cipher = AES.new(key, AES.MODE_CBC, iv)

# Perform CBC encryption
ciphertext = cbc_cipher.encrypt(plaintext)
print("CBC Ciphertext:", ciphertext.hex())

# Define the AES cipher for CFB mode
cfb_cipher = AES.new(key, AES.MODE_CFB, iv)

# Perform CFB encryption
ciphertext = cfb_cipher.encrypt(plaintext)
print("CFB Ciphertext:", ciphertext.hex())

# Perform the decryption on the message using CBC mode
cbc_decipher = AES.new(key, AES.MODE_CBC, iv)
decrypted_cbc = cbc_decipher.decrypt(ciphertext)
print("CBC Decrypted:", decrypted_cbc.decode())

# Perform the decryption on the message using CFB mode
cfb_decipher = AES.new(key, AES.MODE_CFB, iv)
decrypted_cfb = cfb_decipher.decrypt(ciphertext)
print("CFB Decrypted:", decrypted_cfb.decode())
