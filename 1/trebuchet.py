import time

def get_code(line):
    numbers = [char for char in line if char.isdigit()]
    code = int(numbers[0] + numbers[-1])
    return code

def main():
    start_time = time.time()
    total_sum = 0

    with open("input.txt", "r") as file:
        for line in file:
            line = line.strip()
            code = get_code(line)
            total_sum += code

    print(total_sum)
    print(f"Execution time: {(time.time() - start_time) * 1000000:.0f} microseconds")

if __name__ == "__main__":
    main()