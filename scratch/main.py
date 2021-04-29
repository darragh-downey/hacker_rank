


def memo_fact(f):
    print("inside memo_fact")
    memory = {}

    def inner(num):
        print("inside inner")
        if num not in memory:
            print("calculated", num, "factorial")
            memory[num] = f(num)
        print("leaving inner with", memory[num])
        return memory[num]
    
    print("leaving memo_fact")
    return inner


@memo_fact
def facto(num):
    print("inside facto")
    if num == 1:
        print("END with base case 1")
        return 1
    else:
        print("Calc", num, num - 1)
        return num * facto(num - 1)


print(facto(5))