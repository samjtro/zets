/*
Fizz Buzz

Given an integer n, return a string array answer(1-indexed) where:
answer[i] == "FizzBuzz" if i is divisible by 3 and 5, 
answer[i] == "Fizz" if i is divisible by 3, 
answer[i] == "Buzz" if i is divisible by 5,
answer[i] == i (as a string) if none of the above conditions are true.

Success:
Runtime: 3 ms, faster than 26.03% of Java online submissions for Fizz Buzz.
Memory Usage: 48.3 MB, less than 46.74% of Java online submissions for Fizz Buzz.

*/

import java.util.*;

class Solution {
	public static void main(String[] args) {
		System.out.println(fizzBuzz(15).toString());
	}

	public static List<String> fizzBuzz(int n) {
		List<String> arr = new ArrayList<String>();
		for(int i=1; i<=n; i++) {
			if(i%3 == 0 && i%5 == 0) {
				arr.add("FizzBuzz");
			} else if(i%3 == 0) {
				arr.add("Fizz");
			} else if(i%5 == 0) {
				arr.add("Buzz");
			} else {
				arr.add(Integer.toString(i));
			}
		}
		return arr;
	}	
}
