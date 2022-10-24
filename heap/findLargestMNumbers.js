class PriorityQueue {
	constructor() {
		this.store = [];
	}

	insert(value) {
		this.store.push(value);

		let current = this.store.length - 1;
		let parent = Math.floor(current / 2);

		while (this.store[parent] && value > this.store[parent]) {
			// swap parent and current
			let temp = this.store[parent];
			this.store[parent] = value;
			this.store[current] = temp;

			current = parent;
			parent = Math.floor(current / 2);
		}
	}
}

function findLargetstM(nums, m) {
	const heap = new PriorityQueue();
	for (let v of nums) {
		heap.insert(v);
	}
	console.log("heap", heap)
	return heap.store.slice(0, m);
}

console.log(findLargetstM([1,5,4,2,3], 3));

