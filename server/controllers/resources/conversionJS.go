//Genereated by WGF -- DO NOT EDIT
package resources

const ConversionJS = `var fn = function() {
  //console.log(h, v, t, g);
  var i = new Image(),
    c = Math.floor(Math.random() * 1000),
    b = Math.floor(Date.now() / 1000000);
  i.src = "//" + h + "/" + v + "/c.gif?t=" + t + "&g=" + g + "&b=" + b + "&c=" + c;
}
if (document.readyState != 'loading') {
  fn();
} else {
  document.addEventListener('DOMContentLoaded', fn);
}`

