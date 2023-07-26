class Solution(object):
    def addBinary(self, a, b):

        a1=len(a)-1
        b1=len(b)-1
        carry=0
        # print(a1 ," : ",b1)

        result="0"

        while (a1 >= 0) or (b1 >=0):
            # print(a1," <-index-> ",b1)
            # print(a[a1]," <-value-> ",b[b1])

            carry=int(a[a1])+int(b[b1])
            # print("carry : ",carry)

            if carry %2==0:

                result+= "0"
                
                print("if ",result)
            else:
                
                result=result+str(carry)
                print("else ",result)

            print("Output : ", result)
            a1=a1-1
            b1=b1-1

        return result
        """
        :type a: str
        :type b: str
        :rtype: str
        """

s1=Solution()
# print("Result",s1.addBinary("101","11"))
print("Call",s1.addBinary("1010","1011"))

# 101  1010
#  11  1011
# ---------
# 100 10101