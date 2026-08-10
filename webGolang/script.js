function clickButton() {
    console.log("button clicked");
    var count = 10
    count = count + 5
    console.log(count)
    document.body.style.color="green"
    document.getElementById("myp").style.color="red"
    document.getElementById("myp").style.fontSize=50
    
}
function changingColor(){
    document.getElementById("myp").style.color="black"
}