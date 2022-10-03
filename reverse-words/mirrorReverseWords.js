function mirrorReverse(arr, start, end) {
	let temp = null;
	while(start < end) {
		temp = arr[start];
		arr[start] = arr[end];
		arr[end] = temp;
		start++;
		end--;
	}
}

function reverseWords(arr) {
	const n = arr.length;
	mirrorReverse(arr, 0, n - 1);
	
	let wordStart = null;
	for (let i = 0; i < n; i++) {
		// if we reach the end of a white indicated by a white space
		if (arr[i].trim() === '') {
			// we have marked the start of a word earlier,
			if (wordStart !== null) {
				// then we have reached the end of a word, mirror reverse it
				mirrorReverse(arr, wordStart, i - 1);
				wordStart = null;
			}
		// if we have reached the end of the sentence,	
		} else if (i === n - 1) {
			// we have our last word to reverse
			if (wordStart !== null) {
				mirrorReverse(arr, wordStart, i);
			}
		} else {
		// we are at the beginning of a new word
			if (wordStart === null) {
				wordStart = i;
			}
		}
	}
	return arr;
}

console.log(reverseWords(['p', 'e', 'r', 'f', 'e', 'c', 't', '  ', '  ',
                'm', 'a', 'k', 'e', 's', '  ',
                'p', 'r', 'a', 'c', 't', 'i', 'c', 'e' ]));
