let data = {};
let CANVAS_WIDTH = 800;
let CANVAS_HEIGHT = 800;

function preload() {
    data = loadJSON('/api/2021/09')
}

function setup() {
    createCanvas(CANVAS_WIDTH, CANVAS_HEIGHT); //, WEBGL);
    noLoop();
}
  
function draw() {
    background(200);

    let cell_width = CANVAS_WIDTH/data.width;
    let cell_height = CANVAS_HEIGHT/data.height;

    // now draw boxes on every 3rd square
    push();
    for (let col=0; col<data.width; col++) {
      for (let row=0; row<data.height; row++) {
        let key = col + "," + row
        let value = data.points[key]

        let x1 = col*cell_width;
        let x2 = x1 + cell_width;
        let y1 = row*cell_height;
        let y2 = y1 + cell_height
        rect(x1, y1, x2, y2);

        stroke(0);
        strokeWeight(0.8);
        
        fill(0, 255/10*(value+1), 125);
      }
    }
    pop()
    
  }