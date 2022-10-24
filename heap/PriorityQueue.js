class LinkedListNode {
	constructor(value) {
		this.value = value;
		this.next = null;
	}
}

class PriorityQueueLinkedList {
	constructor() {
		this.head = null;
	}

	insert(value) {
		const node = new LinkedListNode(value);
		if (!this.head || this.head.value < value) {
			this.first.next = this.first;
			this.first = node;
		} else {
			let pointer = this.head;
			while(pointer.next && pointer.next.value > value) {
				pointer = pointer.next;
			}
			node.next = pointer.next;
			pointer.next = node;
		}
	}
	pop() {
		if (!this.head) {
			return;
		}
		let temp = this.head;
		this.head = this.head.next;
		return temp.value;
	}
}

class PriorityQueueArray {
	constructor() {
		this.heap = [null];
	}

	insert(value) {
		this.heap.push(value);
		let currentIndex = this.heap.length - 1;
		let currentParentIndex = Math.floor(currentIndex / 2);
		while (this.heap[currentParentIndex] && value > this.heap[currentParentIndex]) {
			const parent = this.heap[currentParentIndex];
			this.heap[currentParentIndex] = value;
			this.heap[currentIndex] = parent;
			currentIndex = currentParentIndex;
			currentParentIndex = Math.floor(currentIndex / 2);
		}
	}
}
