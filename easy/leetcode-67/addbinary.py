class Solution(object):
    def addBinary(self, a, b):

        a1=len(a)-1
        b1=len(b)-1
        carry=0
        # print(a1 ," : ",b1)

        result="0"

        while (a1 >= 0) or (b1 >=0):
            sum=0
            # print("Reset : ",sum)
            print("int(a[a1]) : ",int(a[a1]),"  __  int(b[b1] : ",int(b[b1]))
            sum=int(a[a1])+int(b[b1])
            print("Sum = ", sum, "carry : ",carry)
            if sum == 2 and carry==1:
                result+= "1"
                carry=0
            elif sum ==2 and carry==0:
                result+= "0"
                carry=1
            elif sum ==1 and carry==1:
                result+="0"
                carry=0
            elif sum==1 and carry==0:
                result+="1"
                carry=0
            elif sum==0 and carry==1:
                result+="1"
                carry=0
            # else:
            #     result="1"
            # if carry %2==0:
                
            #     print("if ",result)
            #     carry=1
            # else:
            #     result=result+str(carry)
            #     print("else ",result)


            print("Result : ", result)
            # print("old carry : ", carry)
            a1=a1-1
            b1=b1-1
        if carry >0:
            result+="1"

        return result
        """
        :type a: str
        :type b: str
        :rtype: str
        """

s1=Solution()
print("Result",s1.addBinary("101","11"))
# print("Call",s1.addBinary("1010","1011"))

# 101  1010
#  11  1011
# ---  ----
# 100 10101