const printArray = (array) => {
  array.forEach(row => {
    console.log(row)
  });
}

const printArray2 = (array) => {
  array.forEach(row => {
    console.log(row.map(item => String.fromCharCode(item === 0 ? 9617 : 9619)).join(''));
  });
}

function getGeneration(cells, generations) {
  if (generations === 0) {
    return cells
  }

  const countNeighbors = (cells) => {
    const rows = cells.length
    const columns = cells[0].length
    const neighbors = new Array(rows + 2)
      .fill(0)
      .map(() => new Array(columns + 2)
        .fill(0))

    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < columns; j++) {
        if (cells[i][j] === 1) {
          neighbors[i][j] += 1
          neighbors[i][j + 1] += 1
          neighbors[i][j + 2] += 1
          neighbors[i + 1][j] += 1
          neighbors[i + 1][j + 1] += 0.5
          neighbors[i + 1][j + 2] += 1
          neighbors[i + 2][j] += 1
          neighbors[i + 2][j + 1] += 1
          neighbors[i + 2][j + 2] += 1
        }
      }
    }

    return neighbors
  }

  const cutBorders = (array) => {
    while (array[0].every(item => item === 0)) {
      array.shift()
    }

    while (array[array.length - 1].every(item => item === 0)) {
      array.pop()
    }

    while (array.every(row => row[0] === 0)) {
      array.forEach(row => row.shift())
    }

    while (array.every(row => row[row.length - 1] === 0)) {
      array.forEach(column => column.pop())
    }
  }

  const nextGen = countNeighbors(cells)
    .map(rowItem => rowItem
      .map(item => item >= 2.5 && item <= 3.5 ? 1 : 0))
  cutBorders(nextGen)

  return getGeneration(nextGen, generations - 1)
}

let gliders = [
  [1, 0, 0],
  [0, 1, 1],
  [1, 1, 0]
];

printArray(gliders)
console.log('---------------------------------')
printArray(getGeneration(gliders, 100))