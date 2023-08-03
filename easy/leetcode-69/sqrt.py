import math
class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """

        left=0
        right=46340 # Max inv squr is 46340 of max int value 2147483647. squr shourl not be grater then 46340

        # check is input is negative value 
        if x <0:
            return 0
        print("Squr of : ",x)

        while left <= right:
            # print(" Left : ",left)
            # print(" Right : ",right)
            if left * left == x:
                return left
            if right * right == x:
                return right
            # increment left and decrement rigth
            middle=left+((right-left)/2)

            print(" Middle : ",middle)

            result=middle*middle
            if  result== x:
                return middle
            elif (result < x):
                left = middle + 1
                print(" Left : ",left)
            else:
                right = middle-1
                print(" Right : ",right)
        return right

sq = Solution()
sq.mySqrt(144)