function shortestCellPath(grid, sr, sc, tr, tc) {
	/**
	@param grid: integer[][]
	@param sr: integer
	@param sc: integer
	@param tr: integer
	@param tc: integer
	@return: integer
	*/

	// your code goes here

	const queue = [];
	const visited = new Set();

	queue.push({x: sr, y: sc, d: 0});
	visited.add(`${sr}-${sc}`);

	while (queue.length) {
		const node = queue.shift();
		const { x, y , d } = node;

		if (x === tr && y === tc) return d;

		for (const child of [{x: x - 1, y: y}, {x: x + 1, y: y}, {x: x, y: y - 1}, {x: x, y: y + 1}]) {
			const nx = child.x;
			const ny = child. y;

			if (nx >= 0 && nx < grid.length && ny >= 0 && ny < grid[nx].length && grid[nx][ny] === 1 & !visited.has(`${nx}-${ny}`)) {
				queue.push({x: nx, y: ny, d: d + 1});
				visited.add(`${nx}-${ny}`);
			}
		}
	}
	return -1; 
}
