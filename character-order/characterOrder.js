const getCharacterOrder = (array) => {
	const result = new Set();
	// iterate over the sorted list of words
	for (let i = 0; i < array.length - 1; i++) {
		// iterate through the current word and the subsequent
		const current = array[i];
		const next = array[i+1];

		for (let j = 0; j < current.length; j++) {
			// if current word's current character is different from the subsequent
				// if the current word's letter comes after the next word's letter in the result set
					// delete the next word's letter from the result set
				// add both letters to the result set
		}
	}
	return result;
}
