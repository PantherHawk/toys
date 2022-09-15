import java.lang.String;
import java.lang.Integer;
import java.util.Map;
import java.util.HashMap;
/* Write a function fib(n) that returns the nth Fibonacci number. Fibonacci's sequence is defined as F(n) = F(n-1) + F(n-2) with F(0), F(1) = 1
 * 1, 1, 2, 3, 5, 8, 13 ...
 * */

class FibonacciNumbers {
	private static  Map<Integer, Integer> memo = new HashMap<>();
	public static void main(String[] args) {
		assert fib(0) == 1;
		assert fib(1) == 1;
		assert fib(2) == 2;
		assert fib(3) == 3;
		assert fib(10) == 89;
		assert fib(20) == 10946;
	}
	private static int fib(int n) {
		//  Check if we've encountered this input before.
		if (memo.containsKey(n)) {
			return memo.get(n);
		}
		// Base cases
		// check if it's 0 or 1
		if (n == 0 || n == 1) {
			return 1;
		}
		int output = fib(n-2) + fib(n-1);
		// Memoize the solution
		memo.put(n, output);
		return output;
	}
}
