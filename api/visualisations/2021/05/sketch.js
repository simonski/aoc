class AOCLine {
    constructor(x1, y1, x2, y2) {
      this.x1 = x1;
      this.y1 = y1;
      this.x2 = x2;
      this.y2 = y2;
      this.over = false;
    }
  
    display() {
      stroke(0);
      strokeWeight(0.8);
      noFill();
      line(this.x1, this.y1, this.x2, this.y2);
    }
}

let data = {};
let lines = [];

function preload() {
    data = loadJSON('/api/2021/05')
}

function loadData() {
    let lineData = data['lines'];
    for (let i = 0; i < lineData.length; i++) {
      let l = lineData[i];
      let x1 = l['x1'];
      let y1 = l['y1'];
      let x2 = l['x2'];
      let y2 = l['y2'];
      lines.push(new AOCLine(x1, y1, x2, y2));
    }
  }

function setup() {
    createCanvas(1000, 1000);
    loadData();
}
  
function draw() {
    background(200);
    for (let i=0; i<lines.length; i++) {
        lines[i].display()
    }
}