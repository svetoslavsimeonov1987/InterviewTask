# Question: What will happen when you run this Python script?

NUMBER = 5

def main():
    NUMBER = 10 + NUMBER
    NUMBER = add_five(add_five(NUMBER))
    print(NUMBER)

def add_five(num):
    new_number = num + 5
    return new_number

if __name__ == "__main__":
    main()
