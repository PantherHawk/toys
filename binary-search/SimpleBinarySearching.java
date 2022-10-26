public class SimpleBinarySearching {
	public static void main(String[] args) {
		System.out.println(recursiveBinarySearch(new int[]{1, 2, 3, 4, 5, 6, 7}, 4, 0, 6));
	}

	static int recursiveBinarySearch(int[] nums, int target, int start, int end) {
		if (start < end) {
			int mid = start + ((end - start) / 2);
			if (nums[mid] == target) {
				return mid;
			} else if (nums[mid] > target) {
				return recursiveBinarySearch(nums, target, start, mid - 1);
			} else {
				return recursiveBinarySearch(nums, target, mid + 1, end);
			}
		} else {
			return -1;
		}
	}

	static int binarySearch(int[] nums, int target) {
		// nums is sorted
		int end = nums.length - 1;
		int start = 0;
		
		while (start < end) {
			int mid = start + ((end - start) / 2);
			
			if (nums[mid] > target) {
				end = mid - 1;
			} else if (nums[mid] < target) {
				start = mid + 1;
			} else {
				return mid;
			}
		}
		return -1;
	}
}
