class RecursePatterns {
	public static void main(String[] args) {
		upsideDownTriangle(3, 0);
	}
	
	static void upsideDownTriangle(int r, int c) {
		if (r == 0) {
			return;
		}
		if (c < r) {
			upsideDownTriangle(r, c + 1);
			System.out.print("*");
		} else {
			upsideDownTriangle(r - 1, 0);
			System.out.println();
		}
	}

	static void triangle(int r, int c) {
		if (r == 0) {
			return;
		}
		if (c < r) {
			System.out.print("*");
			triangle(r, c + 1);
		} else {
			System.out.println();
			triangle(r - 1, 0);
		}
	}
}
