// given a rotated array like [5, 6, 7, 8, 9, 1, 2, 3], find a target integer
//

class RotatedBinarySearch {
	public static void main(String[] args) {
		int[] arr = {5, 6, 7, 8, 9, 1, 2, 3};
		System.out.println(rotatedBinarySearch(arr, 7, 0, arr.length - 1));
		System.out.println(rotatedBinarySearch(arr, 2, 0, arr.length - 1));
	}

	static int rotatedBinarySearch(int[] arr, int target, int start, int end) {
		// if arr[start] <= arr[mid]
		if (start > end) {
			return -1;
		}
		int mid = start + (end - start) / 2;
		if (arr[mid] == target) {
			return mid;
		}
		if (arr[start] <= arr[mid]) {
		  // if arr[start} <= target <= arr[mid]
		  if (target >= arr[start] && target <= arr[mid]) {
		    // end = mid - 1
		    return rotatedBinarySearch(arr, target, start, mid - 1);
		  // else
		  } else {
		    // start = mid + 1
		    return rotatedBinarySearch(arr, target, mid + 1, end);
		  }
		} 
		// else, if arr[start] > arr[mid]
		  // if arr[mid] <= target <= arr[end]
		if (target >= arr[mid] && target <= arr[end]) {
		    // start = mid + 1
		    return rotatedBinarySearch(arr, target, mid + 1, end);
		  // else 
		  }
		    // end = mid - 1
		return rotatedBinarySearch(arr, target, start, mid - 1);
	}	
}
