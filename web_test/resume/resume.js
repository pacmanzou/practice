var flag = true;
var img = document.getElementById("img1");
var imgh = img.height;
var imgw = img.width;
img.onclick = function () {
  if (flag) {
    this.style.width = "500px";
    this.style.height = "500px";
    flag = false;
  } else {
    this.style.height = "100px";
    this.style.width = "100px";
    flag = true;
  }
};
