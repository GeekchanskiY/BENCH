import bcrypt

def hash_password(password):
    salt = bcrypt.gensalt()
    hashed_password = bcrypt.hashpw(password.encode('utf-8'), salt)
    return hashed_password


def verify_password(hashed_password: str, input_password):
    
    return bcrypt.checkpw(password=bytes(input_password, 'utf-8'), hashed_password=bytes(hashed_password, 'utf-8'))


if __name__ == "__main__":
    password = "mysecurepassword"

    hashed_password = hash_password(password)
    print(hashed_password)
    is_verified = verify_password(hashed_password, password)

    if is_verified:
        print("Password verified successfully!")
    else:
        print("Incorrect password!")