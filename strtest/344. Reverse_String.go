package strtest

//字符串反转，没啥好说的

func reverseString0(s []byte)  {
	for i,j := 0,len(s)-1;i<j;{
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
}

func reverseString1(s []byte)  {
	lenth := len(s)
	for i:= 0;i<lenth/2;i++{
		s[i],s[lenth-1-i] = s[lenth-1-i],s[i]

	}
}