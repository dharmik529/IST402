from Cryptodome.Cipher import AES
from Cryptodome.Util.Padding import pad, unpad
import os

def cbc_encrypt(plaintext, key, iv):
    cipher = AES.new(key, AES.MODE_CBC, iv)
    ciphertext = cipher.encrypt(pad(plaintext, AES.block_size))
    return ciphertext

def cbc_decrypt(ciphertext, key, iv):
    cipher = AES.new(key, AES.MODE_CBC, iv)
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)
    return plaintext

def cfb_encrypt(plaintext, key, iv):
    cipher = AES.new(key, AES.MODE_CFB, iv)
    ciphertext = cipher.encrypt(plaintext)
    return ciphertext

def cfb_decrypt(ciphertext, key, iv):
    cipher = AES.new(key, AES.MODE_CFB, iv)
    plaintext = cipher.decrypt(ciphertext)
    return plaintext

def main():
    plaintext = input("Enter plaintext to encrypt: ")
    key = os.urandom(16)
    iv = os.urandom(16)
    print("Plaintext:", plaintext)
    
    # CBC mode encryption and decryption
    ciphertext = cbc_encrypt(plaintext.encode('utf-8'), key, iv)
    print("CBC Mode")
    print("Ciphertext:", ciphertext.hex())
    decrypted_plaintext = cbc_decrypt(ciphertext, key, iv).decode()
    print("Decrypted plaintext:", decrypted_plaintext)

    # CFB mode encryption and decryption
    ciphertext = cfb_encrypt(plaintext.encode('utf-8'), key, iv)
    print("\nCFB Mode")
    print("Ciphertext:", ciphertext.hex())
    decrypted_plaintext = cfb_decrypt(ciphertext, key, iv).decode()
    print("Decrypted plaintext:", decrypted_plaintext)

if __name__ == '__main__':
    main()
