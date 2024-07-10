/*
Given a roman numeral, convert it to an int
Numeral to Int:
I - 1
V - 5
X - 10
L - 50
C - 100
D - 500
M - 1000

Subtraction is used:
I before V,X - 4,9
X before L,C - 40,90
C before D,M - 400,900

Success:

Runtime: 8 ms, faster than 55.14% of Java online submissions for Roman to Integer.
Memory Usage: 45.3 MB, less than 57.79% of Java online submissions for Roman to Integer.

*/

class Solution {
	public static int romanToInt(String s) {
		char[] arr = s.toCharArray();
		char[] check = { 'I','X','C' };
		char[] checkI = { 'V','X' };
		char[] checkX = { 'L','C' };
		char[] checkC = { 'D','M' };
		int fin = 0;

		for(int i=0; i<arr.length; i++) {
			int current = conversion(arr[i]);
			if(contains(arr[i],check)) {
				if(i+1 != arr.length) {
					int next = conversion(arr[i+1]);
					if(arr[i] == 'I' && contains(arr[i+1],checkI)) {
						int x = next-current;
						fin += x;
						i++;
					} else if(arr[i] == 'X' && contains(arr[i+1],checkX)) {
						int x = next-current;
						fin += x;
						i++;
					} else if(arr[i] == 'C' && contains(arr[i+1],checkC)) {
						int x = next-current;
						fin += x;
						i++;
					} else {
						fin += current;
					}
				} else {
					fin += current;
				}
			} else {
				fin += current;
			}
		}

		return fin;
	}

	static int conversion(char c) {
		if(c == 'I') {
			return 1;
		} else if(c == 'V') {
			return 5;
		} else if(c == 'X') {
			return 10;
		} else if(c == 'L') {
			return 50;
		} else if(c == 'C') {
			return 100;
		} else if(c == 'D') {
			return 500;
		} else if(c == 'M') {
			return 1000;
		} else {
			return -1;
		}
	}

	static boolean contains(char c, char[] array) {
        	for (char x : array) {
            		if (x == c) {
                		return true;
            		}
        	}
        	return false;
    	}
}
