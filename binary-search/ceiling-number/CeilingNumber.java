public class CeilingNumber {
	public static void main(String[] args) {
		System.out.println(ceilingNumber(new int[]{1, 2, 3, 5}, 4));
	}

	static int ceilingNumber(int[] nums, int target) {
		int s = 0;
		int e = nums.length - 1;

		while(s < e) {
			int mid = s + (e - s) / 2;
			
			if (target < nums[mid]) {
				e = mid - 1;
			} else if (target > nums[mid]) {
				s = mid + 1;
			} else {
				return nums[mid];
			}
		}
		return nums[s];

	}
}
