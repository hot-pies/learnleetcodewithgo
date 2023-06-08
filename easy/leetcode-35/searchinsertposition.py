from typing import List

def searchInsert(nums: List[int], target: int) -> int:

        if len(nums)-1 ==0:
            if nums[0]==target:
                return 0
            if nums[0]<target:
                return 1
            if nums[0]>target:
                return 0

        right=len(nums)-1

        if nums[right]<target:
            return right+1

        left=0
        
        while left <= right:
            print("left and right", left,right)
            mid=round((right+left)/2)
            print("mid : ",mid, nums[mid])
            if nums[mid]==target:
                return mid
            
            if nums[mid]<target:
                print("increment left")
                left=mid+1
            else: 
                right=mid-1
                print("decrement right ", right)
        
        return left


def main():
    nums: List[int] = [1,3,5,6]
    # nums: List[int] = [1,3,5]
    # nums: List[int] = [1]
    
    # target: int = 5
    # target: int = 4
    # target: int = 2
    target: int = 7
    index=searchInsert(nums,target)
    print(index)

if __name__=='__main__':
    main()


# Input: nums = [1,3,5,6], target = 5
# Input: nums = [1,3,5,6], target = 2
# Input: nums = [1,3,5,6], target = 7