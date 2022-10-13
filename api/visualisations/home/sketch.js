let data = {};
let CANVAS_WIDTH = 800;
let CANVAS_HEIGHT = 800;

function preload() {
  data = loadJSON('/api/solutions')
}

function setup() {
  createCanvas(CANVAS_WIDTH, CANVAS_HEIGHT); //, WEBGL);
  noLoop();
}

function draw() {
  background(200);

  data.width = 26;
  data.height = 5;

  let cell_width = CANVAS_WIDTH / data.width;
  let cell_height = cell_width; //CANVAS_HEIGHT / data.height;

  // now draw boxes on every 3rd square
  console.log("draw() data.width=" + data.width + ", data.height=" + data.height);
  console.log("draw() data=" + data);
  console.log("draw() cell_height=" + cell_height + ", cell_width=" + cell_width);

  push();

  let xoffset = cell_width;
  let yoffset = 0;//cell_height;

  for (let day = 1; day <= 24; day++) {
    for (let year = 2015; year <= 2022; year++) {
      let key = year + "." + day;
      let col = day-1;
      let row = year - 2014;

      let x1 = xoffset + (col * cell_width);
      let x2 = (x1 + cell_width);
      let y1 = yoffset + (row * cell_height);
      let y2 = (y1 + cell_height);
      rect(x1, y1, x2, y2);
      let entry = data.solutions[key];
      stroke(0);
      strokeWeight(0.8);

      let value = 0.5;

      if (entry != undefined && entry["c1"] == false) {
        fill(0, 255 / 10 * (value + 1), 90);
      } else {
        fill(128, 255 / 10 * (value + 1), 225);
      }
    }

  }
  fill(0, 0, 0);
  pop()
  console.log("hi-x");

}