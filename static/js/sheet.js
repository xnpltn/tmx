const sContainer = document.getElementById("scon");
const downloadButton = document.getElementById("export-btn");
const activeContainer = document.getElementById("active-container");
const alphabets = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');

const ROWS = 100;
const COLS = alphabets.length + 1;
let spreadSheet = [];

class Cell {
  constructor(isHeader, disabled, data, row, column, rowName, columnName, active = false) {
    this.isHeader = isHeader;
    this.disabled = disabled;
    this.data = data;
    this.row = row;
    this.column = column;
    this.active = active;
    this.rowName = rowName;
    this.columnName = columnName;
    this.name = `${this.columnName}${this.rowName}`;
    this.sheetID = undefined;
  }
}

function getSheetId(location) {
  const url = new URL(location);
  const parts = url.pathname.split('/');
  const id = parts[parts.length - 1];
  return id

}

downloadButton.addEventListener("click", (_) => {
  let data = "";
  for (let i = 0; i < spreadSheet.length; i++) {
    data += spreadSheet[i].filter(item => !item.isHeader).map(item => item.data).join(",") + "\r\n";
  }
  const dataBlob = new Blob([csv]);
  const csvFileUrl = URL.createObjectURL(dataBlob);
  const a = document.createElement("a");
  a.href = csvFileUrl;
  a.download = "Exported.csv";
  a.click();
});


async function fetchData() {
  try {
    const response = await fetch(`/api/cells/${getSheetId(window.location.href)}`);
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    return [];
  }
}

async function initSpreadSheet() {
  const fetchedData = await fetchData();

  for (let i = 0; i < ROWS; i++) {
    let spreadSheetRow = [];
    for (let j = 0; j < COLS; j++) {
      let cellData = "";
      let isHeader = false;
      let isDisabled = false;
      if (j === 0) {
        cellData = i;
        isHeader = true;
        isDisabled = true;
      }
      if (i === 0) {
        cellData = alphabets[j - 1];
        isDisabled = true;
        isHeader = true;
      }
      if (!cellData) {
        cellData = "";
      }
      let rowName = i;
      let columnName = alphabets[j - 1];
      const cell = new Cell(isHeader, isDisabled, cellData, i, j, rowName, columnName, false);
      spreadSheetRow.push(cell);
    }
    spreadSheet.push(spreadSheetRow);
  }

  // Populate the cells with the fetched data
  for (const item of fetchedData) {
    const { row, column, data } = item;
    if (spreadSheet[row] && spreadSheet[row][column]) {
      spreadSheet[row][column].data = data;
    }
  }

  drawSheet();
}

initSpreadSheet();

function drawSheet() {
  sContainer.innerHTML = "";
  for (let i = 0; i < spreadSheet.length; i++) {
    const rowContainerEl = document.createElement("div");
    rowContainerEl.className = "cell-row";
    for (let j = 0; j < spreadSheet[i].length; j++) {
      const cell = spreadSheet[i][j];
      rowContainerEl.append(createCellEl(cell));
    }
    sContainer.append(rowContainerEl);
  }
}

function createCellEl(cell) {
  const cellElement = document.createElement("input");
  cellElement.className = `cell ${cell.isHeader ? 'cell-header' : ''}`;
  cellElement.id = "cell_" + cell.row + cell.column;
  cellElement.value = cell.data;
  cellElement.disabled = cell.disabled;
  cellElement.onclick = () => handlerCellClick(cell);
  cellElement.onchange = async (e) => await handlerCellOnchange(e.target.value, cell);
  return cellElement;
}

function handlerCellClick(cell) {
  clearActiveState();
  const rowHeader = spreadSheet[cell.row][0];
  const columnHeader = spreadSheet[0][cell.column];
  const columnHeaderEl = getElFromRowCol(columnHeader.row, columnHeader.column);
  const rowHeaderEl = getElFromRowCol(rowHeader.row, rowHeader.column);
  columnHeaderEl.classList.add("active");
  rowHeaderEl.classList.add("active");
  activeContainer.innerText = cell.name;
}

console.log(window.location.href);
async function handlerCellOnchange(data, cell) {
  cell.data = data;
  cell.sheetID = parseInt(getSheetId(window.location.href));
  try {
    const response = await fetch('/api/save', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(cell),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const result = await response.json();
    console.log('Save successful:', result);
  } catch (error) {
    console.error('Error saving data:', error);
  }
}

function clearActiveState() {
  for (let i = 0; i < spreadSheet.length; i++) {
    for (let j = 0; j < spreadSheet[i].length; j++) {
      const cell = spreadSheet[i][j];
      if (cell.isHeader) {
        let cellElement = getElFromRowCol(cell.row, cell.column);
        cellElement.classList.remove("active");
      }
    }
  }
}

function getElFromRowCol(row, col) {
  return document.querySelector("#cell_" + row + col);
}

