import string

def brute_force(cipher_text):
    # Create a list of all possible characters
    characters = string.ascii_letters + string.digits + string.punctuation + " "

    # Try all possible combinations of the characters
    for key in range(len(characters)):
        # Decrypt the cipher text using the key
        plain_text = ""
        for char in cipher_text:
            index = characters.find(char)
            index = (index - key) % len(characters)
            plain_text += characters[index]

        # Print the decrypted text and the key used to decrypt it
        print("Key: " + str(key) + " | Decrypted text: " + plain_text)

encrypted = 'Frph ryhu khuh Zdwvrq'
print(encrypted)
brute_force(encrypted)
