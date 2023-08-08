 
def lengthOfLastWord(s: str) -> int:
    s=s.strip()
    split=s.split()
    last=split[len(split)-1]
    return len(last)
 
 


def main():
    print(lengthOfLastWord("hello world"))
 
if __name__=="__main__":
    main()