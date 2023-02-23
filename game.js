var start = document.getElementById("start");
var isClicked = false;
var gameBounds = document.getElementById("game").getBoundingClientRect();
start.addEventListener("mousedown", mouseClicked);

console.log(gameBounds.left);
console.log(gameBounds.top);

function mouseClicked(e){
    isClicked ? isClicked = false: isClicked = true;
    document.addEventListener("mousemove", move);
}

function move(e){
    console.log(e.clientX);
    console.log(e.clientY);
    start.style.top = `${e.clientY - 15 - gameBounds.top}px`;
    start.style.left = `${e.clientX - 15 - gameBounds.left}px`;
}