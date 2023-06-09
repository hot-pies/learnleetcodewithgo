from typing import List

def plusOne(digits: List[int]) -> List[int]:

    number =''
    for i in range(len(digits)):
        number+=str(digits[i])
    
    number=int(number)
    number=number+1
    # digits = [int(d) for d in str(number)]

    digits=[]
    for i in str(number):
        digits.append(int(i))
        
    return digits


def main():
    nums: List[int] =[9,9,9]
    
    print(plusOne(nums))

if __name__ == "__main__":
    main()