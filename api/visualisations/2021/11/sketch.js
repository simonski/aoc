let data = {};
let lastTime = 0;
let frameDuration = 100;
let currentLine = -1;

let CANVAS_WIDTH = 800;
let CANVAS_HEIGHT = 800;

function preload() {
  let callback = postPreload;
  data = loadJSON('/api/2021/11', callback);
  // data = loadJSON('data.json', callback)
}

function postPreload() {
}

function setup() {
    createCanvas(CANVAS_WIDTH, CANVAS_HEIGHT); //, WEBGL);
}
  
function draw() {
    let thisTime = millis();
    if ((thisTime - lastTime) < frameDuration) {
      return;
    }
    background(200);
    let cell_width = CANVAS_WIDTH/10;
    let cell_height = CANVAS_HEIGHT/10;

    translate(cell_width/2, -cell_height/2);
    lastTime = thisTime;
    currentLine++;
    if (currentLine == data["Lines"].length) {
      currentLine = 0;
    }

    let line = data["Lines"][currentLine];
    let row = 0;
    let col = 0;

    console.log("currentLine[" + currentLine + "]=" +line)
    for (let index=0; index<line.length; index++) {
      if (index % 10 == 0) {
        row++;
        col = 0;
        // new
      } else {
        col += 1;
      }
      let character = line.charAt(index);
      let x = cell_width * col;
      let y = cell_height * row;
      text(character, x, y);

    }
   
  }