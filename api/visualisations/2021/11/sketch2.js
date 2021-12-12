let data = {};
let lastTime = 0;
let frameDuration = 100;
let currentLine = -1;

let CANVAS_WIDTH = 900;
let CANVAS_HEIGHT = 800;
let INFO_WIDTH = 100;

let WINDOW_WIDTH = window.innerWidth;
let WINDOW_HEIGHT = window.innerHeight;

let last_beat = 0;
let timing = 4;         // beats to the bar
let bpm_dead = 60;      // bpm or 1 b per 4 seconds
let bpm_slow = 60;      // bpm or 1bps
let bpm_normal = 120;   // bpm or 2bps
let bpm_fast = 240;       

// here I am trying to show the timing of each frame
let bpm = bpm_normal;
let timing_array = [ 1,1,1,1,2,2,2,2,1,1,1,1,0 ];
let millis_per_beat = (60 / bpm) * 1000;

function canvasResize() {
  console.log("resize");
  WINDOW_WIDTH = window.innerWidth;
  WINDOW_HEIGHT = window.innerHeight;
  console.log("window(w,h) (" + WINDOW_WIDTH + "," + WINDOW_HEIGHT + ")");
  // canvas.width = canvasWidth;
  // canvas.height = canvasHeight;
};

window.addEventListener("resize", canvasResize);

function preload() {
  let callback = postPreload;
  // data = loadJSON('/api/2021/11', callback)
  data = loadJSON('data.json', callback);
}

function postPreload() {
}

function setup() {
    createCanvas(CANVAS_WIDTH, CANVAS_HEIGHT); //, WEBGL);
}
  
function draw() {
  drawV2();
}

function drawV1() {

    let thisTime = millis();
    if ((thisTime - lastTime) < frameDuration) {
      return;
    }
    background(255  , 255,255);
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

function drawV2() {
    // [8,2 1,2 1,2 1,2]
    let millis_now = millis();
    let beat = int((millis_now / millis_per_beat) % timing) + 1;
    let double_beat = int((millis_now / (millis_per_beat/4)) % timing/2) + 1;

    if (last_beat != beat) {
      currentLine++;
      last_beat = beat;
    } else {
      return;
    }
    
    // let thisTime = millis();
    // if ((thisTime - lastTime) < frameDuration) {
    //   return;
    // }
    background(255, 255, 255);
    let cell_width = CANVAS_WIDTH/10;
    let cell_height = CANVAS_HEIGHT/10;

    // lastTime = thisTime;
    // currentLine++;
    if (currentLine == data["Lines"].length) {
      currentLine = 0;
    }
    textFont("Courier New");
    textSize(12);
    text("Beat " + beat + ", Doublebeat " + double_beat + ", Step " + currentLine, 0, 10)

    translate(cell_width/2, -cell_height/2);
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
