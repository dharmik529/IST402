# Importing necessary libraries
from Cryptodome.Cipher import AES
from Cryptodome.Util.Padding import pad, unpad
import os

# Defining the CBC mode encryption function
# this function accepts a plaintext (entered by the user, uses the initialized random key, and initialized vector)
# the AES.new creates a new cipher object, where it passes they key, in CBC mode, and the initialized vector
#Then, it encrypts the plaintext using this cipher object by calling its encrypt() method on the plaintext that is padded to the block size of the algorithm
# The pad() function is used to ensure that the plaintext is a multiple of the block size required by the AES encryption algorithm. 
# This is necessary because AES (Advanced Encryption Standard) works on fixed-size blocks of 128 bits (16 bytes) each, and if the plaintext is not a multiple of the block size, 
# it needs to be padded with additional bytes to make it a complete block.


def cbc_encrypt(plaintext, key, iv):
    cipher = AES.new(key, AES.MODE_CBC, iv)      # Creating a new cipher object using AES encryption algorithm with CBC mode and initialization vector iv
    ciphertext = cipher.encrypt(pad(plaintext, AES.block_size))    # Encrypting the plaintext by padding it to the block size and passing it to the cipher object
    return ciphertext

# Defining the CBC mode decryption function
def cbc_decrypt(ciphertext, key, iv):
    cipher = AES.new(key, AES.MODE_CBC, iv)      # Creating a new cipher object using AES encryption algorithm with CBC mode and initialization vector iv
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)   # Decrypting the ciphertext by passing it to the cipher object and then unpadding it to the block size
    return plaintext

# Defining the CFB mode encryption function
def cfb_encrypt(plaintext, key, iv):
    cipher = AES.new(key, AES.MODE_CFB, iv)     # Creating a new cipher object using AES encryption algorithm with CFB mode and initialization vector iv
    ciphertext = cipher.encrypt(plaintext)      # Encrypting the plaintext by passing it to the cipher object
    return ciphertext

# Defining the CFB mode decryption function
def cfb_decrypt(ciphertext, key, iv):
    cipher = AES.new(key, AES.MODE_CFB, iv)     # Creating a new cipher object using AES encryption algorithm with CFB mode and initialization vector iv
    plaintext = cipher.decrypt(ciphertext)      # Decrypting the ciphertext by passing it to the cipher object
    return plaintext

# Main function to get user input and execute the encryption and decryption
def main():
    plaintext = input("Enter plaintext to encrypt: ")    # Get the plaintext input from user
    key = os.urandom(16)    # Generating a 16-byte (128-bit) random key using os.urandom function
    iv = os.urandom(16)     # Generating a 16-byte (128-bit) random initialization vector using os.urandom function
    print("Plaintext:", plaintext)  # Print the plaintext
    print(f'Key: {key}')           # Print the randomly generated key
    print(f'IV: {iv}')             # Print the randomly generated initialization vector
    
    # CBC mode encryption and decryption
    ciphertext = cbc_encrypt(plaintext.encode('utf-8'), key, iv)  # Encrypt the plaintext using CBC mode and the randomly generated key and iv
    print("CBC Mode")
    print("Ciphertext:", ciphertext.hex())     # Print the ciphertext
    decrypted_plaintext = cbc_decrypt(ciphertext, key, iv).decode()   # Decrypt the ciphertext using CBC mode and the randomly generated key and iv
    print("Decrypted plaintext:", decrypted_plaintext)   # Print the decrypted plaintext

    # CFB mode encryption and decryption
    ciphertext = cfb_encrypt(plaintext.encode('utf-8'), key, iv)  # Encrypt the plaintext using CFB mode and the randomly generated key and iv
    print("\nCFB Mode")
    print("Ciphertext:", ciphertext.hex())     # Print the ciphertext
    decrypted_plaintext = cfb_decrypt(ciphertext, key, iv).decode()   # Decrypt the ciphertext using CFB mode and the randomly generated key and iv
    print("Decrypted plaintext:", decrypted_plaintext)   # Print the decrypted plaintext

if __name__ == '__main__':
    main()
