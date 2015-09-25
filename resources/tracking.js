var fn = function() {
  console.log(a, h, v, g);
  var i = new Image();
  i.src = "//" + h + "/" + v + "/v.gif?a=" + a + "&g=" + g;
};
if (document.readyState != 'loading') {
  fn();
} else {
  document.addEventListener('DOMContentLoaded', fn);
}