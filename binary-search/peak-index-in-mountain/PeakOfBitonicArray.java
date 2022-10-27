public class PeakOfBitonicArray {
	
	public static void main(String[] args) {
		System.out.println(findPeak(new int[]{1, 2, 3, 5, 6, 4, 3, 2}));
	}

	static int findPeak(int[] nums) {
		int start = 0;
		int end = nums.length - 1;

		while (start < end) {
			int mid = start + (end - start) / 2;

			if (nums[mid] > nums[mid + 1]) {
				end = mid;
			} else {
				start = mid + 1;
			}
		}
		return nums[start];
	}

}
