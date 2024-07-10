/*
Ransom Note 

*/
class Solution {
	public static void main(String[] args) {
		System.out.println(canConstruct("aa","ab"));
	}

	public static boolean canConstruct(String ransomNote, String magazine) {
		char[] arr = ransomNote.toCharArray();
		char[] arr2 = magazine.toCharArray();
	}

	static char[] remove(char c, char[] array) {
		char[] arr = new char[array.length-1];
		int index;

		for(int i=0; i<array.length; i++) {
			if(array[i] == c) {
				index = i;
				break;
			}
		}
		for(int i=0; i<array.length; i++) {
			if(i != index) {
				arr[i++] = array[i];
			}				
		}

		return arr;

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
