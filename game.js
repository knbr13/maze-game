var boundaries = document.getElementsByClassName("boundary");
var start = document.getElementById("start");
var isClicked = false;
var gameBounds = document.getElementById("game").getBoundingClientRect();
var key = document.createElement('a');
key.textContent = ""
key.style.cssText = "width: 10%; height: 10%; font-size: 12px; position: absolute; left: 38%; top: 33%; background-color: blue; border-radius: 40%; border: 0.1px solid white"
start.appendChild(key);
document.getElementById("end").addEventListener("mousedown", reset);
start.addEventListener("mousedown", reset);
key.addEventListener("mousedown", mouseClicked);
function mouseClicked(e){
    e.stopPropagation();
    isClicked ? isClicked = false: isClicked = true;
    document.addEventListener("mousemove", move);
}
function move(e){
    start.style.top = `${e.clientY - 15 - gameBounds.top}px`;
    start.style.left = `${e.clientX - 15 - gameBounds.left}px`;
    checkForCollision(boundaries);
    checkWin() ? youWin(): null;
    fallDown();
}
function checkForCollision(arrOfElements){
    Array.from(arrOfElements).forEach((element) => {
        var elementBounds = element.getBoundingClientRect();
        if( (start.getBoundingClientRect().left >= elementBounds.left - 40 && start.getBoundingClientRect().right <= elementBounds.right + 40) 
        && (start.getBoundingClientRect().bottom > elementBounds.bottom && start.getBoundingClientRect().top <= elementBounds.bottom) ){
            element.style.borderBottom = `3px solid red`;
            document.removeEventListener("mousemove", move);
            youLose();
        }else if((start.getBoundingClientRect().left >= elementBounds.left - 40  && start.getBoundingClientRect().right <= elementBounds.right + 40) 
        && (start.getBoundingClientRect().bottom >= elementBounds.top && start.getBoundingClientRect().top < elementBounds.top) ){
            element.style.borderTop = `3px solid red`;
            document.removeEventListener("mousemove", move);
            youLose();
        }else if( (start.getBoundingClientRect().top >= elementBounds.top - 40 && start.getBoundingClientRect().bottom <= elementBounds.bottom + 40) 
        && (start.getBoundingClientRect().right >= elementBounds.left && start.getBoundingClientRect().left < elementBounds.left) ){
            element.style.borderLeft = `3px solid red`;
            document.removeEventListener("mousemove", move);
            youLose();
        }else if( (start.getBoundingClientRect().top >= elementBounds.top - 40 && start.getBoundingClientRect().bottom <= elementBounds.bottom + 40) 
        && (start.getBoundingClientRect().left <= elementBounds.right && start.getBoundingClientRect().right > elementBounds.right) ){
            element.style.borderRight = `3px solid red`;
            document.removeEventListener("mousemove", move);
            youLose();
        }
    })
}
function checkWin(){
    return start.getBoundingClientRect().right >= gameBounds.right - 40? true: false;
}
function youWin(){
    var winSound = document.createElement("audio");
    winSound.src = "./sounds/win_sound_effect.mp3";
    winSound.volume = 100;
    winSound.play();
    var you_win = document.createElement("h1");
    you_win.textContent = "You Win!";
    you_win.style.cssText = "z-index: 10; font-family: cursive; position: fixed; left: 45%; color: purple; top: 24%;";
    document.getElementById("game").appendChild(you_win);
    setTimeout(()=> you_win.remove(), 2000);
    boundaries[boundaries.length - 1].textContent = `Score: ${Number((boundaries[boundaries.length - 1].textContent).slice(6)) + 5}`;
    document.removeEventListener("mousemove", move);
}
function fallDown(){
    if(start.getBoundingClientRect().right<gameBounds.left){
        document.removeEventListener("mousemove", move);
        var interval = setInterval(() => {
            start.style.top = `${a}px`;
            a += 1.8;
            start.getBoundingClientRect().bottom > innerHeight - 10? start.style.display = `none`: null;
        }, 1), a = 205;
        setTimeout(()=> {
            clearInterval(interval);
            youLose();
        }, 1000);
    }
}  
function youLose(){
    var alertSound = document.createElement("audio");
    alertSound.src = "./sounds/alert_sound.mp3";
    alertSound.volume = 1;
    alertSound.play();
    var you_lose = document.createElement("h1");
    you_lose.textContent = "You Lose!";
    you_lose.style.cssText = "z-index: 10; font-family: cursive; position: fixed; left: 45%; color: purple; top: 24%;";
    document.getElementById("game").appendChild(you_lose);
    setTimeout(()=> you_lose.remove(), 2000);
    boundaries[boundaries.length - 1].textContent = `Score: ${Number((boundaries[boundaries.length - 1].textContent).slice(6)) - 5}`;
}
function reset(e){
    start.style.cssText = "position: absolute; top: 205px";
}