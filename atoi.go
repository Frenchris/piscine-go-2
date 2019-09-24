package piscine

func Atoi(s string) int{
	stri := []rune(s)

	res:=0
	offset:=0

	if stri[0] == 45 || stri[0]==43 {
		offset = 1
	}else{
		offset = 0
	}
		for i := offset; i < StrLength(s) ; i++  {
			
			if '0' <= stri[i] && stri[i] <= '9'{
				exp:= StrLength(s) - 1 - i
				if stri[i] == '0'{
					res += 0 * Power(10,exp) 
			} else if stri[i] == '1'{
				res += 1 * Power(10,exp) 
			} else if stri[i] == '2'{
				res += 2 * Power(10,exp) 
			} else if stri[i] == '3'{
				res += 3 * Power(10,exp) 
			} else if stri[i] == '4'{
				res += 4 * Power(10,exp) 
			} else if stri[i] == '5'{
				res += 5 * Power(10,exp) 
			} else if stri[i] == '6'{
				res += 6 * Power(10,exp) 
			} else if stri[i] == '7'{
				res += 7 * Power(10,exp) 
			} else if stri[i] == '8'{
				res += 8 * Power(10,exp) 
			} else if stri[i] == '9'{
				res += 9 * Power(10,exp) 
			}
			} else {
				return 0
			}
		} 
			
			return res

}
	






// int toString(char a[]) {
//   int c, sign, offset, n;
 
//   if (a[0] == '-') {  // Handle negative integers
//     sign = -1;
//   }
 
//   if (sign == -1) {  // Set starting position to convert
//     offset = 1;
//   }
//   else {
//     offset = 0;
//   }
 
//   n = 0;
 
//   for (c = offset; a[c] != '\0'; c++) {
//     n = n * 10 + a[c] - '0';
//   }
 
//   if (sign == -1) {
//     n = -n;
//   }
 
//   return n;
// }