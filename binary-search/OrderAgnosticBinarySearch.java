public class OrderAgnosticBinarySearch {
	public static void main(String[] args) {
		System.out.println(orderAgnosticBinarySearch(new int[]{56, 33, 24, 22, 11, 10, 4, 2}, 33));
	}

	static int orderAgnosticBinarySearch(int[] nums, int target) {
		// check first and last elements to see whether its ordered ascending or descending
		int start = 0;
		int end = nums.length - 1;
		
		boolean isAscending = nums[start] < nums[end];

		while (start < end) {
			int mid = start + (end - start) / 2;

			if (target < nums[mid]) {
				end = isAscending ? mid - 1 : end;
				start = isAscending ? start : mid + 1;
			} else if (target > nums[mid]) {
				start = isAscending ? mid + 1 : start;
				end = isAscending ? end : mid - 1;
			} else {
				return mid;
			}
		}
		return -1;
	}
}

