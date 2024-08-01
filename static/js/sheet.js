const sContainer = document.getElementById("scon")
const downloadButton = document.getElementById("export-btn")
const activeContaner = document.getElementById("active-container")

const ROWS = 10
const COLS = 10
let spreadSheet = []

const alphabets = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');

class Cell {
  constructor(isHeader, disabled, data, row, column, rowName, columnName, active = false) {
    this.isHeader = isHeader,
      this.disabled = disabled,
      this.data = data,
      this.row = row,
      this.column = column,
      this.active = active,
      this.rowName = rowName,
      this.columnName = columnName
  }

  get name() {
    return `${this.columnName}${this.rowName}`
  }
}
initSpreadSheet()

downloadButton.addEventListener("click", (_) => {
  let csv = ""
  for (let i = 0; i < spreadSheet.length; i++) {
    csv += spreadSheet[i].filter(item => !item.isHeader).map(item => item.data).join(",") + "\r\n"
  }
  const csvBlob = new Blob([csv])
  const csvFileUrl = URL.createObjectURL(csvBlob)
  const a = document.createElement("a")
  a.href = csvFileUrl
  a.download = "Exported.csv"
  a.click()
})

function initSpreadSheet() {
  for (let i = 0; i < ROWS; i++) {
    let spreadSheetRow = []
    for (let j = 0; j < COLS; j++) {
      let cellData = ""
      let isHeader = false
      let isDisabled = false
      if (j === 0) {
        cellData = i
        isHeader = true
        isDisabled = true
      }

      if (i === 0) {
        cellData = alphabets[j - 1]
        isDisabled = true
        isHeader = true
      }
      if (!cellData) {
        cellData = ""
      }
      let rowName = i
      let columnName = alphabets[j - 1]
      const cell = new Cell(isHeader, isDisabled, cellData, i, j, rowName, columnName, false)
      spreadSheetRow.push(cell)

    }
    spreadSheet.push(spreadSheetRow)
  }

  drawSheet()
}

function drawSheet() {
  for (let i = 0; i < spreadSheet.length; i++) {
    const rowContanerEl = document.createElement("div")
    rowContanerEl.className = "cell-row"
    for (let j = 0; j < spreadSheet[i].length; j++) {
      const cell = spreadSheet[i][j]
      rowContanerEl.append(createCellEl(cell))

    }

    sContainer.append(rowContanerEl)
  }
}

function createCellEl(cell) {
  const cellEl = document.createElement("input")
  cellEl.className = "cell"
  cellEl.id = "cell_" + cell.row + cell.column
  cellEl.value = cell.data
  cellEl.disabled = cell.disabled
  cellEl.onclick = () => handlerCellClick(cell)
  cellEl.onchange = (e) => handlerCellOnchage(e.target.value, cell)
  return cellEl
}


function handlerCellClick(cell) {
  clearActiveState()
  const rowHeader = spreadSheet[cell.row][0]
  const columnHeader = spreadSheet[0][cell.column]
  const columnHeaderEl = getElFromRowCol(columnHeader.row, columnHeader.column)
  const rowHeaderEl = getElFromRowCol(rowHeader.row, rowHeader.column)
  columnHeaderEl.classList.add("active")
  rowHeaderEl.classList.add("active")
  activeContaner.innerText = cell.name

  console.log(cell.name)
  /*
  console.log("clicked cell: ", cell)
  console.log("header: ", columnHeaderEl)
  console.log("row header: ", rowHeaderEl)
  */
}


function handlerCellOnchage(data, cell) {
  cell.data = data
}


function clearActiveState() {

  for (let i = 0; i < spreadSheet.length; i++) {
    for (let j = 0; j < spreadSheet[i].length; j++) {
      const cell = spreadSheet[i][j]
      if (cell.isHeader) {
        let cellEl = getElFromRowCol(cell.row, cell.column)
        cellEl.classList.remove("active")
      }
    }
  }
}

function getElFromRowCol(row, col) {
  return document.querySelector("#cell_" + row + col)
}

