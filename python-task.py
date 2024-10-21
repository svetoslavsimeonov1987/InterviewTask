# Question: What will happen when you run this Python script?

def main():
    NUMBER = 2
    try:
        ARR = [1, 2, 3, 4, 5]
        NUMBER = 2 + NUMBER**3 + ARR[::-1][0]
        add_five(add_five(NUMBER))
        print(NUMBER)
    except:
        print("An error occured.")

def add_five(num):
    new_number = num + 5
    return new_number

if __name__ == "__main__":
    main()
